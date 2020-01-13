package command

import (
	"fmt"
	"strings"
	"time"

	"github.com/hirakiuc/daily-workflow/config"
	"github.com/hirakiuc/daily-workflow/service"
	"github.com/pkg/errors"
	cli "github.com/urfave/cli/v2"
)

type DailyCommand struct {
	Date time.Time

	Conf *config.Config
}

const JSTDiffSeconds = 9 * 60 * 60

func NewDailyCommand() *cli.Command {
	srv := DailyCommand{}

	t := time.Now()

	loc, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		loc = time.FixedZone("Asia/Tokyo", JSTDiffSeconds)
	}

	t = t.In(loc)
	today := t.Format("2006-0102")

	return &cli.Command{
		Name:    "daily",
		Aliases: []string{"d"},
		Usage:   "Add a daily memo",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "date,d",
				Aliases: []string{"d"},
				Value:   today,
				Usage:   "yyyy-mmdd",
			},
		},
		Subcommands: makeSubCommands(srv),
	}
}

func makeSubCommands(srv DailyCommand) []*cli.Command {
	return []*cli.Command{
		{
			Name:   "edit",
			Usage:  "Edit the target daily reports",
			Action: srv.EditAction,
		},
		{
			Name:   "list",
			Usage:  "list daily reports",
			Action: srv.ListAction,
			Flags: []cli.Flag{
				&cli.BoolFlag{
					Name:    "edit,e",
					Aliases: []string{"e"},
					Value:   false,
				},
			},
		},
		{
			Name:   "find",
			Usage:  "find daily reports",
			Action: srv.FindAction,
			Flags: []cli.Flag{
				&cli.BoolFlag{
					Name:    "edit",
					Aliases: []string{"e"},
					Value:   false,
				},
			},
		},
	}
}

func (s *DailyCommand) parseArgs(c *cli.Context) error {
	conf, err := config.LoadConfig("./config.toml")
	if err != nil {
		return err
	}

	s.Conf = conf

	d := c.String("date")

	s.Date, err = time.Parse("2006-0102", d)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalida date:%s", d))
	}

	return nil
}

func (s *DailyCommand) EditAction(c *cli.Context) error {
	err := s.parseArgs(c)
	if err != nil {
		return err
	}

	fs := service.NewFsService(s.Conf.DailyPath())

	dirPath := fmt.Sprintf("%04d", s.Date.Year())
	if err := fs.MakeDirs(dirPath); err != nil {
		return err
	}

	fPath := s.Conf.DailyFilePath(
		dirPath,
		fmt.Sprintf("%02d%02d.md", s.Date.Month(), s.Date.Day()),
	)

	// Open vim with the target path.
	cmd := service.NewCmdService(s.Conf)

	return cmd.EditAndWait(fPath, "")
}

func (s *DailyCommand) findCandidates(_ *cli.Context, words []string) ([]string, error) {
	fs := service.NewFsService(s.Conf.Common.Root)

	if len(words) == 0 {
		return fs.ListFiles(s.Conf.DailyPath())
	}

	return fs.FindFiles(
		s.Conf.Common.Finder,
		s.Conf.Common.FinderOpts,
		strings.Join(words, " "),
	)
}

const CaseCandidateIsOnlyPath int = 1
const CaseCandidateIsVimdiff int = 3

func (s *DailyCommand) chooseAndEdit(_ *cli.Context, candidates []string) error {
	srv := service.NewCmdService(s.Conf)

	results, err := srv.Choose(candidates)
	if err != nil {
		return err
	}

	switch len(results) {
	case 0:
		return nil
	case CaseCandidateIsOnlyPath:
		target := results[0]
		opts := ""

		parts := strings.Split(target, ":")

		if len(parts) == CaseCandidateIsVimdiff {
			target = parts[0]
			row := parts[1]
			// col := parts[2]

			opts = fmt.Sprintf("+%s", row)
		}

		srv := service.NewCmdService(s.Conf)

		return srv.EditAndWait(target, opts)
	default:
		return fmt.Errorf("unsupported case: can't select multiple items")
	}
}

func (s *DailyCommand) ListAction(c *cli.Context) error {
	err := s.parseArgs(c)
	if err != nil {
		return err
	}

	paths, err := s.findCandidates(c, []string{})
	if err != nil {
		return err
	}

	if c.Bool("e") {
		return s.chooseAndEdit(c, paths)
	}

	for _, path := range paths {
		fmt.Println(path)
	}

	return nil
}

func (s *DailyCommand) FindAction(c *cli.Context) error {
	if err := s.parseArgs(c); err != nil {
		return err
	}

	candidates, err := s.findCandidates(c, c.Args().Slice())
	if err != nil {
		fmt.Println("find failure")
		return err
	}

	if len(candidates) == 0 {
		return nil
	}

	if c.Bool("e") {
		return s.chooseAndEdit(c, candidates)
	}

	for _, v := range candidates {
		fmt.Println(v)
	}

	return nil
}

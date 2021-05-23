package daily

import (
	"fmt"
	"strings"
	"time"

	base "github.com/hirakiuc/daily-workflow/command"
	"github.com/hirakiuc/daily-workflow/service"
	"github.com/pkg/errors"
	cli "github.com/urfave/cli/v2"
)

const JSTDiffSeconds = 9 * 60 * 60

// ErrUnsupportedCase represents an un-expected case error.
var ErrUnsupportedCase = errors.New("unsupported case: can't select multiple items")

type Command struct {
	Date time.Time

	*base.Base
}

func NewCommand(iostream *base.IoStream) *cli.Command {
	srv := Command{
		Base: base.NewBase(iostream),
	}

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
		Subcommands: makeDailySubCommands(srv),
	}
}

func makeDailySubCommands(srv Command) []*cli.Command {
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
				&cli.StringFlag{
					Name:    "path",
					Aliases: []string{"p"},
					Value:   "",
				},
			},
		},
	}
}

func (s *Command) parseArgs(c *cli.Context) error {
	err := s.LoadConfig("./config.toml")
	if err != nil {
		return errors.Wrap(err, "Failed to load config")
	}

	d := c.String("date")

	s.Date, err = time.Parse("2006-0102", d)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalida date:%s", d))
	}

	return nil
}

func (s *Command) EditAction(c *cli.Context) error {
	err := s.parseArgs(c)
	if err != nil {
		return err
	}

	fs := service.NewFsService(s.Conf.DailyPath())

	dirPath := fmt.Sprintf("%04d", s.Date.Year())
	if err := fs.MakeDirs(dirPath); err != nil {
		return fmt.Errorf("failed to create directories: %w", err)
	}

	fPath := s.Conf.DailyFilePath(
		dirPath,
		fmt.Sprintf("%02d%02d.md", s.Date.Month(), s.Date.Day()),
	)

	// Open vim with the target path.
	cmd := service.NewCmdService(s.Conf)

	err = cmd.EditAndWait(fPath, "")
	if err != nil {
		return fmt.Errorf("failed to wait the command result: %w", err)
	}

	return nil
}

func (s *Command) findCandidates(_ *cli.Context, words []string) ([]string, error) {
	fs := service.NewFsService(s.Conf.Common.Root)

	if len(words) == 0 {
		lines, err := fs.ListFiles(s.Conf.DailyPath())
		if err != nil {
			return []string{}, fmt.Errorf("failed to show the file list: %w", err)
		}

		return lines, nil
	}

	lines, err := fs.FindFiles(
		s.Conf.Common.Finder,
		s.Conf.Common.FinderOpts,
		strings.Join(words, " "),
	)
	if err != nil {
		return []string{}, fmt.Errorf("failed to show the file list: %w", err)
	}

	return lines, nil
}

const (
	// CaseCandidateIsOnlyPath represents the case that the candidate is only path string.
	CaseCandidateIsOnlyPath int = 1

	// CaseCandidateIsVimdiff represents the case that the candidate is vimdiff.
	CaseCandidateIsVimdiff int = 3
)

func (s *Command) chooseAndEdit(_ *cli.Context, candidates []string) error {
	srv := service.NewCmdService(s.Conf)

	results, err := srv.Choose(candidates)
	if err != nil {
		return fmt.Errorf("failed to choose target: %w", err)
	}

	for _, path := range results {
		s.Out().Println(path)
	}

	if len(results) == 0 {
		return nil
	} else if len(results) > 1 {
		return fmt.Errorf("%w", ErrUnsupportedCase)
	}

	target := results[0]
	opts := ""

	parts := strings.Split(target, ":")
	if len(parts) > CaseCandidateIsVimdiff {
		target = parts[0]
		row := parts[1]
		// col := parts[2]

		opts = fmt.Sprintf("+%s", row)
	}

	err = srv.EditAndWait(target, opts)
	if err != nil {
		return fmt.Errorf("failed: %w", err)
	}

	return nil
}

func (s *Command) ListAction(c *cli.Context) error {
	if err := s.parseArgs(c); err != nil {
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
		s.Out().Println(path)
	}

	return nil
}

func (s *Command) FindAction(c *cli.Context) error {
	if err := s.parseArgs(c); err != nil {
		return err
	}

	founds, err := s.findCandidates(c, c.Args().Slice())
	if err != nil {
		s.Out().Println("find failure")

		return fmt.Errorf("failed to find candidates: %w", err)
	}

	candidates := []string{}

	if len(c.String("path")) > 0 {
		pathKey := c.String("path")

		for _, path := range founds {
			if strings.Contains(path, pathKey) {
				candidates = append(candidates, path)
			}
		}
	} else {
		candidates = founds
	}

	if len(candidates) == 0 {
		return nil
	}

	if c.Bool("e") {
		return s.chooseAndEdit(c, candidates)
	}

	for _, v := range candidates {
		s.Out().Println(v)
	}

	return nil
}

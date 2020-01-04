package command

import (
	"fmt"
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
				Name: "date,d", Aliases: []string{"d"},
				Value: today,
				Usage: "yyyy-mmdd",
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
		},
		{
			Name:   "find",
			Usage:  "find daily reports",
			Action: srv.FindAction,
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
	cmd := service.NewCmdService()

	return cmd.Exec(s.Conf.Common.Editor, fPath)
}

func (s *DailyCommand) ListAction(c *cli.Context) error {
	err := s.parseArgs(c)
	if err != nil {
		return err
	}

	//	fs := service.NewFsService(s.Conf.Common.Root)

	return nil
}

func (s *DailyCommand) FindAction(c *cli.Context) error {
	fmt.Println("find diaries")
	return nil
}

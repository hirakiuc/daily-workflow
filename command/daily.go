package command

import (
	"fmt"
	"time"

	cli "github.com/urfave/cli/v2"
)

type DailyCommand struct {
	Date time.Time
}

func NewDailyCommand() cli.Command {
	srv := DailyCommand{}

	return cli.Command{
		Name:    "daily",
		Aliases: []string{"d"},
		Usage:   "Add a daily memo",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name: "date,d",
			},
		},
		Action: srv.Action,
	}
}

func (s *DailyCommand) parseArgs(c *cli.Context) error {
	if len(c.Args()) == 0 {
		// Use system date
		s.Date = time.Now()
	}

	return nil
}

func (s *DailyCommand) Action(c *cli.Context) error {
	err := s.parseArgs(c)
	if err != nil {
		return err
	}

	fmt.Println("Daily Action")

	return nil
}

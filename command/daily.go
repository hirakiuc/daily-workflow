package command

import (
	"fmt"

	"github.com/urfave/cli"
)

type DailyService struct {
}

func NewDailyCommand() cli.Command {
	srv := DailyService{}

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

func (s *DailyService) Action(c *cli.Context) error {
	fmt.Println("Daily Action")
	return nil
}

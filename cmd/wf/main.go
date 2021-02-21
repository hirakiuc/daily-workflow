package main

import (
	"fmt"
	"os"

	"github.com/hirakiuc/daily-workflow/command"
	cli "github.com/urfave/cli/v2"
)

func main() {
	app := cli.App{
		Name:                 "wf",
		Usage:                "workflow for daily tasks",
		EnableBashCompletion: true,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "config",
				Aliases: []string{"c"},
				Value:   "~/.config/wf/config.toml",
			},
		},
	}

	iostream := command.NewIoStream()

	app.Commands = []*cli.Command{
		command.NewDailyCommand(iostream),
		command.NewJiraCommand(iostream),
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)

		return
	}
}

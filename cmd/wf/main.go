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
				Value:   "./config.toml",
			},
		},
	}

	app.Commands = []*cli.Command{
		command.NewDailyCommand(),
		command.NewIssueCommand(),
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return
	}
}

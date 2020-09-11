package command

import (
	cli "github.com/urfave/cli/v2"
)

type JiraCommand struct {
}

func NewJiraCommand() *cli.Command {
	srv := JiraCommand{}

	return &cli.Command{
		Name:        "jira",
		Aliases:     []string{"j"},
		Usage:       "jira ticket",
		Subcommands: makeJiraSubCommands(srv),
	}
}

func makeJiraSubCommands(srv JiraCommand) []*cli.Command {
	return []*cli.Command{
		{
			Name:   "issues",
			Usage:  "Show issues",
			Action: srv.ShowIssuesAction,
		},
		{
			Name:   "boards",
			Usage:  "Show boards",
			Action: srv.ShowBoardsAction,
		},
	}
}

func (s *JiraCommand) ShowIssuesAction(c *cli.Context) error {
	return nil
}

func (s *JiraCommand) ShowBoardsAction(c *cli.Context) error {
	return nil
}

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
			Action: srv.ShowIssueAction,
		},
	}
}

func (s *JiraCommand) ShowIssueAction(c *cli.Context) error {
	return nil
}

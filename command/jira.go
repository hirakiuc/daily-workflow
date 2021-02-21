package command

import (
	cli "github.com/urfave/cli/v2"
)

type JiraCommand struct {
	*Base
}

func NewJiraCommand(iostream *IoStream) *cli.Command {
	srv := JiraCommand{
		Base: NewBase(iostream),
	}

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

func (s *JiraCommand) parseArgs(_ *cli.Context) error {
	err := s.LoadConfig("./config.toml")
	if err != nil {
		return err
	}

	return nil
}

func (s *JiraCommand) ShowIssuesAction(c *cli.Context) error {
	err := s.parseArgs(c)
	if err != nil {
		return err
	}

	return nil
}

func (s *JiraCommand) ShowBoardsAction(c *cli.Context) error {
	err := s.parseArgs(c)
	if err != nil {
		return err
	}

	return nil
}

package jira

import (
	base "github.com/hirakiuc/daily-workflow/command"
	"github.com/pkg/errors"
	cli "github.com/urfave/cli/v2"
)

type Command struct {
	*base.Base
}

func NewCommand(iostream *base.IoStream) *cli.Command {
	srv := Command{
		Base: base.NewBase(iostream),
	}

	return &cli.Command{
		Name:        "jira",
		Aliases:     []string{"j"},
		Usage:       "jira ticket",
		Subcommands: makeJiraSubCommands(srv),
	}
}

func makeJiraSubCommands(srv Command) []*cli.Command {
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

func (s *Command) parseArgs(_ *cli.Context) error {
	err := s.LoadConfig("./config.toml")
	if err != nil {
		return errors.Wrap(err, "Failed to load config")
	}

	return nil
}

func (s *Command) ShowIssuesAction(c *cli.Context) error {
	err := s.parseArgs(c)
	if err != nil {
		return err
	}

	return nil
}

func (s *Command) ShowBoardsAction(c *cli.Context) error {
	err := s.parseArgs(c)
	if err != nil {
		return err
	}

	return nil
}

package command

import (
	"fmt"

	"github.com/urfave/cli"
)

type IssueCommand struct {
}

func NewIssueCommand() cli.Command {
	srv := IssueCommand{}

	return cli.Command{
		Name:    "issue",
		Aliases: []string{"i"},
		Usage:   "Add issue memo",
		Action:  srv.Action,
	}
}

func (s *IssueCommand) Action(c *cli.Context) error {
	fmt.Println("Issue Action")
	return nil
}

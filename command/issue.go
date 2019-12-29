package command

import (
	"fmt"

	cli "github.com/urfave/cli/v2"
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

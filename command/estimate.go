package command

import (
	"fmt"

	"github.com/urfave/cli"
)

type EstimateCommand struct {
}

func NewEstimateCommand() cli.Command {
	srv := EstimateCommand{}

	return cli.Command{
		Name:    "estimate",
		Aliases: []string{"e"},
		Usage:   "Add a estimation memo",
		Action:  srv.Action,
	}
}

func (s *EstimateCommand) Action(c *cli.Context) error {
	fmt.Println("Estimate Action")
	return nil
}

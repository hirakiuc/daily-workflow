package command

import (
	"fmt"
	"io"
	"os"

	"github.com/hirakiuc/daily-workflow/config"
)

type Base struct {
	Conf *config.Config

	Stdout io.Writer
	Stderr io.Writer
}

func NewBase() *Base {
	return &Base{
		Conf: nil,

		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}
}

func (b *Base) LoadConfig(path string) error {
	c, err := config.LoadConfig(path)
	if err != nil {
		return fmt.Errorf("failed to load config file: %w", err)
	}

	b.Conf = c

	return nil
}

func (b *Base) Output(msg string, a ...interface{}) {
	fmt.Fprintf(b.Stdout, msg, a...)
}

func (b *Base) Err(msg string, a ...interface{}) {
	fmt.Fprintf(b.Stderr, msg, a...)
}

package command

import (
	"fmt"

	"github.com/hirakiuc/daily-workflow/config"
)

type Base struct {
	Conf *config.Config

	IO *IoStream
}

func NewBase(stream *IoStream) *Base {
	return &Base{
		Conf: nil,

		IO: stream,
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

func (b *Base) Out() *Printer {
	return b.IO.Out
}

func (b *Base) Err() *Printer {
	return b.IO.Err
}

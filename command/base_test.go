package command_test

import (
	"testing"

	"github.com/hirakiuc/daily-workflow/command"
	"github.com/stretchr/testify/assert"
)

func baseCommand() *command.Base {
	return command.NewBase(
		command.NewIoStream(),
	)
}

func TestNewBase(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name string
		io   *command.IoStream
	}{
		{
			name: "Create Base instance",
			io:   command.NewIoStream(),
		},
	}

	for _, c := range cases {
		c := c

		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			ast := assert.New(t)

			b := command.NewBase(c.io)
			ast.NotNil(b, "NewBase should return non-nil value")
		})
	}
}

func TestLoadConfig_withInvalid(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name string
		path string
	}{
		{
			name: "With empty config",
			path: "../testdata/configs/empty.toml",
		},
	}

	for _, c := range cases {
		c := c

		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			ast := assert.New(t)

			b := baseCommand()
			ast.NotNil(b)

			err := b.LoadConfig(c.path)
			ast.NotNil(err)

			ast.Nil(b.Conf)
		})
	}
}

func TestLoadConfig_withValid(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name string
		path string
	}{
		{
			name: "With valid config",
			path: "../testdata/configs/valid.toml",
		},
	}

	for _, c := range cases {
		c := c

		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			ast := assert.New(t)

			b := baseCommand()
			ast.NotNil(b)

			err := b.LoadConfig(c.path)
			ast.Nil(err)

			ast.NotNil(b.Conf)
		})
	}
}

func TestOut(t *testing.T) {
	t.Parallel()

	ast := assert.New(t)

	b := baseCommand()
	ast.NotNil(b)

	p := b.Out()
	ast.NotNil(p)
	ast.Equal(p, b.IO.Out)
}

func TestErr(t *testing.T) {
	t.Parallel()

	ast := assert.New(t)

	b := baseCommand()
	ast.NotNil(b)

	p := b.Err()
	ast.NotNil(p)
	ast.Equal(p, b.IO.Err)
}

package service

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/hirakiuc/daily-workflow/config"
	"gopkg.in/pipe.v2"
)

type CmdService struct {
	Config *config.Config
}

func NewCmdService(conf *config.Config) *CmdService {
	return &CmdService{
		Config: conf,
	}
}

func (s *CmdService) ExecAndWait(name string, args ...string) error {
	opts := []string{}

	for _, v := range args {
		if len(v) > 0 {
			opts = append(opts, v)
		}
	}

	cmd := exec.Command(name, opts...)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	cmd.Env = os.Environ()

	err := cmd.Start()
	if err != nil {
		return fmt.Errorf("failed to execute command: %w", err)
	}

	return cmd.Wait()
}

func (s *CmdService) EditAndWait(fpath string, opts string) error {
	return s.ExecAndWait(
		s.Config.Common.Editor,
		fpath,
		opts,
	)
}

func (s *CmdService) Choose(candidates []string) ([]string, error) {
	p := pipe.Line(
		pipe.Println(strings.Join(candidates, "\n")),
		pipe.Exec(s.Config.Common.Chooser),
	)

	output, err := pipe.CombinedOutput(p)
	if err != nil {
		return []string{}, fmt.Errorf("failed to extract output: %w", err)
	}

	lines := strings.TrimRight(string(output), "\n")

	return strings.Split(lines, "\n"), nil
}

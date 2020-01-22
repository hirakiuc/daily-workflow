package service

import (
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

	// nolint:gosec
	cmd := exec.Command(name, opts...)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	cmd.Env = os.Environ()

	err := cmd.Start()
	if err != nil {
		return err
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
	var chooserPipe pipe.Pipe
	if len(s.Config.Common.ChooserOpts) > 0 {
		chooserPipe = pipe.Exec(
			s.Config.Common.Chooser,
			s.Config.Common.ChooserOpts,
		)
	} else {
		chooserPipe = pipe.Exec(
			s.Config.Common.Chooser,
		)
	}

	p := pipe.Line(
		pipe.Print(strings.Join(candidates, "\n")+"\n"),
		chooserPipe,
	)

	output, err := pipe.CombinedOutput(p)
	if err != nil {
		return []string{}, err
	}

	lines := strings.TrimRight(string(output), "\n")

	return strings.Split(lines, "\n"), nil
}

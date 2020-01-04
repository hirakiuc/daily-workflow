package service

import (
	"os"
	"os/exec"
)

type CmdService struct {
}

func NewCmdService() *CmdService {
	return &CmdService{}
}

func (s *CmdService) Exec(name string, args ...string) error {
	cmd := exec.Command(name, args...)

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

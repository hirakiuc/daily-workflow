package service

import (
	"os"
	"os/exec"

	"github.com/hirakiuc/daily-workflow/config"
)

type EditorService struct {
	Config *config.Config
}

func NewEditorService(conf *config.Config) *EditorService {
	return &EditorService{
		Config: conf,
	}
}

func (s *EditorService) EditAndWait(path string, opts string) error {
	// nolint:gosec
	cmd := exec.Command(
		s.Config.Common.Editor,
		path,
		opts,
	)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Start()
	if err != nil {
		return err
	}

	return cmd.Wait()
}

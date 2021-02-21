package service

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type FsService struct {
	Root string
}

func NewFsService(root string) *FsService {
	return &FsService{
		Root: root,
	}
}

func (s *FsService) MakeDirs(path string) error {
	absPath := filepath.Join(s.Root, path)
	if _, err := os.Stat(absPath); !os.IsNotExist(err) {
		// Already exists
		return nil
	}

	err := os.MkdirAll(absPath, 0777)
	if err != nil {
		return fmt.Errorf("failed to mkdir 0777: %w", err)
	}

	return nil
}

func (s *FsService) ListFiles(root string) ([]string, error) {
	founds := []string{}

	err := filepath.Walk(root, func(path string, info os.FileInfo, e error) error {
		if info.IsDir() {
			return nil
		}

		rel, err := filepath.Rel(root, path)
		if err != nil {
			return fmt.Errorf("failed to get a relative path: %w", err)
		}

		if rel == "." {
			return nil
		}

		if info.IsDir() {
			return filepath.SkipDir
		}

		founds = append(founds, path)

		return nil
	})
	if err != nil {
		return []string{}, fmt.Errorf("failed to recuesively scan the directory: %w", err)
	}

	return founds, nil
}

func (s *FsService) FindFiles(finder string, finderOpts string, word string) ([]string, error) {
	// exec ag with words
	cmd := exec.Command(
		finder,
		finderOpts,
		word,
		s.Root,
	)

	// nolint:forbidigo
	fmt.Println(cmd.String())

	bytes, err := cmd.Output()
	if err != nil {
		return []string{}, fmt.Errorf("failed to get the output: %w", err)
	}

	outputs := strings.TrimRight(string(bytes), "\n")

	return strings.Split(outputs, "\n"), nil
}

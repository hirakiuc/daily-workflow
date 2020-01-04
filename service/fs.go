package service

import (
	"os"
	"path/filepath"
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
		return err
	}

	return nil
}

func (s *FsService) AbsPath(paths ...string) (string, error) {
	parts := append([]string{s.Root}, paths...)

	absPath, err := filepath.Abs(filepath.Join(parts...))
	if err != nil {
		return "", err
	}

	return absPath, nil
}

func (s *FsService) ListFiles(root string, words []string) ([]string, error) {
	founds := []string{}

	err := filepath.Walk(root, func(path string, info os.FileInfo, e error) error {
		if info.IsDir() {
			return nil
		}

		rel, err := filepath.Rel(root, path)
		if err != nil {
			return err
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
		return []string{}, err
	}

	return founds, nil
}

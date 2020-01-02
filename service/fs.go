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

func (s *FsService) Open(path string) (*os.File, error) {
	absPath, err := filepath.Abs(filepath.Join(s.Root, path))
	if err != nil {
		return nil, err
	}

	f, err := os.OpenFile(absPath, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return nil, err
	}

	return f, nil
}

func (s *FsService) MakeDirs(path string) error {
	absPath, err := filepath.Abs(filepath.Join(s.Root, path))
	if err != nil {
		return err
	}

	err = os.MkdirAll(absPath, 0777)
	if err != nil {
		return err
	}

	return nil
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

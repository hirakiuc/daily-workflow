package service

import (
	"os"
	"path/filepath"
)

type FileService struct {
}

func NewFileService() *FileService {
	return &FileService{}
}

func (s *FileService) Open(path string) (*os.File, error) {
	absPath, err := filepath.Abs(path)
	if err != nil {
		return nil, err
	}

	f, err := os.OpenFile(absPath, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return f, nil
}

func (s *FileService) ListFiles(root string, words []string) ([]string, error) {
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

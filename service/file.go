package service

type FileService struct {
}

func NewFileService() *FileService {
	return &FileService{}
}

func (s *FileService) ListFiles(prefix string) []string {
	return []string{}
}

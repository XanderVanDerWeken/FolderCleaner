package services

import (
	"os"
	"path/filepath"
	"strings"
)

type FileService interface {
	ListFiles(path string, extensions []string) ([]string, error)
	DeleteFile(filename string) error
}

type fileService struct {
}

func NewFileService() FileService {
	return &fileService{}
}

func (f fileService) ListFiles(path string, extensions []string) ([]string, error) {
	var filesToDelete []string

	err := filepath.Walk(path, func(filePath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			for _, ext := range extensions {
				if strings.HasSuffix(info.Name(), ext) {
					filesToDelete = append(filesToDelete, filePath)
				}
			}
		}

		return nil
	})

	return filesToDelete, err
}

func (f fileService) DeleteFile(filename string) error {
	return os.Remove(filename)
}

package services

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type FileService interface {
	ListFiles(path string, extensions []string)
	DeleteFiles(path string, extensions []string)
}

type fileService struct {
}

func NewFileService() FileService {
	return &fileService{}
}

func (f fileService) ListFiles(path string, extensions []string) {
	err := filepath.Walk(path, func(filePath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			for _, ext := range extensions {
				if strings.HasSuffix(info.Name(), ext) {
					fmt.Printf("Found %s\n", filePath)
				}
			}
		}

		return nil
	})

	if err != nil {
		fmt.Printf("Error cleaning folder: %v\n", err)
	} else {
		fmt.Println("Cleaning complete!")
	}
}

func (f fileService) DeleteFiles(path string, extensions []string) {
	err := filepath.Walk(path, func(filePath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			for _, ext := range extensions {
				if strings.HasSuffix(info.Name(), ext) {
					fmt.Printf("Deleting %s\n", filePath)
					if err := os.Remove(filePath); err != nil {
						fmt.Printf("Failed to delete: %s: %v\n", filePath, err)
					}
				}
			}
		}

		return nil
	})

	if err != nil {
		fmt.Printf("Error cleaning folder: %v\n", err)
	} else {
		fmt.Println("Cleaning complete!")
	}
}

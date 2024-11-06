package services

import "strings"

type ExtensionService interface {
	GetExtensions(extensionName string) []string
}

type extensionService struct {
}

func NewExtensionService() ExtensionService {
	return &extensionService{}
}

var latexExtensions = []string{
	".aux",
	".log",
	".nav",
	".out",
	".snm",
	".synctex.gz",
	".toc",
	".vrb",
}

var macExtensions = []string{
	".DS_Store",
}

var intellij = []string{
	".idea",
	".target",
}

func (e extensionService) GetExtensions(extensionName string) []string {
	switch strings.ToLower(extensionName) {
	case "latex":
		return latexExtensions
	case "macos", "mac":
		return macExtensions
	case "intellij":
		return intellij
	default:
		return []string{}
	}
}

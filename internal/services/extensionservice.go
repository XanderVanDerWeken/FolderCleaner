package services

type ExtensionService interface {
	GetExtensions(extensionName string) []string
}

type extensionService struct {
}

func NewExtensionService() ExtensionService {
	return &extensionService{}
}

var latexExtensions []string = []string{
	".aux",
	".log",
	".nav",
	".out",
	".snm",
	".synctex.gz",
	".toc",
}

func (e extensionService) GetExtensions(extensionName string) []string {
	switch extensionName {
	case "latex":
		return latexExtensions
	default:
		return []string{}
	}
}

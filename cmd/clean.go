package cmd

import (
	"github.com/spf13/cobra"
	"github.com/xandervanderweken/FolderCleaner/internal/services"
)

var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "Clean specified files from a folder",
	Run: func(cmd *cobra.Command, args []string) {
		fileService := services.NewFileService()
		extensionService := services.NewExtensionService()

		var extensions = extensionService.GetExtensions(template)

		if isDryRun {
			fileService.ListFiles(path, extensions)
		} else {
			fileService.DeleteFiles(path, extensions)
		}
	},
}

var path string
var isDryRun bool
var template string

func init() {
	rootCmd.AddCommand(cleanCmd)

	cleanCmd.Flags().StringVarP(&path, "path", "p", "", "Path of folder")
	cleanCmd.MarkFlagDirname("path")
	cleanCmd.MarkFlagRequired("path")

	cleanCmd.Flags().BoolVarP(&isDryRun, "dry", "d", false, "Should not delete - Running without deleting")

	cleanCmd.Flags().StringVarP(&template, "template", "t", "", "Template to use (latex, ...)")
}

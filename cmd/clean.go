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

		var tempExtensions []string = []string{".go"}

		if isDryRun {
			fileService.ListFiles(path, tempExtensions)
		} else {
			fileService.DeleteFiles(path, tempExtensions)
		}
	},
}

var path string
var isDryRun bool

func init() {
	rootCmd.AddCommand(cleanCmd)

	cleanCmd.Flags().StringVarP(&path, "path", "p", "", "Path of folder")
	cleanCmd.MarkFlagDirname("path")
	cleanCmd.MarkFlagRequired("path")

	cleanCmd.Flags().BoolVarP(&isDryRun, "dry", "d", false, "Should not delete - Running without deleting")
}

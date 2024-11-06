package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/xandervanderweken/FolderCleaner/internal/services"
)

var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "Clean specified files from a folder",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fileService := services.NewFileService()
		extensionService := services.NewExtensionService()

		var path = args[0]

		var extensions = extensionService.GetExtensions(template)

		if len(extensions) == 0 {
			fmt.Printf("No Template was found for: %s\n", template)
			return
		} else {
			fmt.Printf("Using template: %s\n", template)
		}

		var filesToDelete, err = fileService.ListFiles(path, extensions)

		if err != nil {
			fmt.Printf("Error cleaning folder: %v\n", err)
			return
		}

		if len(filesToDelete) == 0 {
			fmt.Println("No matching files found to delete")
			return
		}

		fmt.Println("Files that would be deleted")
		for _, file := range filesToDelete {
			fmt.Println(" -", file)
		}

		if !isForcing {
			fmt.Print("Do you want to delte those files (y/N): ")
			scanner := bufio.NewScanner(os.Stdin)
			scanner.Scan()
			input := strings.ToLower(strings.TrimSpace(scanner.Text()))
			if input != "y" {
				fmt.Println("Deletion canceled")
				return
			}
		}

		for _, file := range filesToDelete {
			fmt.Printf("Deleting %s\n", file)
			if err := fileService.DeleteFile(file); err != nil {
				fmt.Printf("Failed to delete %s: %v\n", file, err)
			}
		}

		fmt.Println("Cleaning Complete!")
	},
}

var isForcing bool
var template string

func init() {
	rootCmd.AddCommand(cleanCmd)

	cleanCmd.Flags().BoolVarP(&isForcing, "force", "f", false, "Should delete files, without asking")

	cleanCmd.Flags().StringVarP(&template, "template", "t", "", "Template to use (latex, ...)")
}

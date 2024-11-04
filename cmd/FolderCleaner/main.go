package main

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/xandervanderweken/FolderCleaner/internal/cleaner"
)

func main() {
	rootcmd := &cobra.Command{
		Use:   "folder-cleaner",
		Short: "A CLI Tool for cleaning folders",
	}

	rootcmd.AddCommand(cleaner.CleanCommand)

	if err := rootcmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

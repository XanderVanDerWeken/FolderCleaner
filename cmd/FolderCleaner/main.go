package main

import (
	"log"

	"github.com/spf13/cobra"
)

func main() {
	rootcmd := &cobra.Command{
		Use:   "folder-cleaner",
		Short: "A CLI Tool for cleaning folders",
	}

	if err := rootcmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

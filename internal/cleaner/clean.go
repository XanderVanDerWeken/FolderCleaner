package cleaner

import (
	"fmt"

	"github.com/spf13/cobra"
)

var CleanCommand = &cobra.Command{
	Use:   "clean",
	Short: "Clean specified files from a folder",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("clean called")
	},
}

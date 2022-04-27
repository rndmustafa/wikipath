package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(pathCmd)
}

var pathCmd = &cobra.Command{
	Use:   "path",
	Short: "Find a path between two wiki articles",
	Run: func(cmd *cobra.Command, args []string) {
		start := args[0]
		end := args[1]
		fmt.Println(start + " " + end)
	},
}

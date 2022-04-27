package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "wiki",
	Short: "Find a path between two Wikipedia articles",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}

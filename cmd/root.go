package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var verbose bool

func init() {
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose flag for debug level logs")
}

var rootCmd = &cobra.Command{
	Use:   "wiki",
	Short: "Find a path between two Wikipedia articles",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Panic(err)
	}
}

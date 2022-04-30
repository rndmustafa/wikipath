package cmd

import (
	"time"
	"wikipath/pathing"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(pathCmd)
}

var pathCmd = &cobra.Command{
	Use:   "path",
	Short: "Find a path between two wiki articles",
	Run: func(cmd *cobra.Command, args []string) {
		if verbose {
			logrus.SetLevel(logrus.DebugLevel)
		}

		start := args[0]
		end := args[1]

		startTime := time.Now()
		path, err := pathing.Search(start, end)
		elapsedTime := time.Since(startTime)

		if err != nil {
			logrus.Panic(err)
		}
		logrus.Infof("Path: %v Time elapsed: %v", path, elapsedTime)
	},
}

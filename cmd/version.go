package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/stolenzc/got/config"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Got version",
	Long:  "Got version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Got version %s\n", config.Version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

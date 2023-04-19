package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(startCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Pomogo",
	Long:  `All software has versions. This is Pomogo's`,
	Run: func(cmd *cobra.Command, args []string) {
		panic("not impleted yet")
	},
}

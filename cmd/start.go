package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(startCmd)
}

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the Pomogo",
	Long:  ``, // TODO: write
	Run: func(cmd *cobra.Command, args []string) {
		panic("not impleted yet")
	},
}

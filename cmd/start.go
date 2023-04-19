package cmd

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"time"

	"github.com/peterramaldes/pomogo/internal/pomo"
	"github.com/spf13/cobra"
)

var description string

func init() {
	rootCmd.AddCommand(startCmd)
	startCmd.Flags().StringVarP(&description, "description", "d", "", "Description (required)")
}

var startCmd = &cobra.Command{
	Use: "start",
	RunE: func(cmd *cobra.Command, args []string) error {
		// Get the home dir path
		path, err := os.UserHomeDir()
		if err != nil {
			return err
		}

		// Get the filejson if exists and add a new content there

		var mock = []pomo.Pomo{
			{
				Start:       time.Now(),
				Description: description,
			},
		}

		j, err := json.MarshalIndent(mock, "", " ")
		if err != nil {
			return err
		}

		err = ioutil.WriteFile(path+"/.pomo.json", j, 0666) // 0666 indicates we going to create if not exists
		if err != nil {
			return err
		}

		return nil
	},
}

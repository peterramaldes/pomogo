package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/peterramaldes/pomogo/internal/pomo"
	"github.com/spf13/cobra"
)

var acitivity string

func init() {
	rootCmd.AddCommand(startCmd)
	startCmd.Flags().StringVarP(&acitivity, "activity", "a", "", "Activity (required)")
}

var startCmd = &cobra.Command{
	Use: "start",
	RunE: func(cmd *cobra.Command, args []string) error {
		// Description are required
		if len(acitivity) == 0 {
			return errors.New("description is required")
		}

		currentTrackingFile, err := currentTrackingFile()
		if err != nil {
			return err
		}

		currentTrackingFile.StoreCurrentPomo()

		// TODO: Put in the result the last Pomo
		// TODO: Start a new Pomo

		j, err := json.MarshalIndent(currentTrackingFile, "", " ")
		if err != nil {
			return err
		}

		filepath, err := getFilePath()
		if err != nil {
			return err
		}

		// NOTE: 0666 indicates we going to create if not exists
		err = ioutil.WriteFile(filepath, j, 0666)
		if err != nil {
			return err
		}

		return nil
	},
}

// currentTrackingFile is reponsible for get all pomodoros created at moment. If the file
// is not created storing the pomo's this method is responsible for create it
func currentTrackingFile() (pomo.TrackingFile, error) {
	filepath, err := getFilePath()
	if err != nil {
		return pomo.TrackingFile{}, err
	}

	f, err := os.OpenFile(filepath, os.O_RDONLY|os.O_CREATE, 0666)
	if err != nil {
		return pomo.TrackingFile{}, err
	}
	defer f.Close()

	b, err := ioutil.ReadAll(f)
	if err != nil {
		return pomo.TrackingFile{}, err
	}

	var currentTrackingFile pomo.TrackingFile
	json.Unmarshal(b, &currentTrackingFile)

	return currentTrackingFile, nil
}

// getFilePath is responsible for get the entire path for the pomo stored file
func getFilePath() (string, error) {
	path, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s/%s", path, ".pomo.json"), nil
}

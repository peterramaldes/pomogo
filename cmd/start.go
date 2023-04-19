package cmd

import (
	"encoding/json"
	"errors"
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
		// Description are required
		if len(description) == 0 {
			return errors.New("description is required")
		}

		pomos, err := getPomos()
		if err != nil {
			return err
		}

		newPomo := pomo.NewPomo(time.Now(), description)
		pomos = append(pomos, newPomo)

		j, err := json.MarshalIndent(pomos, "", " ")
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

// getPomos is reponsible for get all pomodoros created at moment. If the file
// is not created storing the pomo's this method is responsible for create it
func getPomos() ([]pomo.Pomo, error) {
	filepath, err := getFilePath()
	if err != nil {
		return []pomo.Pomo{}, err
	}

	f, err := os.OpenFile(filepath, os.O_RDONLY|os.O_CREATE, 0666)
	if err != nil {
		return []pomo.Pomo{}, err
	}
	defer f.Close()

	b, err := ioutil.ReadAll(f)
	if err != nil {
		return []pomo.Pomo{}, err
	}

	var pomos []pomo.Pomo
	json.Unmarshal(b, &pomos)

	return pomos, nil
}

func getFilePath() (string, error) {
	// Get the home dir path
	path, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return path + "/.pomo.json", nil
}

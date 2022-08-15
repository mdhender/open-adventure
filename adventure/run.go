// Open Adventure - a clone of Open Adventure written in Go
// Copyright (c) 2022 Michael D Henderson
// SPDX-License-Identifier: BSD-2-clause

package adventure

import (
	"encoding/json"
	"fmt"
	"github.com/mdhender/open-adventure/state"
	"os"
)

func New() (*state.Adventure, error) {
	adv, err := state.NewAdventure()
	if err != nil {
		return nil, err
	}
	return adv, nil
}

func SaveState(adv *state.Adventure, saveFile string) error {
	if b, err := json.MarshalIndent(adv, "", "  "); err != nil {
		return err
	} else if err = os.WriteFile(saveFile, b, 0666); err != nil {
		return err
	}
	return nil
}

// ExecuteCommand executes a single command, returning any errors and output
func ExecuteCommand(adv *state.Adventure, cmd string) (output string, err error) {
	return "", fmt.Errorf("not implemented")
}

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

func Execute(adv *state.Adventure, cmd string) (*state.Adventure, string, error) {
	if adv == nil {
		if adv, err := state.NewAdventure(); err != nil {
			return nil, "", err
		} else {
			return adv, adv.ArbitraryMessages.Map["WELCOME_YOU"], nil
		}
	}

	return adv, "", fmt.Errorf("not implemented")
}

func Save(adv *state.Adventure, saveFile string) error {
	if b, err := json.MarshalIndent(adv, "", "  "); err != nil {
		return err
	} else if err = os.WriteFile(saveFile, b, 0666); err != nil {
		return err
	}
	return nil
}

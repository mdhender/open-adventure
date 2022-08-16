// Open Adventure - a clone of Open Adventure written in Go
// Copyright (c) 2022 Michael D Henderson
// SPDX-License-Identifier: BSD-2-clause

package adventure

import (
	"fmt"
	"github.com/mdhender/open-adventure/state"
)

func Execute(adv *state.Adventure, cmd string) (*state.Adventure, string, error) {
	if adv == nil {
		if cmd != "" {
			return nil, "", fmt.Errorf("command must be blank when starting adventure")
		} else if adv, err := state.NewAdventure(); err != nil {
			return nil, "", err
		} else {
			return adv, adv.ArbitraryMessages.Map["WELCOME_YOU"], nil
		}
	}

	return adv, "", fmt.Errorf("not implemented")
}

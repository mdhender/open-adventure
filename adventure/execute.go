// Open Adventure - a clone of Open Adventure written in Go
// Copyright (c) 2022 Michael D Henderson
// SPDX-License-Identifier: BSD-2-clause

package adventure

import (
	"fmt"
	"github.com/mdhender/open-adventure/state"
	"strings"
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

	// treat empty input as end of input and quit the game
	if cmd == "" {
		return nil, `You scored 32 out of a possible 430, using 1 turn.

You are obviously a rank amateur.  Better luck next time.

To achieve the next higher rating, you need 14 more points.`, nil
	}

	if cmd == "n" {
		return adv, strings.Join(adv.Locations.Map["LOC_START"].LongDescr, "\n"), nil
	}

	// return nothing happens when we don't understand a command?
	return adv, adv.ArbitraryMessages.Map["NOTHING_HAPPENS"], nil
}

// Open Adventure - a clone of Open Adventure written in Go
// Copyright (c) 2022 Michael D Henderson
// SPDX-License-Identifier: BSD-2-clause

// Package adventure implements the main program of Open Adventure.
//
// * MAIN PROGRAM
// *
// *  Adventure (rev 2: 20 treasures)
// *  History: Original idea & 5-treasure version (adventures) by Willie Crowther
// *           15-treasure version (adventure) by Don Woods, April-June 1977
// *           20-treasure version (rev 2) by Don Woods, August 1978
// *           Errata fixed: 78/12/25
// *           Revived 2017 as Open Adventure.
package adventure

import "fmt"

type State struct {
	oldstyle bool
	seedval  int64
}

func WithLogFile(name string) func(*State) error {
	return func(s *State) error {
		return fmt.Errorf("log file: not implemented")
	}
}

func WithOldStyle() func(*State) error {
	return func(s *State) error {
		s.oldstyle = true
		return nil
	}
}

func WithPrompt() func(*State) error {
	return func(s *State) error {
		return fmt.Errorf("prompt: not implemented")
	}
}

func WithRestoreFile(name string) func(*State) error {
	return func(s *State) error {
		return fmt.Errorf("restore file: not implemented")
	}
}

func WithSeedValue(seed int64) func(*State) error {
	return func(s *State) error {
		s.seedval = seed
		return nil
	}
}

func (game *game_t) printf(format string, args ...interface{}) {
	if game != nil && game.stdout != nil {
		_, _ = game.stdout.Write([]byte(fmt.Sprintf(format, args...)))
	}
}

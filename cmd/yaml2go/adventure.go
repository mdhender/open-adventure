/*
 * Open Adventure - a clone of Open Adventure written in Go
 * Copyright (c) 2022 Michael D Henderson
 * SPDX-License-Identifier: BSD-2-clause
 */

package main

import (
	"fmt"
	"io"
)

func (y *YAML) GenerateAdventure(w io.Writer) error {
	_, _ = fmt.Fprintln(w, "// Adventure holds all the data needed for the game.")
	_, _ = fmt.Fprintln(w, "// In theory, writing out this would be all that's needed to save and restore a game.")
	_, _ = fmt.Fprintln(w, "type Adventure struct {")
	_, _ = fmt.Fprintln(w, "\tVersion           struct{ Major, Minor, Patch int; Suffix string }")
	_, _ = fmt.Fprintln(w, "\tActions           *Actions")
	_, _ = fmt.Fprintln(w, "\tArbitraryMessages *ArbitraryMessages")
	_, _ = fmt.Fprintln(w, "\tClasses           []*Class")
	_, _ = fmt.Fprintln(w, "\tLocations         *Locations")
	_, _ = fmt.Fprintln(w, "\tMotions           *Motions")
	_, _ = fmt.Fprintln(w, "\tObituaries        []*Obituary")
	_, _ = fmt.Fprintln(w, "\tObjects           *Objects")
	_, _ = fmt.Fprintln(w, "\tTurnThresholds    []*TurnThreshold")
	_, _ = fmt.Fprintln(w, "}")
	_, _ = fmt.Fprintln(w, "")
	_, _ = fmt.Fprintln(w, "// NewAdventure returns an initialized adventure.")
	_, _ = fmt.Fprintln(w, "func NewAdventure() (*Adventure, error) {")
	_, _ = fmt.Fprintln(w, "\treturn &Adventure{")
	_, _ = fmt.Fprintln(w, "\t\tVersion:           struct{ Major, Minor, Patch int; Suffix string }{Major: 1, Minor: 11},")
	_, _ = fmt.Fprintln(w, "\t\tActions:           generateActions(),")
	_, _ = fmt.Fprintln(w, "\t\tArbitraryMessages: generateArbitraryMessages(),")
	_, _ = fmt.Fprintln(w, "\t\tClasses:           generateClasses(),")
	_, _ = fmt.Fprintln(w, "\t\tLocations:         generateLocations(),")
	_, _ = fmt.Fprintln(w, "\t\tMotions:           generateMotions(),")
	_, _ = fmt.Fprintln(w, "\t\tObituaries:        generateObituaries(),")
	_, _ = fmt.Fprintln(w, "\t\tObjects:           generateObjects(),")
	_, _ = fmt.Fprintln(w, "\t\tTurnThresholds:    generateTurnThresholds(),")
	_, _ = fmt.Fprintln(w, "\t}, nil")
	_, _ = fmt.Fprintln(w, "}")
	return nil
}

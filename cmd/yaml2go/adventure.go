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
	fmt.Fprintln(w, "// Adventure holds all the data needed for the game.")
	fmt.Fprintln(w, "// In theory, writing out this would be all that's needed to save and restore a game.")
	fmt.Fprintln(w, "type Adventure struct {")
	fmt.Fprintln(w, "\tVersion           struct{ Major, Minor, Patch int; Suffix string }")
	fmt.Fprintln(w, "\tActions           *Actions")
	fmt.Fprintln(w, "\tArbitraryMessages *ArbitraryMessages")
	fmt.Fprintln(w, "\tClasses           []*Class")
	fmt.Fprintln(w, "\tLocations         *Locations")
	fmt.Fprintln(w, "\tMotions           *Motions")
	fmt.Fprintln(w, "\tObituaries        []*Obituary")
	fmt.Fprintln(w, "\tObjects           *Objects")
	fmt.Fprintln(w, "\tTurnThresholds    []*TurnThreshold")
	fmt.Fprintln(w, "}")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "// NewAdventure returns an initialized adventure.")
	fmt.Fprintln(w, "func NewAdventure() (*Adventure, error) {")
	fmt.Fprintln(w, "\treturn &Adventure{")
	fmt.Fprintln(w, "\t\tVersion:           struct{ Major, Minor, Patch int; Suffix string }{Major: 1, Minor: 8},")
	fmt.Fprintln(w, "\t\tActions:           generateActions(),")
	fmt.Fprintln(w, "\t\tArbitraryMessages: generateArbitraryMessages(),")
	fmt.Fprintln(w, "\t\tClasses:           generateClasses(),")
	fmt.Fprintln(w, "\t\tLocations:         generateLocations(),")
	fmt.Fprintln(w, "\t\tMotions:           generateMotions(),")
	fmt.Fprintln(w, "\t\tObituaries:        generateObituaries(),")
	fmt.Fprintln(w, "\t\tObjects:           generateObjects(),")
	fmt.Fprintln(w, "\t\tTurnThresholds:    generateTurnThresholds(),")
	fmt.Fprintln(w, "\t}, nil")
	fmt.Fprintln(w, "}")
	return nil
}

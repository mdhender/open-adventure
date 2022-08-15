// Open Adventure - a clone of Open Adventure written in Go
// Copyright (c) 2022 Michael D Henderson
// SPDX-License-Identifier: BSD-2-clause

// Package main implements the application.
package main

import (
	"encoding/json"
	"fmt"
	"github.com/mdhender/open-adventure/state"
	"os"
)

func main() {
	if err := run(); err != nil {
		fmt.Printf("%+v\n", err)
		os.Exit(2)
	}
	os.Exit(0)
}

func run() error {
	adv, err := state.NewAdventure()
	if err != nil {
		return err
	}
	b, err := json.MarshalIndent(adv, "", "  ")
	if err != nil {
		return err
	}
	if err = os.WriteFile("D:\\open-adventure\\testdata\\state.json", b, 0666); err != nil {
		return err
	}
	return nil
}

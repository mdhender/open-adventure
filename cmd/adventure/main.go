// Open Adventure - a clone of Open Adventure written in Go
// Copyright (c) 2022 Michael D Henderson
// SPDX-License-Identifier: BSD-2-clause

// Package main implements the application.
package main

import (
	"fmt"
	"github.com/mdhender/open-adventure/adventure"
	"os"
)

func main() {
	adv, err := adventure.New()
	if err != nil {
		fmt.Printf("%+v\n", err)
		os.Exit(2)
	}

	if err = adventure.SaveState(adv, "D:\\open-adventure\\testdata\\state.json"); err != nil {
		fmt.Printf("%+v\n", err)
		os.Exit(2)
	}

	os.Exit(0)
}

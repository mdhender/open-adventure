// Open Adventure - a clone of Open Adventure written in Go
// Copyright (c) 2022 Michael D Henderson
// SPDX-License-Identifier: BSD-2-clause

package tests

import (
	"github.com/mdhender/open-adventure/adventure"
	"testing"
)

func TestWelcome(t *testing.T) {
	adv, output, err := adventure.Execute(nil, "")
	if err != nil {
		t.Fatalf("?")
	} else if adv == nil {
		t.Fatalf("?")
	} else if output != "Welcome to Adventure!!  Would you like instructions?" {
		t.Fatalf("?")
	}
}

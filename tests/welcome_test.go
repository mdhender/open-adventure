// Open Adventure - a clone of Open Adventure written in Go
// Copyright (c) 2022 Michael D Henderson
// SPDX-License-Identifier: BSD-2-clause

package tests

import (
	"github.com/mdhender/open-adventure/adventure"
	"testing"
)

func Test_welcome(t *testing.T) {
	adv, got, err := adventure.Execute(nil, "")
	if err != nil {
		t.Fatalf("%v", err)
	}
	expected := adv.ArbitraryMessages.Map["WELCOME_YOU"]
	if expected != got {
		t.Fatalf("expected %q: got %q", expected, got)
	}
}

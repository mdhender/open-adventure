// Open Adventure - a clone of Open Adventure written in Go
// Copyright (c) 2022 Michael D Henderson
// SPDX-License-Identifier: BSD-2-clause

package tests

import (
	"github.com/mdhender/open-adventure/adventure"
	"strings"
	"testing"
)

func Test_isofoo(t *testing.T) {
	adv, got, err := adventure.Execute(nil, "")
	if err != nil {
		t.Fatalf("%v", err)
	}
	expected := adv.ArbitraryMessages.Map["WELCOME_YOU"]
	if expected != got {
		t.Fatalf("expected %q: got %q", expected, got)
	}

	if _, got, err = adventure.Execute(adv, "n"); err != nil {
		t.Fatalf("%v", err)
	}
	expected = strings.Join(adv.Locations.Map["LOC_START"].LongDescr, "\n")
	if expected != got {
		t.Fatalf("expected %q: got %q", expected, got)
	}

	if _, got, err = adventure.Execute(adv, "foo"); err != nil {
		t.Fatalf("?")
	}
	expected = adv.ArbitraryMessages.Map["NOTHING_HAPPENS"]
	if expected != got {
		t.Fatalf("expected %q: got %q", expected, got)
	}

	if _, got, err = adventure.Execute(adv, "foo"); err != nil {
		t.Fatalf("?")
	}
	expected = `You scored 32 out of a possible 430, using 1 turn.

You are obviously a rank amateur.  Better luck next time.

To achieve the next higher rating, you need 14 more points.`
	if expected != got {
		t.Fatalf("expected %q: got %q", expected, got)
	}
}

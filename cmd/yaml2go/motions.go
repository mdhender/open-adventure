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

// MOTION is used to load the YAML data for motions
type MOTION struct {
	Words []string
	// RawOldStyle is the value from the YAML file. it's a pointer because the
	// default value for it is true while Go expects it to be false. the loader
	// will test for nil; if found, it sets the value to true. otherwise, it uses
	// the value from the input.
	RawOldStyle *bool `yaml:"oldstyle"`
	// OldStyle will be populated by our loader.
	OldStyle bool `yaml:"-"`
}

// GenerateMotions converts
func (y *YAML) GenerateMotions(w io.Writer) error {
	fmt.Fprintf(w, "// Motions is an ordered map of motion data.\n")
	fmt.Fprintf(w, "// Use `Seq` to access sequentially or `Map` to access by key.\n")
	fmt.Fprintf(w, "type Motions struct{\n")
	fmt.Fprintf(w, "    Seq []*Motion            // sequential accessor\n")
	fmt.Fprintf(w, "    Map map[string]*Motion   // tag accessor\n")
	fmt.Fprintf(w, "}\n\n")

	fmt.Fprintf(w, "// Motion is autogenerated\n")
	fmt.Fprintf(w, "type Motion struct {\n")
	fmt.Fprintf(w, "    Words    []string\n")
	fmt.Fprintf(w, "    OldStyle bool\n")
	fmt.Fprintf(w, "}\n\n")

	fmt.Fprintf(w, "// generateMotions returns the initial state of motions\n")
	fmt.Fprintf(w, "func generateMotions() *Motions{\n")
	fmt.Fprintf(w, "    motions := Motions{\n")
	fmt.Fprintf(w, "        Map: make(map[string]*Motion),\n")
	fmt.Fprintf(w, "    }\n")
	for i, motionMap := range y.Motions {
		fmt.Fprintf(w, "\n")
		for name, motion := range motionMap {
			fmt.Fprintf(w, "    // seq: %d name: %s\n", i, name)
			if i == 0 {
				fmt.Fprintf(w, "    motion := &Motion{\n")
			} else {
				fmt.Fprintf(w, "    motion = &Motion{\n")
			}
			if len(motion.Words) != 0 {
				fmt.Fprintf(w, "        Words: []string{\n")
				for _, s := range motion.Words {
					fmt.Fprintf(w, "            %q,\n", s)
				}
				fmt.Fprintf(w, "        },\n")
			}
			if motion.OldStyle {
				fmt.Fprintf(w, "        OldStyle: true,\n")
			}
			fmt.Fprintf(w, "    }\n")
			fmt.Fprintf(w, "    motions.Seq = append(motions.Seq, motion)\n")
			fmt.Fprintf(w, "    motions.Map[%q] = motion\n", name)
		}
	}
	fmt.Fprintf(w, "\n")
	fmt.Fprintf(w, "    return &motions\n")
	fmt.Fprintf(w, "}\n")
	fmt.Fprintf(w, "\n")

	return nil
}

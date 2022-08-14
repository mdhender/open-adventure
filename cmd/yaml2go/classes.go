/*
 * Open Adventure - a clone of Open Adventure written in Go
 * Copyright (c) 2022 Michael D Henderson
 * SPDX-License-Identifier: BSD-2-clause
 */

package main

import (
	"fmt"
	"io"
	"strings"
)

// CLASS is used to load the YAML data for classes
type CLASS struct {
	Threshold int
	Message   string
}

// GenerateClasses converts class data.
func (y *YAML) GenerateClasses(w io.Writer) error {
	fmt.Fprintf(w, "// Class holds data for rating by points scored.\n")
	fmt.Fprintf(w, "type Class struct {\n")
	fmt.Fprintf(w, "    Threshold int       // minimum score needed for the rating\n")
	fmt.Fprintf(w, "    Message   []string\n")
	fmt.Fprintf(w, "}\n")
	fmt.Fprintf(w, "\n")
	fmt.Fprintf(w, "// generateClasses returns the initial state for classes\n")
	fmt.Fprintf(w, "func generateClasses() []*Class {\n")
	fmt.Fprintf(w, "    return []*Class{\n")
	for _, v := range y.Classes {
		fmt.Fprintf(w, "        &Class{\n")
		if v.Threshold != 0 {
			fmt.Fprintf(w, "            Threshold: %d,\n", v.Threshold)
		}
		if len(v.Message) != 0 {
			fmt.Fprintf(w, "            Message: []string{\n")
			for _, s := range strings.Split(v.Message, "\n") {
				fmt.Fprintf(w, "                %q,\n", s)
			}
			fmt.Fprintf(w, "            },\n")
		}
		fmt.Fprintf(w, "        },\n")
	}
	fmt.Fprintf(w, "    }\n")
	fmt.Fprintf(w, "}\n")
	fmt.Fprintf(w, "\n")

	return nil
}

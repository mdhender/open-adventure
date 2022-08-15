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

// OBITUARY is used to load the YAML data for obituaries
type OBITUARY struct {
	Query       string
	YesResponse string `yaml:"yes_response"`
}

// GenerateObituaries converts
func (y *YAML) GenerateObituaries(w io.Writer) error {
	_, _ = fmt.Fprintf(w, "// Obituary is data\n")
	_, _ = fmt.Fprintf(w, "type Obituary struct {\n")
	_, _ = fmt.Fprintf(w, "    Query       []string\n")
	_, _ = fmt.Fprintf(w, "    YesResponse []string\n")
	_, _ = fmt.Fprintf(w, "}\n\n")

	_, _ = fmt.Fprintf(w, "// generateObituaries returns the initial state of obituaries\n")
	_, _ = fmt.Fprintf(w, "func generateObituaries() []*Obituary {\n")
	_, _ = fmt.Fprintf(w, "    return []*Obituary{\n")
	for _, v := range y.Obituaries {
		_, _ = fmt.Fprintf(w, "        &Obituary{\n")
		if len(v.Query) != 0 {
			_, _ = fmt.Fprintf(w, "            Query: []string{\n")
			for _, s := range strings.Split(v.Query, "\n") {
				_, _ = fmt.Fprintf(w, "                %q,\n", s)
			}
			_, _ = fmt.Fprintf(w, "            },\n")
		}
		if len(v.YesResponse) != 0 {
			_, _ = fmt.Fprintf(w, "            YesResponse: []string{\n")
			for _, s := range strings.Split(v.YesResponse, "\n") {
				_, _ = fmt.Fprintf(w, "                %q,\n", s)
			}
			_, _ = fmt.Fprintf(w, "            },\n")
		}
		_, _ = fmt.Fprintf(w, "        },\n")
	}
	_, _ = fmt.Fprintf(w, "   }\n")
	_, _ = fmt.Fprintf(w, "}\n")
	_, _ = fmt.Fprintf(w, "\n")

	return nil
}

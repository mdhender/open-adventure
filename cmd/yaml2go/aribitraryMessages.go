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

// GenerateArbitraryMessages converts
func (y *YAML) GenerateArbitraryMessages(w io.Writer) error {
	fmt.Fprintf(w, "// ArbitraryMessages is an ordered map of messages\n")
	fmt.Fprintf(w, "// Use `Seq` to access sequentially or `Map` to access by key.\n")
	fmt.Fprintf(w, "// Note that no message string ends with a new-line!\n")
	fmt.Fprintf(w, "type ArbitraryMessages struct {\n")
	fmt.Fprintf(w, "    Seq []string           // sequential accessor\n")
	fmt.Fprintf(w, "    Map map[string]string  // tag accessor\n")
	fmt.Fprintf(w, "}\n")
	fmt.Fprintf(w, "\n")
	fmt.Fprintf(w, "// generateArbitraryMessages returns the initial arbitrary message state\n")
	fmt.Fprintf(w, "func generateArbitraryMessages() *ArbitraryMessages {\n")
	fmt.Fprintf(w, "    arbitraryMessages := &ArbitraryMessages{\n")
	fmt.Fprintf(w, "        Map: make(map[string]string),\n")
	fmt.Fprintf(w, "    }\n")
	// spew.Dump(y.ArbitraryMessages)
	for i, arbitraryMessageOMap := range y.ArbitraryMessages {
		fmt.Fprintf(w, "\n")
		if i == 0 {
			// grrr. NO_MESSAGE is !!null in the YAML, so
			// arbitraryMessageOMap is an empty map
			fmt.Fprintf(w, "    tag := %q\n", "NO_MESSAGE")
			fmt.Fprintf(w, "    arbitraryMessages.Map[tag] = \"\"\n")
			fmt.Fprintf(w, "    arbitraryMessages.Seq = append(arbitraryMessages.Seq, tag)\n")
		} else {
			for tag, arbitraryMessageText := range arbitraryMessageOMap {
				fmt.Fprintf(w, "    tag = %q\n", tag)
				fmt.Fprintf(w, "    arbitraryMessages.Seq = append(arbitraryMessages.Seq, tag)\n")
				fmt.Fprintf(w, "    arbitraryMessages.Map[tag] = %q\n", arbitraryMessageText)
			}
		}
	}
	fmt.Fprintf(w, "    return arbitraryMessages\n")
	fmt.Fprintf(w, "}\n")
	fmt.Fprintf(w, "\n")

	return nil
}

/*
 * Open Adventure - a clone of Open Adventure written in Go
 * Copyright (c) 2022 Michael D Henderson
 * SPDX-License-Identifier: BSD-2-clause
 */

package main

import "io/ioutil"

// LoadFile is a helper function to load a file from storage and return a slice of bytes or an error.
func LoadFile(filename string) ([]byte, error) {
	raw, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return raw, nil
}

// Open Adventure - a clone of Open Adventure written in Go
// Copyright (c) 2022 Michael D Henderson
// SPDX-License-Identifier: BSD-2-clause

package seeder

import (
	crand "crypto/rand"
	"encoding/binary"
)

// Seed returns the best seed we can get from the operating system.
// It is suitable to be used directly as a seed or to build a new source.
// For example:
//    seed, err := seeder.Seed()
// Then
//    rand.New(rand.NewSource(seed))
// Or
//    rand.Seed(seed)
// Please don't ignore errors from this function.
func Seed() (int64, error) {
	var seed [8]byte
	if _, err := crand.Read(seed[:]); err != nil {
		return 0, err
	}
	return int64(binary.LittleEndian.Uint64(seed[:])), nil
}

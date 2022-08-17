// Open Adventure - a clone of Open Adventure written in Go
// Copyright (c) 2022 Michael D Henderson
// SPDX-License-Identifier: BSD-2-clause

// Package lcg implements Knuth's generator per the original Open Adventure code.
package lcg

import "log"

// New returns an initialized generator
func New(seedval int32) *Type {
	/* LCG PRNG parameters tested against Knuth vol. 2. by the original authors */
	ng := &Type{a: 1093, c: 221587, m: 1048576}
	ng.setSeed(seedval)
	return ng
}

// GetCurrent returns the current value from the default Source.
func GetCurrent() int32 {
	return defaultSource.x
}

// GetNext returns the next value from the default Source.
func GetNext() int32 {
	return defaultSource.getNext()
}

// Intn returns, as an int, a non-negative pseudo-random number in [0,n) from the default Source.
func Intn(n int32) int32 {
	return defaultSource.intn(n)
}

// SetSeed sets the seed for the default source.
func SetSeed(seedval int32) {
	defaultSource.setSeed(seedval)
}

// ZZWord returns the Z`ZZZ word.
func ZZWord() string {
	return defaultSource.zzword
}

// Type implements an LCG for our PRNG.
type Type struct {
	x       int32
	a, c, m int32
	zzword  string
}

// generateZZWord generates a Z`ZZZ word.
func (ng *Type) generateZZWord() string {
	// the original generated the second letter then changed it to a tic.
	// if we don't duplicate, then the number sequence will be off by one.
	// that causes the tests to fail.
	a, b, c, d, e := ng.randomUpperAlpha(), ng.randomUpperAlpha(), ng.randomUpperAlpha(), ng.randomUpperAlpha(), ng.randomUpperAlpha()
	b = "'"
	return a + b + c + d + e
}

// getNext implements get_next_lcg_value.
//
// int32_t get_next_lcg_value (void) at misc.c:678
// Return the LCG's current value, and then iterate it.
func (ng *Type) getNext() (x int32) {
	x = ng.x
	ng.x = (ng.a*ng.x + ng.c) % ng.m
	return x
}

// intn returns, as an int, a non-negative pseudo-random number in [0,n) from the default Source.
func (ng *Type) intn(n int32) int32 {
	return n * ng.getNext() / ng.m
}

// randomUpperAlpha returns a pseudo-random letter
func (ng *Type) randomUpperAlpha() string {
	set := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	return string(set[ng.intn(int32(len(set)))])
}

// setSeed sets the seed.
func (ng *Type) setSeed(seedval int32) {
	ng.x = seedval % ng.m
	if ng.x < 0 {
		ng.x = ng.m + ng.x
		log.Printf("seed2: %d m: %d x: %d\n", seedval, ng.m, ng.x)
	}

	// once seed is set, we need to generate the Z`ZZZ word
	ng.zzword = ng.generateZZWord()
}

// defaultSource is the default generator.
var defaultSource *Type

func init() {
	defaultSource = New(0)
}

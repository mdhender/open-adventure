// Open Adventure - a clone of Open Adventure written in Go
// Copyright (c) 2022 Michael D Henderson
// SPDX-License-Identifier: BSD-2-clause

package main

import (
	"github.com/mdhender/open-adventure/cmd"
	"github.com/mdhender/open-adventure/internal/seeder"
	"log"
	"math/rand"
	"time"
)

func main() {
	defer func(started time.Time) {
		elapsed := time.Now().Sub(started)
		log.Printf("adventure: total time %v\n", elapsed)
	}(time.Now())

	// default log format to UTC
	log.SetFlags(log.Ldate | log.Ltime | log.LUTC)

	// seed the default PRNG source.
	if seed, err := seeder.Seed(); err != nil {
		log.Fatalln(err)
	} else {
		rand.Seed(seed)
	}

	// run the command as given
	cmd.Execute()
}

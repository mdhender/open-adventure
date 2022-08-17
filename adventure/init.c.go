// Open Adventure - a clone of Open Adventure written in Go
// Copyright (c) 2022 Michael D Henderson
// SPDX-License-Identifier: BSD-2-clause

package adventure

import (
	"io"
)

func initialise(w io.Writer, settings *settings_t) *game_t {
	game := &game_t{}

	if settings.oldstyle {
		game.printf("Initialising...\n")
	}

	var seedval int32 = 0xBEEF
	game.set_seed(seedval)

	game.dloc[1] = LOC_KINGHALL
	game.dloc[2] = LOC_WESTBANK
	game.dloc[3] = LOC_Y2
	game.dloc[4] = LOC_ALIKE3
	game.dloc[5] = LOC_COMPLEX

	/*  Sixth dwarf is special (the pirate).  He always starts at his
	 *  chest's eventual location inside the maze. This loc is saved
	 *  in chloc for ref. The dead end in the other maze has its
	 *  loc stored in chloc2. */
	game.dloc[6] = LOC_MAZEEND12
	game.chloc = LOC_MAZEEND12
	game.chloc2 = LOC_DEADEND13

	game.abbnum = 5
	game.clock1 = WARNTIME
	game.clock2 = FLASHTIME
	game.newloc = LOC_START
	game.loc = LOC_START
	game.limit = GAMELIMIT
	game.foobar = WORD_EMPTY

	for i := 1; i <= NOBJECTS; i++ {
		game.place[i] = LOC_NOWHERE
	}

	return game
}

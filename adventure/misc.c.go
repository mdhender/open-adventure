// Open Adventure - a clone of Open Adventure written in Go
// Copyright (c) 2022 Michael D Henderson
// SPDX-License-Identifier: BSD-2-clause

package adventure

func (game *game_t) get_next_lcg_value() (x int32) {
	old_x := game.lcg_x
	game.lcg_x = (LCG_A*game.lcg_x + LCG_C) % LCG_M
	return old_x
}

/* Return a random integer from [0, range). */
func (game *game_t) randrange(r int32) int32 {
	return r * game.get_next_lcg_value() / LCG_M
}

func (game *game_t) set_seed(seedval int32) {
	game.lcg_x = seedval % LCG_M
	if game.lcg_x < 0 {
		game.lcg_x = LCG_M + game.lcg_x
	}
	// once seed is set, we need to generate the Z`ZZZ word
	for i := 0; i < 5; i++ {
		game.zzword[i] = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"[game.randrange(26)]
	}
	game.zzword[1] = '\'' // force second char to apostrophe
	game.zzword[5] = 0
}

// Open Adventure - a clone of Open Adventure written in Go
// Copyright (c) 2022 Michael D Henderson
// SPDX-License-Identifier: BSD-2-clause

package adventure

import (
	"encoding/json"
	"github.com/mdhender/open-adventure/state"
	"os"
)

func Save(adv *state.Adventure, saveFile string) error {
	if b, err := json.MarshalIndent(adv, "", "  "); err != nil {
		return err
	} else if err = os.WriteFile(saveFile, b, 0666); err != nil {
		return err
	}
	return nil
}

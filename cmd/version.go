// Open Adventure - a clone of Open Adventure written in Go
// Copyright (c) 2022 Michael D Henderson
// SPDX-License-Identifier: BSD-2-clause

package cmd

import (
	"fmt"
	"github.com/mdhender/open-adventure/state"
	"github.com/spf13/cobra"
	"log"
)

var cmdVersion = &cobra.Command{
	Use:   "version",
	Short: "show application version",
	Long:  `Show application version.`,
	Run: func(cmd *cobra.Command, args []string) {
		adv, err := state.NewAdventure()
		if err != nil {
			log.Fatal(err)
		}

		// format the version per https://semver.org/ rules
		version := fmt.Sprintf("%d.%d.%d", adv.Version.Major, adv.Version.Minor, adv.Version.Patch)
		if adv.Version.Suffix != "" {
			version += "+" + adv.Version.Suffix
		}

		fmt.Printf("%s\n", version)
	},
}

func init() {
	cmdBase.AddCommand(cmdVersion)
}

// Open Adventure - a clone of Open Adventure written in Go
// Copyright (c) 2022 Michael D Henderson
// SPDX-License-Identifier: BSD-2-clause

package cmd

import (
	"github.com/mdhender/open-adventure/adventure"
	"github.com/spf13/cobra"
	"log"
)

var cmdState = &cobra.Command{
	Use:   "state",
	Short: "save state to a JSON file",
	Long:  `Create a new adventure and write its state to a persistent store.`,
	Run: func(cmd *cobra.Command, args []string) {
		adv, output, err := adventure.Execute(nil, "")
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("%s\n", output)

		stateFile := "D:\\open-adventure\\testdata\\state.json"
		if err = adventure.Save(adv, stateFile); err != nil {
			log.Fatal(err)
		}

		log.Printf("state: saved to %q\n", stateFile)
	},
}

func init() {
	cmdBase.AddCommand(cmdState)
}

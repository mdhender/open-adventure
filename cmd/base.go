// Open Adventure - a clone of Open Adventure written in Go
// Copyright (c) 2022 Michael D Henderson
// SPDX-License-Identifier: BSD-2-clause

package cmd

import (
	"fmt"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
)

var globalBase struct {
	TestFlag    bool
	VerboseFlag bool
	ConfigFile  string // configuration file from command line flag

	envPrefix  string // value to prepend when converting flags to env variables
	cfgName    string // default configuration file name
	homeFolder string // derived path to home directory
}

// cmdBase represents the base command when called without any subcommands
var cmdBase = &cobra.Command{
	Use:   "adventure",
	Short: "Open Adventure game server",
	Long:  `adventure is the game server for Open Adventure.`,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		// now bind viper and cobra configuration since this hook always runs early
		return bindConfig(cmd)
	},
	Run: func(cmd *cobra.Command, args []string) {
		log.Printf("env: %-30s == %q\n", "HOME", globalBase.homeFolder)
		log.Printf("env: %-30s == %q\n", "ADVENTURE_CONFIG", viper.ConfigFileUsed())
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(cmdBase.Execute())
}

func init() {
	// set the env and config
	globalBase.envPrefix, globalBase.cfgName = "ADVENTURE", ".adventure"
	// find home directory
	if home, err := homedir.Dir(); err == nil {
		globalBase.homeFolder = home
	}

	cmdBase.PersistentFlags().StringVar(&globalBase.ConfigFile, "config", "", fmt.Sprintf("config file (default is ~/%s.json)", globalBase.cfgName))
	cmdBase.PersistentFlags().BoolVar(&globalBase.TestFlag, "test", false, "test mode")
	cmdBase.PersistentFlags().BoolVar(&globalBase.VerboseFlag, "verbose", false, "verbose mode")
}

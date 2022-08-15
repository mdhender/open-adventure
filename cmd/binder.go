// Open Adventure - a clone of Open Adventure written in Go
// Copyright (c) 2022 Michael D Henderson
// SPDX-License-Identifier: BSD-2-clause

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"log"
	"path/filepath"
	"strings"
)

// logic for binding viper and cobra taken from https://carolynvanslyck.com/blog/2020/08/sting-of-the-viper/

// bindConfig reads in config file and ENV variables if set.
func bindConfig(cmd *cobra.Command) error {
	if globalBase.ConfigFile == "" { // use default location of ~/.adventure.json
		viper.AddConfigPath(globalBase.homeFolder)
		viper.SetConfigName(globalBase.cfgName)
		viper.SetConfigType("json")
		globalBase.ConfigFile = filepath.Clean(filepath.Join(globalBase.homeFolder, globalBase.cfgName+".json"))
	} else { // Use config file from the flag.
		viper.SetConfigFile(globalBase.ConfigFile)
	}

	// Try to read the config file. Ignore file-not-found errors.
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return err
		}
	} else {
		if globalBase.VerboseFlag {
			log.Printf("[viper] using config file: %q\n", viper.ConfigFileUsed())
		}
		globalBase.ConfigFile = filepath.Clean(viper.ConfigFileUsed())
	}

	// read in environment variables that match
	viper.SetEnvPrefix(globalBase.envPrefix)
	viper.AutomaticEnv()

	// bind the current command's flags to viper
	cmd.Flags().VisitAll(func(f *pflag.Flag) {
		// Environment variables can't have dashes in them, so bind them to their equivalent
		// keys with underscores, e.g. --favorite-color to ADVENTURE_FAVORITE_COLOR
		if strings.Contains(f.Name, "-") {
			envVarSuffix := strings.ToUpper(strings.ReplaceAll(f.Name, "-", "_"))
			_ = viper.BindEnv(f.Name, fmt.Sprintf("%s_%s", globalBase.envPrefix, envVarSuffix))
		}

		// Apply the viper config value to the flag when the flag is not set and viper has a value
		if !f.Changed && viper.IsSet(f.Name) {
			val := viper.Get(f.Name)
			_ = cmd.Flags().Set(f.Name, fmt.Sprintf("%v", val))
		}
	})

	return nil
}

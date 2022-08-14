/*
 * Open Adventure - a clone of Open Adventure written in Go
 * Copyright (c) 2022 Michael D Henderson
 * SPDX-License-Identifier: BSD-2-clause
 */

package main

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"reflect"
	"regexp"
	"strings"
)

// YAML is
type YAML struct {
	Actions           []map[string]*ACTION
	ArbitraryMessages []map[string]string `yaml:"arbitrary_messages"`
	Classes           []*CLASS
	Hints             []map[string]*HINT
	Locations         []map[string]*LOCATION
	Motions           []map[string]*MOTION
	Obituaries        []*OBITUARY
	Objects           []map[string]*OBJECT
	TurnThresholds    []*TURNTHRESHOLD `yaml:"turn_thresholds"`
}

func loadYAML(raw []byte) (*YAML, error) {
	y := &YAML{}
	if err := yaml.Unmarshal(raw, y); err != nil {
		return nil, fmt.Errorf("unable to unmarshal yaml: %w", err)
	}

	// there are some subtleties with the input that we have to work with.
	// one is boolean values that default to true. another is that there's
	// a field which can be either a string or slice of strings. the last
	// is that some hints have embedded newlines. we must convert
	// those to the escaped-newlines that Go expects.

	// scan actions.oldstyle replacing the default value with true.
	// (we can identify default values because the YAML reader sets them to nil.)
	for _, actionMap := range y.Actions {
		for _, action := range actionMap {
			action.OldStyle = true
			if action.RawOldStyle != nil {
				action.OldStyle = *action.RawOldStyle
			}
		}
	}

	// scan motions.oldstyle replacing the default value with true.
	// (we can identify default values because the YAML reader sets them to nil.)
	for _, motionMap := range y.Motions {
		for _, motion := range motionMap {
			motion.OldStyle = true
			if motion.RawOldStyle != nil {
				motion.OldStyle = *motion.RawOldStyle
			}
		}
	}

	// apologies for this, but object.locations is the field that can
	// be either a string or an array of strings in the YAML.
	for _, objectMap := range y.Objects {
		for _, object := range objectMap {
			if object.RawLocations != nil {
				kind := reflect.TypeOf(object.RawLocations).Kind()
				switch kind {
				case reflect.Slice:
					val := reflect.ValueOf(object.RawLocations)
					for x := 0; x < val.Len(); x++ {
						i := val.Index(x).Interface()
						kind := reflect.TypeOf(i).Kind()
						switch kind {
						case reflect.String:
							object.Locations = append(object.Locations, reflect.ValueOf(i).String())
						default:
							return nil, fmt.Errorf("object: locations: unknown %q", kind)
						}
					}
				case reflect.String:
					s := reflect.ValueOf(object.RawLocations).String()
					if len(s) != 0 {
						object.Locations = append(object.Locations, s)
					}
				default:
					return nil, fmt.Errorf("object: locations: unknown %q", kind)
				}
			}
		}
	}

	// regex matches a literal backslash n in the input.
	var reLiteralNewline = regexp.MustCompile(`\\n`)

	// loop through all locations in the map, then all tags in the location, then all hints.
	for _, locationMap := range y.Locations {
		for _, location := range locationMap {
			for _, hint := range location.Hints {
				// replace every backslash n in the hints with an escaped newline.
				for _, s := range strings.Split(reLiteralNewline.ReplaceAllLiteralString(hint.Hint, "\n"), "\n") {
					hint.Hints = append(hint.Hints, s)
				}
			}
		}
	}

	return y, nil
}

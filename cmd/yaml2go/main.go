/*
 * Open Adventure - a clone of Open Adventure written in Go
 * Copyright (c) 2022 Michael D Henderson
 * SPDX-License-Identifier: BSD-2-clause
 */

// Package main implements a translator from YAML to Go.
package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func main() {
	input := "adventure.yaml"
	fmt.Printf("info: %-30s == %q\n", "INPUT", input)

	outputPath := "./state/"
	fmt.Printf("info: %-30s == %q\n", "OUTPUT_PATH", outputPath)

	if err := run(input, outputPath); err != nil {
		fmt.Printf("%+v\n", err)
		os.Exit(2)
	}

	fmt.Println("info: completed successfully")
	os.Exit(0)
}

func run(input, outputPath string) error {
	raw, err := LoadFile(input)
	if err != nil {
		return fmt.Errorf("unable to load input: %w", err)
	}

	y, err := loadYAML(raw)
	if err != nil {
		return fmt.Errorf("unable to parse input: %w", err)
	}

	d := &dungeon{
		// use header format from https://github.com/golang/go/issues/13560#issuecomment-288457920
		DoNotEditHeader: fmt.Sprintf("Code generated by yaml2go on %s; DO NOT EDIT.", time.Now().UTC().Format(time.RFC3339)),
	}

	for i, locationMap := range y.Locations {
		for tag := range locationMap {
			d.LocationsRefs = append(d.LocationsRefs, &location_ref{
				Seq:  i,
				Enum: tag,
			})
		}
	}
	d.NLocations = len(d.LocationsRefs)

	for _, t := range []string{"advent", "dungeon"} {
		d.TemplateName = t
		b := &bytes.Buffer{}
		if err = y.GenerateTemplate(d, b); err != nil {
			return err
		} else if err = os.WriteFile(filepath.Join("adventure", fmt.Sprintf("%s.h.go", t)), b.Bytes(), 0666); err != nil {
			return err
		}
	}

	// use header format from https://github.com/golang/go/issues/13560#issuecomment-288457920
	header := fmt.Sprintf("// Code generated by yaml2go on %s; DO NOT EDIT.", time.Now().UTC().Format(time.RFC3339))

	for _, t := range []string{"adventure", "actions", "arbitraryMessages", "classes", "locations", "motions", "obituaries", "objects", "turnThresholds"} {
		f, err := os.Create(filepath.Join(outputPath, fmt.Sprintf("%s.go", t)))
		if err != nil {
			return fmt.Errorf("creating actions file: %w", err)
		}
		w := bufio.NewWriter(f)
		_, _ = fmt.Fprintln(w, header)
		_, _ = fmt.Fprintln(w, "")
		_, _ = fmt.Fprintln(w, "// Open Adventure - a clone of Open Adventure written in Go")
		_, _ = fmt.Fprintln(w, "// Copyright (c) 2022 Michael D Henderson")
		_, _ = fmt.Fprintln(w, "// SPDX-License-Identifier: BSD-2-clause")
		_, _ = fmt.Fprintln(w, "")
		_, _ = fmt.Fprintln(w, "package state")
		_, _ = fmt.Fprintln(w, "")

		switch t {
		case "adventure":
			err = y.GenerateAdventure(w)
		case "actions":
			err = y.GenerateActions(w)
		case "arbitraryMessages":
			err = y.GenerateArbitraryMessages(w)
		case "classes":
			err = y.GenerateClasses(w)
		case "locations":
			err = y.GenerateLocations(w)
		case "motions":
			err = y.GenerateMotions(w)
		case "obituaries":
			err = y.GenerateObituaries(w)
		case "objects":
			err = y.GenerateObjects(w)
		case "turnThresholds":
			err = y.GenerateTurnThresholds(w)
		}
		if err != nil {
			return fmt.Errorf("generating %s code: %w", t, err)
		}

		if err = w.Flush(); err != nil {
			return err
		} else if err = f.Close(); err != nil {
			return err
		}
	}

	return nil
}

// Open Adventure - a clone of Open Adventure written in Go
// Copyright (c) 2022 Michael D Henderson
// SPDX-License-Identifier: BSD-2-clause

package main

import (
	"fmt"
	"io"
	"path/filepath"
	"text/template"
)

type dungeon struct {
	TemplateName    string
	DoNotEditHeader string
	LocationsRefs   []*location_ref
	NLocations      int
	NObjects        int
	NHints          int
}

type location_ref struct {
	Seq  int
	Enum string
}

func (y *YAML) GenerateTemplate(d *dungeon, w io.Writer) error {
	t, err := template.ParseFiles(filepath.Join("cmd", "yaml2go", fmt.Sprintf("%s.go.tmpl", d.TemplateName)))
	if err != nil {
		return err
	}
	return t.Execute(w, d)
}

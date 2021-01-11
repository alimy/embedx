// Copyright 2021 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package embedx

import (
	"bytes"
	"embed"
	"strings"
	"testing"
	"text/template"
)

func TestParseFS(t *testing.T) {
	//go:embed testdata
	var content embed.FS

	embedFS := ChangeRoot(content, "testdata")
	tmpl, err := ParseFS(embedFS, "templates/*.tmpl", "templates/b/*.tmpl")
	if err != nil {
		t.Errorf("parse embed fs to template error: %s", err)
	}

	for _, ctx := range []struct {
		Name string
	}{
		{"templates/a.tmpl"},
		{"templates/b/c.tmpl"},
	} {
		bs := &bytes.Buffer{}
		if err := tmpl.ExecuteTemplate(bs, ctx.Name, ctx); err != nil {
			t.Errorf("ExcuteTemplate(%s) error: %s", ctx.Name, err)
		}
		rs := strings.Split(bs.String(), "=")
		if len(rs) != 2 {
			t.Errorf("result split is not 2: %s", bs.String())
		}
		lh, rh := strings.Trim(rs[0], " "), strings.Trim(rs[1], " ")
		if lh != rh {
			t.Errorf("result of rendered is not correct: %s", bs.String())
		}
	}
}

func TestParseWith(t *testing.T) {
	//go:embed testdata
	var content embed.FS

	embedFS := ChangeRoot(content, "testdata")
	funcTmpl := template.New("embedx").Funcs(template.FuncMap{
		"notEmptyStr": notEmptyStr,
	})
	tmpl, err := ParseWith(funcTmpl, embedFS, "templates/*.tmpl", "templates/b/*.tmpl", "templates/d/*.tmpl")
	if err != nil {
		t.Errorf("parse embed fs to template error: %s", err)
	}

	for _, ctx := range []struct {
		Name string
	}{
		{"templates/a.tmpl"},
		{"templates/b/c.tmpl"},
		{"templates/d/e.tmpl"},
		{"templates/d/f.tmpl"},
	} {
		bs := &bytes.Buffer{}
		if err := tmpl.ExecuteTemplate(bs, ctx.Name, ctx); err != nil {
			t.Errorf("ExcuteTemplate(%s) error: %s", ctx.Name, err)
		}
		rs := strings.Split(bs.String(), "=")
		if len(rs) != 2 {
			t.Errorf("result split is not 2: %s", bs.String())
		}
		lh, rh := strings.Trim(rs[0], " "), strings.Trim(rs[1], " ")
		if lh != rh {
			t.Errorf("result of rendered is not correct: %s", bs.String())
		}
	}
}

func notEmptyStr(s string) bool {
	return s != ""
}

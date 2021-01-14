// Copyright 2021 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package embedx

import (
	"path"
	"testing"
)

func TestNaming(t *testing.T) {
	for origin, expect := range map[string]string{
		"templates/a.tmpl":   "templates/a.tmpl",
		"templates/b/c.tmpl": "templates/b/c.tmpl",
		"templates/d/e.tmpl": "templates/d/e.tmpl",
		"templates/d/f.tmpl": "templates/d/f.tmpl",
	} {
		if name := Naming(origin); name != expect {
			t.Errorf("expect: %s got %s", expect, name)
		}
	}
}

func TestRegisterNamer(t *testing.T) {
	results := make(map[string]string)
	RegisterNamer(NamerFunc(func(filepath string) string {
		ext := path.Ext(filepath)
		results[filepath] = filepath[:len(filepath)-len(ext)]
		return filepath
	}))
	for origin, expect := range map[string]string{
		"templates/a.tmpl":   "templates/a",
		"templates/b/c.tmpl": "templates/b/c",
		"templates/d/e.tmpl": "templates/d/e",
		"templates/d/f.tmpl": "templates/d/f",
	} {
		name := Naming(origin)
		if value, exist := results[name]; exist && value != expect {
			t.Errorf("expect: %s got %s", expect, value)
		}
	}
}

// Copyright 2021 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package embedx

import (
	"fmt"
	"io/fs"
	"text/template"
)

// ParseFS creates a new text/Template and parses the template definitions from
// the files identified by the pattern. The files are matched according to the
// semantics of filepath.Match, and the pattern must match at least one file.
// The returned template will have the (base) name and (parsed) contents of the
// first file matched by the pattern.
func ParseFS(fsys fs.FS, patterns ...string) (*template.Template, error) {
	return ParseWith(nil, fsys, patterns...)
}

// ParseWith like ParseFS but need provide a *template.Template instance as parameter.
func ParseWith(t *template.Template, fsys fs.FS, patterns ...string) (*template.Template, error) {
	var filenames []string

	for _, pattern := range patterns {
		list, err := fs.Glob(fsys, pattern)
		if err != nil {
			return nil, err
		}
		if len(list) == 0 {
			return nil, fmt.Errorf("template: pattern matches no files: %#q", pattern)
		}
		filenames = append(filenames, list...)
	}

	if len(filenames) == 0 {
		// Not really a problem, but be consistent.
		return nil, fmt.Errorf("text/template: no files named in call to ParseFiles")
	}

	for _, name := range filenames {
		b, err := fs.ReadFile(fsys, name)
		if err != nil {
			return nil, err
		}
		s := string(b)
		// First template becomes return value if not already defined,
		// and we use that one for subsequent New calls to associate
		// all the templates together.
		tmplName := namer.Naming(name)
		if t == nil {
			t, err = template.New(tmplName).Parse(s)
			if err != nil {
				return nil, err
			}
			continue
		}
		tmpl, err := t.Clone()
		if err != nil {
			return nil, err
		}
		if _, err = tmpl.Parse(s); err == nil {
			_, err = t.AddParseTree(tmplName, tmpl.Tree)
		}
		if err != nil {
			return nil, err
		}
	}
	return t, nil
}

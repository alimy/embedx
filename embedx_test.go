// Copyright 2021 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package embedx

import "testing"

func TestAmendPath(t *testing.T) {
	for _, path := range []struct {
		origin   string
		expect   string
		isDotDir bool
	}{
		{" /", "", true},
		{"", "", true},
		{".", ".", true},
		{" abc ", "abc", false},
		{"/abc", "abc", false},
		{"/abc/", "abc", false},
		{"//abc//", "abc", false},
		{`\abc\`, "abc", false},
		{` \abc \ `, "abc", false},
		{` \abc\abc\ `, "abc/abc", false},
		{`\ /abc/ `, "abc", false},
	} {
		name, isDotDir := amendPath(path.origin)
		if name != path.expect || isDotDir != path.isDotDir {
			t.Errorf("path=>expect:%s got: %s, isDot=> expect: %t got %t", path.expect, name, path.isDotDir, isDotDir)
		}
	}
}

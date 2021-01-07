// Copyright 2021 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package embedx

import "testing"

func TestName(t *testing.T) {
	for origin, expect := range map[string]string{
		" abc ":    "abc",
		"/abc":     "abc",
		"/abc/":    "abc",
		"//abc//":  "abc",
		`\abc\`:    "abc",
		` \abc \ `: "abc",
		`\ /abc/ `: "abc",
	} {
		a := &args{category: categoryNone}
		ChangeRoot(origin).apply(a)
		if a.category != categoryRootFS || a.rootDir != expect {
			t.Errorf("category=>expect:%d got: %d, rootdir=> expect: %s got %s", categoryRootFS, a.category, expect, a.rootDir)
		}

		a = &args{category: categoryNone}
		AttachRoot(origin).apply(a)
		if a.category != categoryAttachFS || a.attachDir != expect {
			t.Errorf("category=>expect:%d got: %d, rootdir=> expect: %s got %s", categoryAttachFS, a.category, expect, a.attachDir)
		}

		a = &args{category: categoryNone}
		AttachRoot(origin).apply(a)
		ChangeRoot(origin).apply(a)
		if a.category != categoryBundleFS || a.rootDir != expect || a.attachDir != expect {
			t.Errorf("category=>expect:%d got: %d, rootdir=> expect: %s got %s, attachdir=> expect: %s got %s", categoryBundleFS, a.category, expect, a.rootDir, expect, a.attachDir)
		}

		a = &args{category: categoryNone}
		for range [6]int{} {
			AttachRoot(origin).apply(a)
			ChangeRoot(origin).apply(a)
		}

		if a.category != categoryBundleFS || a.rootDir != expect || a.attachDir != expect {
			t.Errorf("category=>expect:%d got: %d, rootdir=> expect: %s got %s, attachdir=> expect: %s got %s", categoryBundleFS, a.category, expect, a.rootDir, expect, a.attachDir)
		}
	}
}

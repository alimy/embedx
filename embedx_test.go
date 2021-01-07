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
		ChangeRootOpt(origin).apply(a)
		if a.category != categoryRootFS || a.rootDir != expect {
			t.Errorf("category=>expect:%d got: %d, rootdir=> expect: %s got %s", categoryRootFS, a.category, expect, a.rootDir)
		}

		a = &args{category: categoryNone}
		AttachRootOpt(origin).apply(a)
		if a.category != categoryAttachFS || a.attachDir != expect {
			t.Errorf("category=>expect:%d got: %d, rootdir=> expect: %s got %s", categoryAttachFS, a.category, expect, a.attachDir)
		}

		a = &args{category: categoryNone}
		AttachRootOpt(origin).apply(a)
		ChangeRootOpt(origin).apply(a)
		if a.category != categoryBundleFS || a.rootDir != expect || a.attachDir != expect {
			t.Errorf("category=>expect:%d got: %d, rootdir=> expect: %s got %s, attachdir=> expect: %s got %s", categoryBundleFS, a.category, expect, a.rootDir, expect, a.attachDir)
		}

		a = &args{category: categoryNone}
		for range [6]int{} {
			AttachRootOpt(origin).apply(a)
			ChangeRootOpt(origin).apply(a)
		}

		if a.category != categoryBundleFS || a.rootDir != expect || a.attachDir != expect {
			t.Errorf("category=>expect:%d got: %d, rootdir=> expect: %s got %s, attachdir=> expect: %s got %s", categoryBundleFS, a.category, expect, a.rootDir, expect, a.attachDir)
		}
	}
}

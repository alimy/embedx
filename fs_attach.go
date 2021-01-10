// Copyright 2021 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package embedx

import (
	"io/fs"
	"strings"
)

type attachEmbedFS struct {
	EmbedFS
	root      string
	rootSlash string
}

// Open opens the named file for reading and returns it as an fs.File.
func (f *attachEmbedFS) Open(name string) (fs.File, error) {
	return f.EmbedFS.Open(f.path(name))
}

// ReadDir reads and returns the entire named directory.
func (f *attachEmbedFS) ReadDir(name string) ([]fs.DirEntry, error) {
	return f.EmbedFS.ReadDir(f.path(name))
}

// ReadFile reads and returns the content of the named file.
func (f *attachEmbedFS) ReadFile(name string) ([]byte, error) {
	return f.EmbedFS.ReadFile(f.path(name))
}

func (f *attachEmbedFS) path(name string) string {
	if name == f.root {
		return "."
	}
	return strings.TrimPrefix(name, f.rootSlash)
}

func newAttachEmbedFS(fs EmbedFS, root string) EmbedFS {
	return &attachEmbedFS{
		EmbedFS:   fs,
		root:      root,
		rootSlash: root + "/",
	}
}

// Copyright 2021 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package embedx

import (
	"embed"
	"io/fs"
	"strings"
)

type rootEmbedFS struct {
	*embed.FS
	root string
}

// Open opens the named file for reading and returns it as an fs.File.
func (f *rootEmbedFS) Open(name string) (fs.File, error) {
	return f.FS.Open(f.path(name))
}

// ReadDir reads and returns the entire named directory.
func (f *rootEmbedFS) ReadDir(name string) ([]fs.DirEntry, error) {
	return f.FS.ReadDir(f.path(name))
}

// ReadFile reads and returns the content of the named file.
func (f *rootEmbedFS) ReadFile(name string) ([]byte, error) {
	return f.FS.ReadFile(f.path(name))
}

func (f *rootEmbedFS) path(name string) string {
	if name == "." {
		return f.root
	}
	return strings.Join([]string{f.root, name}, "/")
}

func newRootEmbedFS(content *embed.FS, root string) EmbedFS {
	return &rootEmbedFS{
		FS:   content,
		root: root,
	}
}

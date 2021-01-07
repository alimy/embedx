// Copyright 2021 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package embedx

import (
	"embed"
	"io/fs"
	"strings"
)

type attachEmbedFS struct {
	*embed.FS
	attachDir      string
	attachDirSlash string
}

// Open opens the named file for reading and returns it as an fs.File.
func (f *attachEmbedFS) Open(name string) (fs.File, error) {
	return f.FS.Open(f.path(name))
}

// ReadDir reads and returns the entire named directory.
func (f *attachEmbedFS) ReadDir(name string) ([]fs.DirEntry, error) {
	return f.FS.ReadDir(f.path(name))
}

// ReadFile reads and returns the content of the named file.
func (f *attachEmbedFS) ReadFile(name string) ([]byte, error) {
	return f.FS.ReadFile(f.path(name))
}

func (f *attachEmbedFS) path(name string) string {
	if name == f.attachDir {
		return "."
	}
	return strings.TrimPrefix(name, f.attachDirSlash)
}

func newAttachEmbedFS(content *embed.FS, attachDir string) EmbedFS {
	return &attachEmbedFS{
		FS:             content,
		attachDir:      attachDir,
		attachDirSlash: attachDir + "/",
	}
}

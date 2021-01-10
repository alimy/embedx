// Copyright 2021 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package embedx

import (
	"io/fs"
	"strings"
)

// EmbedFS embed.FS public method re-defined as  interface
type EmbedFS interface {
	fs.FS
	ReadDir(name string) ([]fs.DirEntry, error)
	ReadFile(name string) ([]byte, error)
}

// AttachRoot attach to a virtual root directory like mount to a new directory.
func AttachRoot(fs EmbedFS, root string) EmbedFS {
	if name, isDotDir := amendPath(root); isDotDir {
		return fs
	} else {
		return newAttachEmbedFS(fs, name)
	}

}

// ChangeRoot change to a new directory like cd cmd in shell but not check whether
// this directory is exist.
func ChangeRoot(fs EmbedFS, root string) EmbedFS {
	if name, isDotDir := amendPath(root); isDotDir {
		return fs
	} else {
		return newRootEmbedFS(fs, name)
	}
}

// note: will not process path like ".."
func amendPath(name string) (string, bool) {
	name = strings.Trim(strings.ReplaceAll(name, `\`, "/"), " /")
	isDotDir := name == "" || name == "."
	return name, isDotDir
}

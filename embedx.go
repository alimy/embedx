// Copyright 2021 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package embedx

import (
	"embed"
	"io/fs"
	"strings"
)

const (
	categoryNone     int8 = 0
	categoryAttachFS int8 = 1
	categoryRootFS   int8 = 2
	categoryBundleFS      = categoryAttachFS | categoryRootFS
)

// EmbedFS embed.FS public method re-defined as  interface
type EmbedFS interface {
	fs.FS
	ReadDir(name string) ([]fs.DirEntry, error)
	ReadFile(name string) ([]byte, error)
}

type args struct {
	category  int8
	attachDir string
	rootDir   string
}

type option interface {
	apply(*args)
}

type argFunc func(*args)

func (f argFunc) apply(arg *args) {
	f(arg)
}

// AttachRoot setup attached to root directory name.
func AttachRoot(rootDir string) option {
	rootDir = strings.Trim(rootDir, `\ /`)
	return argFunc(func(a *args) {
		a.category |= categoryAttachFS
		a.attachDir = rootDir
	})
}

// ChangeRoot setup changed to root directory name.
func ChangeRoot(rootDir string) option {
	rootDir = strings.Trim(rootDir, `\ /`)
	return argFunc(func(a *args) {
		a.category |= categoryRootFS
		a.rootDir = rootDir
	})
}

// NewFileSystem make an EmbedFS instance that contain embed.FS resource.
func NewFileSystem(content *embed.FS, opts ...option) EmbedFS {
	a := &args{category: categoryNone}
	for _, opt := range opts {
		opt.apply(a)
	}
	switch {
	case a.category == categoryRootFS && a.rootDir != "":
		return newRootEmbedFS(content, a.rootDir)
	case a.category == categoryAttachFS && a.attachDir != "":
		return newAttachEmbedFS(content, a.attachDir)
	case a.category == categoryBundleFS && a.rootDir != "" && a.attachDir != "":
		return newBundleEmbedFS(content, a.rootDir, a.attachDir)
	default:
		return content
	}
}

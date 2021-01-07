package embedx

import (
	"embed"
	"net/http"
	"strings"
)

const (
	categoryNone     = 0
	categoryAttachFS = 1
	categoryRootFS   = 2
	categoryBundleFS = categoryAttachFS | categoryRootFS
)

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

// AttachRootOpt setup attached to root directory name
func AttachRootOpt(rootDir string) option {
	rootDir = strings.Trim(rootDir, `\ /`)
	return argFunc(func(a *args) {
		a.category |= categoryAttachFS
		a.attachDir = rootDir
	})
}

// ChangeRootOpt setup changed to root directory name
func ChangeRootOpt(rootDir string) option {
	rootDir = strings.Trim(rootDir, `\ /`)
	return argFunc(func(a *args) {
		a.category |= categoryRootFS
		a.rootDir = rootDir
	})
}

// NewFileSystem make a http.FileSystem instance that contain embed.FS resource.
func NewFileSystem(content *embed.FS, opts ...option) http.FileSystem {
	a := &args{category: categoryNone}
	for _, opt := range opts {
		opt.apply(a)
	}
	switch {
	case a.category == categoryRootFS && a.rootDir != "":
		return newRootEmbedFS(content, a.rootDir)
	case a.category == categoryAttachFS && a.attachDir != "":
		return newAttachEmbedFS(content, a.rootDir)
	case a.category == categoryBundleFS && a.rootDir != "" && a.attachDir != "":
		return newBundleEmbedFS(content, a.rootDir, a.attachDir)
	default:
		return http.FS(content)
	}
}

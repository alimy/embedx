package embedx

import (
	"embed"
	"io/fs"
	"strings"
)

type bundleEmbedFS struct {
	*embed.FS
	root           string
	attachDir      string
	attachDirSlash string
}

// Open opens the named file for reading and returns it as an fs.File.
func (f *bundleEmbedFS) Open(name string) (fs.File, error) {
	return f.FS.Open(f.path(name))
}

// ReadDir reads and returns the entire named directory.
func (f *bundleEmbedFS) ReadDir(name string) ([]fs.DirEntry, error) {
	return f.FS.ReadDir(f.path(name))
}

// ReadFile reads and returns the content of the named file.
func (f *bundleEmbedFS) ReadFile(name string) ([]byte, error) {
	return f.FS.ReadFile(f.path(name))
}

func (f *bundleEmbedFS) path(name string) string {
	if name == f.attachDir || name == "." {
		return f.root
	}
	name = strings.TrimPrefix(name, f.attachDirSlash)
	return strings.Join([]string{f.root, name}, "/")
}

func newBundleEmbedFS(content *embed.FS, rootDir string, attachDir string) EmbedFS {
	return &bundleEmbedFS{
		FS:             content,
		root:           rootDir,
		attachDir:      attachDir,
		attachDirSlash: attachDir + "/",
	}
}

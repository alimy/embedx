## embedx
just an extension for go:embed

#### Usage
```bash
%> cd demo # change to your golang project root directory; cd <your-project-dir>
%> go get github.com/alimy/embedx
%> tree
 |- public
    |- ...
    |- index.html
    |- ...
 |- conf
    |- app.ini
    |- ...
    |- conf.go
 |- ...
 |- main.go
 |- go.mod
 |- ...
```

```go
// file: conf/assets.go

package assets

import (
	"embed"
	"net/http"
	
	"github.com/alimy/embedx"
)

func NewConfigFS() http.FileSystem {
	//go:embed app.ini
	var content embed.FS

	// attach a root to conf dir then access files in this returned FS will
	// need add  'conf' prefix. eg: access app.ini need FS.ReadFile("conf/app.ini").
	return embedx.NewFileSystem(&content, embedx.AttachRoot("conf"))
}
```
```go
// file: main.go

package assets

import (
	"embed"
	"net/http"
	
	"github.com/alimy/embedx"
)

func newPublicFS() http.FileSystem {
	//go:embed public
	var content embed.FS
	
	// change the root to public dir then access files in this returned FS will
	// not need  'public' prefix. eg: access public/index.html just need FS.ReadFile("index.html").
	return embedx.NewFileSystem(&content, embedx.ChangeRoot("public"))
}
```

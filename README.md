## embedx
just an extension for go:embed

#### Usage
```bash
%> cd demo # change to your golang project root directory; cd <your-project-dir>
%> go get github.com/alimy/embedx
%> tree
 |- public
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
	
	return embedx.NewFileSystem(&content, embedx.AttachRootOpt("conf"))
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
	
	return embedx.NewFileSystem(&content, embedx.ChangeRootOpt("public"))
}
```


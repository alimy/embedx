## embedx
[![Build Status](https://api.travis-ci.com/alimy/embedx.svg?branch=master)](https://travis-ci.com/alimy/embedx)
[![GoDoc](https://godoc.org/github.com/alimy/embedx?status.svg)](https://pkg.go.dev/github.com/alimy/embedx)
[![sourcegraph](https://img.shields.io/badge/view%20on-Sourcegraph-brightgreen.svg?logo=sourcegraph)](https://sourcegraph.com/github.com/alimy/embedx)

Just an extension for go:embed.

#### Usage
```bash
%> cd demo # change to your golang project root directory; cd <your-project-dir>
%> go get github.com/alimy/embedx
%> tree
.
 |- assets
    |- public
       |- ...
       |- index.html
       |- ...
    |- conf
       |- ...
       |- app.ini
       |- ...
    |- templates
       |- account
          |- login.tmpl
          |- ...
          |- register.tmpl
       |- ...
       |- index.tmpl
       |- ...
    |- assets.go
 |- ...
 |- main.go
 |- go.mod
 |- ...
``` 
```go
// file: assets/assets.go

package assets

import (
    "embed"
    "log"
    "path"
    "text/template"

    "github.com/alimy/embedx"
)

var (
    //go:embed conf/app.ini
    configFS embed.FS
    
    //go:embed public
    publicFS embed.FS
    
    //go:embed templates
    tmplFS   embed.FS
)

func NewConfigFS() embedx.EmbedFS {
    // attach a root to conf dir then access files in this returned FS will
    // need add  'custom' prefix. eg: access app.ini need FS.ReadFile("custom/conf/app.ini").
    return embedx.AttachRoot(configFS, "custom")
}

func NewPublicFS() embedx.EmbedFS {
    // change the root to public dir then access files in this returned FS will
    // not need  'public' prefix. eg: access public/index.html just need FS.ReadFile("index.html").
    return embedx.ChangeRoot(publicFS, "public")
}

func NewTemplate() *template.Template {
    // change root to templates directory
    embedFS := embedx.ChangeRoot(tmplFS, "templates")
	
    // register custom namer that return filepath without file name extension.
    embedx.RegisterNamer(embedx.NamerFunc(func (filepath string) string {
        ext := path.Ext(filepath)
        return filepath[:len(filepath)-len(ext)]	
    }))
	
    // parse template files from embed.FS then execute template by template path file name. 
    // eg: tmpl.ExecuteTemplate(w, "account/login", data)
    tmpl, err := embedx.ParseFS(embedFS, "*.tmpl", "account/*.tmpl")
    if err != nil {
        log.Fatal(err)
    }
    return tmpl
}
```

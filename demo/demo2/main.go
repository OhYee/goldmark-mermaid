package main

import (
	"bytes"
	"io/ioutil"
	"path"
	"runtime"

	ext "github.com/OhYee/goldmark-fenced_codeblock_extension"
	mermaid "github.com/OhYee/goldmark-mermaid"

	"github.com/yuin/goldmark"
	highlighting "github.com/yuin/goldmark-highlighting"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
)

var raw = "```go\npackage main\n\nimport ()\nfunc main(){}\n```\n\n```mermaid\npie\n    title Key elements in Product X\n    \"Calcium\" : 42.96\n    \"Potassium\" : 50.05\n    \"Magnesium\" : 10.01\n    \"Iron\" :  5\n```"

func main() {
	md := goldmark.New(
		goldmark.WithExtensions(
			extension.GFM,
			ext.NewExt(
				mermaid.RenderMap(50, "mermaid"),
				ext.RenderMap{
					Languages:      []string{"*"},
					RenderFunction: ext.GetFencedCodeBlockRendererFunc(highlighting.NewHTMLRenderer()),
				},
			),
		),
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(),
		),
		goldmark.WithRendererOptions(),
	)
	buf := bytes.Buffer{}
	if err := md.Convert([]byte(raw), &buf); err != nil {
		panic(err.Error())
	}

	_, file, _, _ := runtime.Caller(0)
	if err := ioutil.WriteFile(path.Join(path.Dir(file), "output.html"), buf.Bytes(), 0744); err != nil {
		panic(err.Error())
	}
}

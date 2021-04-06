// Package mermaid is a extension for the goldmark(http://github.com/yuin/goldmark).
//
// This extension adds svg picture output from mermaid language using
// go-mermaid(https://github.com/OhYee/go-mermaid).
package mermaid

import (
	"bytes"
	"crypto/sha1"

	mermaid "github.com/OhYee/go-mermaid"
	ext "github.com/OhYee/goldmark-fenced_codeblock_extension"
	fp "github.com/OhYee/goutils/functional"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/util"
)

// Default Mermaid extension when there is no other fencedCodeBlock goldmark render extensions
var Default = NewMermaidExtension(50, "mermaid")

// RenderMap return the goldmark-fenced_codeblock_extension.RenderMap
func RenderMap(length int, languages ...string) ext.RenderMap {
	return ext.RenderMap{
		Languages:      languages,
		RenderFunction: NewMermaid(length, languages...).Renderer,
	}
}

// NewMermaidExtension return the goldmark.Extender
func NewMermaidExtension(length int, languages ...string) goldmark.Extender {
	return ext.NewExt(RenderMap(length, languages...))
}

// Mermaid render struct
type Mermaid struct {
	Languages []string
	buf       map[string][]byte
	MaxLength int
}

// NewMermaid initial a Mermaid struct
func NewMermaid(length int, languages ...string) *Mermaid {
	return &Mermaid{Languages: languages, buf: make(map[string][]byte), MaxLength: length}
}

// Renderer render function
func (u *Mermaid) Renderer(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	n := node.(*ast.FencedCodeBlock)
	language := string(n.Language(source))

	if fp.AnyString(func(l string, idx int) bool {
		return l == language
	}, u.Languages) {
		if !entering {
			raw := u.getLines(source, node)
			h := sha1.New()
			h.Write(raw)
			hash := string(h.Sum([]byte{}))
			if result, exist := u.buf[hash]; exist {
				w.Write(result)
			} else {
				svg := []byte(mermaid.Render(string(raw)))
				if len(u.buf) >= u.MaxLength {
					u.buf = make(map[string][]byte)
				}
				u.buf[hash] = svg
				w.Write(svg)
			}
		}
	}
	return ast.WalkContinue, nil
}

func (u *Mermaid) getLines(source []byte, n ast.Node) []byte {
	buf := bytes.NewBuffer([]byte{})
	l := n.Lines().Len()
	for i := 0; i < l; i++ {
		line := n.Lines().At(i)
		buf.Write(line.Value(source))
	}
	return buf.Bytes()
}

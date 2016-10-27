package parser

import (
	"io"
	"log"

	"github.com/bowei/vu/ast"
)

/*
LineParser creates the following DOM:

	<line>
		<indent />
	    <text />
*/
type LineParser struct {
	input   io.Reader
	seenEOF bool
	buf     []byte

	line *ast.LineNode
	doc  *ast.DocumentNode
}

func NewLineParser(input io.ReadCloser) *LineParser {
	return &LineParser{
		input: input,
		doc:   &ast.DocumentNode{},
	}
}

func (p *LineParser) Parse() {
	for !p.atEOF() {
		p.line = &ast.LineNode{}

		p.parseIndent()
		p.parseText()

		p.doc.Children = append(p.doc.Children, p.line)
	}

	log.Println(p.doc)
}

func (p *LineParser) parseIndent() {
	indent := 0

	for {
		ch, ok := p.peek()
		if !ok {
			return
		}

		switch ch {
		case ' ':
			fallthrough
		case '\t':
			p.consume(1)
			indent++
		default:
			p.line.Children = append(
				p.line.Children,
				&ast.IndentNode{Indent: indent})
			return
		}
	}
}

func (p *LineParser) parseText() {
	text := ""

	for {
		ch, ok := p.peek()
		if !ok {
			return
		}

		p.consume(1)

		if p.atEOF() || ch == '\n' {
			p.line.Children = append(
				p.line.Children,
				&ast.TextNode{Text: text})
			return
		}

		text += string(ch)
	}
}

func (p *LineParser) peek() (byte, bool) {
	if len(p.buf) == 0 {
		p.fill()
	}

	if len(p.buf) == 0 {
		return 0, false
	}

	return p.buf[0], true
}

func (p *LineParser) consume(n int) {
	p.buf = p.buf[n:]
}

func (p *LineParser) atEOF() bool {
	return p.seenEOF && len(p.buf) == 0
}

func (p *LineParser) fill() {
	buf := make([]byte, 4096)
	len, err := p.input.Read(buf)

	p.buf = append(p.buf, buf[:len]...)
	if err == io.EOF {
		p.seenEOF = true
		return
	} else if err != nil {
		panic(err)
	}
}

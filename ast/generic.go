/*
Copyright 2016 Bowei Du

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package ast

import "fmt"

type DocumentNode struct {
	location Location
	Children []Node
}

type LineNode struct {
	location Location
	Children []Node
}

// NewLine ast node.
func NewLine(location Location) Node {
	return &LineNode{
		location: location,
	}
}

func (n *LineNode) Location() Location {
	return n.location
}

func (n *LineNode) Attributes() map[string]string {
	return make(map[string]string)
}

func (n *LineNode) String() string {
	ret := "<line>\n"
	for i := range n.Children {
		ret += "  " + n.Children[i].String()
	}
	ret += "</line>\n"

	return ret
}

type IndentNode struct {
	location Location
	Indent   int
}

// NewIndent ast node.
func NewIndent(location Location, indent int) Node {
	return &IndentNode{
		location: location,
		Indent:   indent,
	}
}

func (n *IndentNode) Location() Location {
	return n.location
}

func (n *IndentNode) Attributes() map[string]string {
	return make(map[string]string)
}

func (n *IndentNode) String() string {
	return fmt.Sprintf("<indent %d />", n.Indent)
}

type TextNode struct {
	location Location
	Text     string
}

// NewText ast node.
func NewText(location Location, text string) Node {
	return &TextNode{
		location: location,
		Text:     text,
	}
}

func (n *TextNode) Location() Location {
	return n.location
}

func (n *TextNode) Attributes() map[string]string {
	return make(map[string]string)
}

func (n *TextNode) String() string {
	return "<text>" + n.Text + "</text>"
}

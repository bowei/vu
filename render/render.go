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

package render

type Render interface {
	Render(data.
	AstNode tree)
}

func (l *Line) ToStr() string {
	return l.Content
}

type Render struct {
	Lines []Line
}

func (render *Render) ToStr() string {
	str := ""

	for i := range render.Lines {
		str += render.Lines[i].ToStr()
		str += "\n"
	}
	return str
}

// NewFromStr creates a new Render from string
func NewFromStr(contents string) *Render {
	return &Render{
		Lines: []Line{
			Line{
				Content: "one",
			},
			Line{
				Content: "two",
			},
		},
	}
}

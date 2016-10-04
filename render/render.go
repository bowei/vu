package render

type Line struct {
	Content string
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

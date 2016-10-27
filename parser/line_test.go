package parser

import (
	"os"
	"testing"
)

func TestLineParser(t *testing.T) {
	file, err := os.Open("test.txt")
	if err != nil {
		t.Fatal(err)
	}

	parser := NewLineParser(file)
	parser.Parse()
}

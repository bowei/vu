package parser

import (
	"bufio"
	"log"
	"os"
)

type fileParser struct {
	output chan string
}

func (fp *fileParser) Output() <-chan string {
	return fp.output
}

func parseFile(filename string) (Parser, error) {
	parser := &fileParser{
		output: make(chan string),
	}

	fh, err := os.Open(filename)
	if err != nil {
		log.Printf("Error opening %v: %v", filename, err)
		return nil, err
	}

	log.Printf("Opened %v", filename)

	go func() {
		log.Printf("Scanning file")

		scanner := bufio.NewScanner(fh)
		for scanner.Scan() {
			line := scanner.Text()
			log.Printf("Parsing line %v", line)
			parser.output <- line
		}
		close(parser.output)

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}()

	return parser, nil
}

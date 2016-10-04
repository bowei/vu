package parser

import (
	"bufio"
	"log"
	"net/http"
)

type httpParser struct {
	output chan string
}

func (fp *httpParser) Output() <-chan string {
	return fp.output
}

func parseHTTP(src string) (Parser, error) {
	parser := &httpParser{
		output: make(chan string),
	}

	resp, err := http.Get(src)
	log.Printf("URI %v: %v", src, err)

	if err != nil {
		return nil, err
	}

	go func() {
		scanner := bufio.NewScanner(resp.Body)
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

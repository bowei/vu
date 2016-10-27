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

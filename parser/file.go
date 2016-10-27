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

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

/*
VU executable.
*/
package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/bowei/vu/ui"
)

// Params from the command line.
type Params struct {
	// HTTP listen address.
	HTTP string
}

func main() {
	params := parseArgs()

	WebUI := ui.WebUI{}
	WebUI.Setup()

	log.Printf("params %+v\n", params)

	server := &http.Server{
		Addr: params.HTTP,
	}
	log.Fatal(server.ListenAndServe())
}

func parseArgs() *Params {
	params := &Params{}
	flag.StringVar(&params.Http, "http", ":9800", "http server")

	return params
}

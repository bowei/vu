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
Package ui is the web UI.
*/
package ui

import (
	"fmt"
	"html"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/bowei/vu/parser"
)

const (
	mainHead = `<!DOCTYPE html>
<head>
	<title>vu</title>
	<script src="https://ajax.googleapis.com/ajax/libs/jquery/3.1.1/jquery.min.js"></script>
</head>

<body>
	<div id="pasteArea" style="background-color: red">#</div>
	<form method="get" action="/">
		<label>source</label>
		<input name="src"/>
	</form>
	<script>
		$('#pasteArea').on('paste', function (event) {
				var element = this;
				console.log(event.originalEvent.clipboardData.getData('text'))
				return false
		});
	</script>
<hr/>

<pre>`
	mainTail = `</pre>`
)

// WebUI is the object that encapsulates the web server.
type WebUI struct {
}

// Setup the web UI.
func (w *WebUI) Setup() {
	http.HandleFunc("/", w.mainPage)
	http.HandleFunc("/paste", w.pasteContents)
	http.HandleFunc("/saved/", w.savedContents)
}

func (*WebUI) mainPage(w http.ResponseWriter, r *http.Request) {
	log.Printf("mainPage")

	src := r.URL.Query().Get("src")
	log.Printf("src = %q", src)

	w.Header().Set("Content-Type", "text/html")
	head := mainHead

	if _, err := w.Write([]byte(head)); err != nil {
		log.Printf("Error writing page: %v", err)
		return
	}

	parser, err := parser.Parse(src)
	if err == nil {
		for line := range parser.Output() {
			fmt.Fprintf(w, "%s\n", html.EscapeString(line))
		}
	} else {
		fmt.Fprintf(w, "error loading from %v: %v", src, err)
	}

	tail := mainTail
	if _, err := w.Write([]byte(tail)); err != nil {
		log.Printf("Error writing page: %v", err)
		return
	}
}

func (*WebUI) pasteContents(w http.ResponseWriter, r *http.Request) {
	log.Printf("pasteContents")

	if r.Method != "POST" {
		http.Error(w, "invalid method", 400)
		return
	}

	bytes, err := ioutil.ReadAll(r.Body)
	type jsonData struct {
	}

	fmt.Println(bytes, err) // XXX
}

func (*WebUI) savedContents(w http.ResponseWriter, r *http.Request) {
}

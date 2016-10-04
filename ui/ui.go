package ui

import (
	"fmt"
	"html"
	"log"
	"net/http"

	"github.com/bowei/vu/parser"
)

func SetupServer() {
	http.HandleFunc("/", mainPage)
}

func mainPage(w http.ResponseWriter, r *http.Request) {
	src := r.URL.Query().Get("src")
	log.Printf("src = %q", src)

	w.Header().Set("Content-Type", "text/html")
	head := `<!DOCTYPE html>
		<html>
		<head><title>vu</title></head>
		<body>` + html.EscapeString(src) +
		`<form method="get" action="/">
		<label>source</label>
		<input name="src"/>
		</form>
		<hr/>
		<pre>`

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

	tail := `</pre></html>`
	if _, err := w.Write([]byte(tail)); err != nil {
		log.Printf("Error writing page: %v", err)
		return
	}
}

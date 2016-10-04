package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/bowei/vu/ui"
)

type Params struct {
	Addr string
}

func main() {
	params := parseArgs()
	ui.SetupServer()

	log.Printf("params %+v\n", params)

	server := &http.Server{
		Addr: params.Addr,
	}
	log.Fatal(server.ListenAndServe())
}

func parseArgs() *Params {
	params := &Params{}
	flag.StringVar(
		&params.Addr, "addr",
		"localhost:9800", "http server addr")

	return params
}

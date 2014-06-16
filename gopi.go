/*
Package gopi provides an out-of-box web server with reasonable defaults.

Example:
	package main

	import (
		"fmt"
		"net/http"

		"github.com/ludmiloff/gopi"
		"github.com/ludmiloff/gopi/web"
	)

	func hello(c web.C, w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %s!", c.URLParams["name"])
	}

	func main() {
		gopi.UseDefaults()
		gopi.Get("/hello/:name", hello)
		gopi.Serve()
	}

This package exists purely as a convenience to programmers who want to get
started as quickly as possible. It draws almost all of its code from gopi's
subpackages, the most interesting of which is gopi/web, and where most of the
documentation for the web framework lives.

A side effect of this package's ease-of-use is the fact that it is opinionated.
If you don't like (or have outgrown) its opinions, it should be straightforward
to use the APIs of gopi's subpackages to reimplement things to your liking. Both
methods of using this library are equally well supported.

GoPI requires Go 1.2 or newer.
*/
package gopi

import (
	"flag"
	"log"
	"net/http"

	"github.com/ludmiloff/gopi/bind"
	"github.com/ludmiloff/gopi/graceful"
)

// Serve starts GoPI
func Serve() {
	if !flag.Parsed() {
		flag.Parse()
	}

	log.SetFlags(log.Flags() | log.Lmicroseconds)

	// Install our handler at the root of the standard net/http default mux.
	// This allows packages like expvar to continue working as expected.
	http.Handle("/", DefaultMux)

	listener := bind.Default()
	log.Println("Starting GoPI on", listener.Addr())

	bind.Ready()

	err := graceful.Serve(listener, http.DefaultServeMux)

	if err != nil {
		log.Fatal(err)
	}

	graceful.Wait()
}

package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/FATHOM5/rest"
	"github.com/FATHOM5/ryer-server/shared"
)

// Routes is a method that creates handlers for the type ExampleServer
func (e *ExampleServer) Routes() {
	e.Router.Handle("/greeting", rest.M.StrictPOST(e.greetingHandler()))
	e.Router.Handle("/version", e.versionHandler())
}

// versionHandler is a handler function for the ExampleServer
func (ExampleServer) versionHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_ = rest.RespondOK(w, rest.ServiceVersion{GitTag: "v1.0.0"})
	})
}

func (ExampleServer) greetingHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Client request parsing
		// AddParams is the expected request. If the request params are unknown,
		// we send back a bad request code and error.
		var name shared.Params
		err := rest.ParseClientRequest(r, &name)
		if err != nil {
			_ = rest.RespondUnknownClientParams(w, rest.ErrorFields{
				"reason": err,
			})
			return
		}

		// This is if the params unmarshaled correctly, but are not valid
		// prams due to application logic (I expect more than 1 number to add up).
		// We can pass some additional info to the invalid params for human debugging.
		if err := name.Valid(); err != nil {
			_ = rest.RespondInvalidParams(w, rest.ErrorFields{
				"reason": err,
				"type":   "GreetingParams",
			})
			return
		}

		// Call the implementation
		gr, err := Greeting(name)
		if err != nil {
			_ = rest.RespondBadRequest(w, fmt.Errorf("failed to create greeting"), rest.ErrorFields{
				"reason": err,
			})
			return
		}

		// Log for debugging/logging
		log.Printf("greeting = %s", gr)

		// Response
		_ = rest.RespondOK(w, gr)
	})
}

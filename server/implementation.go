package main

import (
	"fmt"

	"github.com/FATHOM5/ryer-server/shared"
)

// Greeting uses the name parameter to return a greeting string
func Greeting(params shared.Params) (shared.Str, error) {
	if len(params) == 0 {
		return "", fmt.Errorf("greeting expects at least 1 argument")
	}
	var str shared.Str = "My name is"
	for _, s := range params {
		str += " " + shared.Str(s)
	}
	return str, nil
}

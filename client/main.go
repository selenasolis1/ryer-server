package main

import (
	"errors"
	"flag"
	"fmt"
	"os"

	"github.com/FATHOM5/rest/httptypes"
	"github.com/FATHOM5/ryer-server/shared"
)

var (
	ServerAddress string
	DebugClient   bool
)

func main() {
	flag.StringVar(&ServerAddress, "s", "localhost:6500", "The url of the server")
	flag.BoolVar(&DebugClient, "d", false, "Turn on client request logging")
	flag.Parse()

	if len(flag.Args()) == 0 {
		usage()
		os.Exit(1)
	}

	// Optionally pass a client with the proper timeout settings
	cli := NewClient(nil, ServerAddress)
	if DebugClient {
		cli.Debug()
	}

	if flag.Args()[0] == "version" {
		ver, err := cli.Version()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("GreetingServer version is: %s\n", ver.GitTag)
		return
	}

	var params shared.Params
	// Gather params
	for _, arg := range flag.Args() {
		params = append(params, arg)
	}

	reply, err := cli.Greeting(params)
	if err != nil {
		fmt.Println("Request failed")
		if errors.Is(err, httptypes.ErrorInvalidParams) {
			fmt.Println("Invalid parameters")
		}
		fmt.Println(err)
		return
	}

	fmt.Println("greeting: ", reply)
}

func usage() {
	fmt.Println("Client will communicate with the server.")
	fmt.Println("It will send your name to the server to create the greeting")
	fmt.Println("Usage: client <name>")
}

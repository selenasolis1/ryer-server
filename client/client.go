package main

import (
	"net/http"

	"github.com/FATHOM5/rest"
	"github.com/FATHOM5/ryer-server/shared"
)

type Client struct {
	client        *rest.RestfulClient
	serverAddress string
}

func (cli *Client) Debug() {
	cli.client.Debug()
}

func NewClient(cli *http.Client, addr string) *Client {
	a := new(Client)
	if cli == nil {
		cli = http.DefaultClient
	}
	a.client = rest.NewRestClient(cli)
	a.serverAddress = addr

	return a
}

func (cli Client) Greeting(params shared.Params) (reply shared.Str, err error) {
	rc := cli.client
	_, err = rc.Post(rc.Join(cli.serverAddress, "greeting"), params, &reply)
	return
}

func (cli Client) Version() (reply rest.ServiceVersion, err error) {
	rc := cli.client
	_, err = rc.Post(rc.Join(cli.serverAddress, "version"), nil, &reply)
	return
}

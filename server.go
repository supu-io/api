package main

import (
	"github.com/go-martini/martini"
	"github.com/nats-io/nats"
)

var c *nats.EncodedConn
var nc *nats.Conn
var env string
var mockedResponse []byte

func setupRouter() *martini.ClassicMartini {
	m := martini.Classic()

	m.Get("/", GetHome)
	m.Get("/issues", GetIssues)
	m.Get("/issues/:issue", GetIssue)
	m.Put("/issues/:issue", UpdateIssue)

	return m
}

func main() {
	nc, _ = nats.Connect(nats.DefaultURL)
	c, _ = nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	defer c.Close()

	m := setupRouter()
	m.Run()
}

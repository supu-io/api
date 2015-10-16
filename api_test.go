package main

import (
	"net/http"

	"github.com/nats-io/nats"
)

var w http.ResponseWriter
var r http.Request

func setup() {
	nc, _ = nats.Connect(nats.DefaultURL)
	c, _ = nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	// defer c.Close()
}

func subscribe(subject string, respond string) {
	sub, _ := nc.Subscribe(subject, func(m *nats.Msg) {
		if respond == "" {
			nc.Publish(m.Reply, m.Data)
		} else {
			nc.Publish(m.Reply, []byte(respond))
		}
	})
	sub.AutoUnsubscribe(1)
}

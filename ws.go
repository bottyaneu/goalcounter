package main

import (
	"context"

	"github.com/bndrmrtn/go-gale"
)

func NewWSServer(app *gale.Gale) gale.WSServer {
	server := gale.NewWSServer(context.Background())

	app.WS("/ws", func(conn gale.WSConn) {
		conn.Ctx().Set("ip_addr", conn.Ctx().IP())
		server.AddConn(conn)
	}).Name("ws.connect")

	server.OnMessage(func(s gale.WSServer, conn gale.WSConn, msg []byte) error {
		// Keep connection alive by responding to ping messages
		if string(msg) == "ping" {
			return conn.Send([]byte("pong"))
		}

		// Otherwise skip message
		return nil
	})

	return server
}

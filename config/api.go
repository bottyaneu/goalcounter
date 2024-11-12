package config

import (
	"net/http"

	"github.com/bndrmrtn/go-gale"
	"github.com/coder/websocket"
)

func Api() gale.Config {
	return gale.Config{
		Mode: Mode(),
		NotFoundHandler: func(c gale.Ctx) error {
			return c.Status(http.StatusNotFound).JSON(gale.Map{
				"error": "Not Found",
			})
		},
		Session: &gale.SessionConfig{
			Enabled: false,
		},
		Websocket: &gale.WSConfig{
			AcceptOptions: &websocket.AcceptOptions{
				InsecureSkipVerify: true,
			},
		},
	}
}

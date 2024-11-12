package middlewares

import (
	"github.com/bndrmrtn/go-gale"
)

func CORSMiddleware(c gale.Ctx) error {
	c.Header().Add("Access-Control-Allow-Origin", c.Request().Host)
	c.Header().Add("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	c.Header().Add("Access-Control-Allow-Headers", "Content-Type, Authorization")
	c.Header().Add("Access-Control-Allow-Credentials", "true")

	if c.Method() == "OPTIONS" {
		c.Status(204).Break() // No content and stop the middleware chain
		return nil            // No content
	}

	return nil
}

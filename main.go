package main

import (
	"flag"

	"github.com/sirupsen/logrus"
)

var listenAddr = flag.String("listenAddr", ":3000", "The address to listen on for HTTP requests.")

// Simple API server for goal counting
// With websocket support to keep track of goals in real-time
// Made with ❤️ by Martin Binder
// https://mrtn.vip
func main() {
	// Parse flags
	flag.Parse()

	// Create new API server
	api := NewApiServer()

	// Start API server
	logrus.Fatal(api.Serve(*listenAddr))
}

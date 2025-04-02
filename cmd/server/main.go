package main

import (
	"github.com/eliasmeireles/go-pdf-generator/pkg/handler"
	"github.com/softwareplace/goserve/logger"
	"github.com/softwareplace/goserve/server"
)

func init() {
	logger.LogSetup()
}

func main() {
	server.Default().
		ContextPath("/api/pdf-generator/v1/").
		Get(handler.PDFRequestHandler, "/pdf").
		StartServer()
}

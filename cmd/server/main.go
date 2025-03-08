package main

import (
	"github.com/softwareplace/http-utils/server"
	"go-pdf-generator/pkg/handler"
)

func main() {
	server.Default().
		WithContextPath("/api/pdf-generator/v1/").
		Get(handler.PDFRequestHandler, "/pdf").
		StartServer()
}

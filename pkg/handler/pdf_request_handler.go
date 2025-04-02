package handler

import (
	"encoding/base64"
	"github.com/eliasmeireles/go-pdf-generator/pkg/pdf"
	"github.com/google/uuid"
	goservectx "github.com/softwareplace/goserve/context"
)

func PDFRequestHandler(ctx *goservectx.Request[*goservectx.DefaultContext]) {
	url := ctx.QueryOf("url")
	isBase64 := ctx.QueryOf("base64")
	expectedIds := ctx.QueriesOf("id")
	expectedClasses := ctx.QueriesOf("class")
	appendText := ctx.QueryOf("appendText")

	if isBase64 == "true" {
		decodedUrl, err := base64.StdEncoding.DecodeString(url)
		if err != nil {
			ctx.InternalServerError("Failed to decode base64: " + err.Error())
			return
		}
		url = string(decodedUrl)
	}

	if url != "" {
		generatedPDF, err := pdf.GetPDFFromURL(url, expectedIds, expectedClasses, appendText)
		if err != nil {
			ctx.InternalServerError("Failed to generate the pdf: " + err.Error())
		}

		fileName := ctx.QueryOfOrElse("fileName", uuid.NewString())

		err = ctx.WriteFile(generatedPDF, fileName)

		if err != nil {
			ctx.InternalServerError("Failed to stream file: " + err.Error())
			return
		}
	}

	ctx.BadRequest("Requires url query parameter")
}

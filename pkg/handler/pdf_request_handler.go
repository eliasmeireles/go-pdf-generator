package handler

import (
	"github.com/eliasmeireles/go-pdf-generator/pkg/pdf"
	"github.com/google/uuid"
	goservectx "github.com/softwareplace/goserve/context"
)

func PDFRequestHandler(ctx *goservectx.Request[*goservectx.DefaultContext]) {
	url := ctx.QueryOf("url")
	expectedIds := ctx.QueriesOf("id")
	expectedClasses := ctx.QueriesOf("class")
	appendText := ctx.QueryOf("appendText")

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

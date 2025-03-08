package handler

import (
	"github.com/google/uuid"
	"github.com/softwareplace/http-utils/api_context"
	"go-pdf-generator/pkg/pdf"
)

func PDFRequestHandler(ctx *api_context.ApiRequestContext[*api_context.DefaultContext]) {
	url := ctx.QueryOf("url")
	expectedIds := ctx.QueriesOf("id")
	expectedClasses := ctx.QueriesOf("class")

	if url != "" {
		generatedPDF, err := pdf.GetPDFFromURL(url, expectedIds, expectedClasses)
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

package handler

import (
	"bytes"
	"fmt"
	"github.com/google/uuid"
	"github.com/softwareplace/http-utils/api_context"
	"go-pdf-generator/pkg/pdf"
	"io"
)

func PDFRequestHandler(ctx *api_context.ApiRequestContext[*api_context.DefaultContext]) {
	url := ctx.QueryValues["url"][0]
	expectedIds := ctx.QueryValues["id"]

	if url != "" {
		generatedPDF, err := pdf.GetPDFFromURL(url, expectedIds)
		if err != nil {
			ctx.InternalServerError("Failed to generate the pdf: " + err.Error())
		}
		writer := *ctx.Writer

		writer.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s.pdf", uuid.NewString()))
		writer.Header().Set("Content-Type", "application/octet-stream")

		_, err = io.Copy(writer, bytes.NewReader(generatedPDF))
		if err != nil {
			ctx.InternalServerError("Failed to stream file: " + err.Error())
		}
		return
	}

	ctx.BadRequest("Requires url query parameter")

}

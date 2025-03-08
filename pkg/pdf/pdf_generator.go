package pdf

import (
	"context"
	"fmt"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
	"time"
)

func GetPDFFromURL(templateUrl string, ids []string) ([]byte, error) {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	var pdfBuffer []byte
	err := chromedp.Run(ctx, printPdf(&pdfBuffer, templateUrl, ids))
	if err != nil {
		return nil, err
	}

	return pdfBuffer, nil
}
func printPdf(pdfBuffer *[]byte, url string, ids []string) chromedp.Tasks {
	tasks := chromedp.Tasks{
		// Use the data: URL scheme to load the HTML content directly
		chromedp.Navigate(url),
	}

	// Add visibility checks for each ID in the ids slice
	for _, id := range ids {
		tasks = append(tasks, chromedp.Tasks{
			chromedp.Evaluate(fmt.Sprintf(`console.log("Checking visibility of #%s...");`, id), nil),
			chromedp.Evaluate(fmt.Sprintf(`console.log(document.getElementById("%s") ? "Element exists" : "Element does not exist");`, id), nil),
			chromedp.Evaluate(fmt.Sprintf(`console.log(document.getElementById("%s").offsetParent ? "Element is visible" : "Element is hidden");`, id), nil),
			chromedp.WaitVisible("#"+id, chromedp.ByQuery),
		}...)
	}

	// Add the PDF generation task
	tasks = append(tasks, chromedp.Tasks{
		// Optional: Add a delay to ensure everything is fully rendered
		chromedp.Sleep(2 * time.Second), // Adjust the duration as needed
		chromedp.ActionFunc(func(ctx context.Context) error {
			var err error
			// Configure PDF options
			*pdfBuffer, _, err = page.PrintToPDF().
				WithPrintBackground(true). // Include background colors/images
				WithPaperWidth(8.5).
				WithPaperHeight(11).
				WithMarginTop(0.5).
				WithMarginBottom(0.5).
				WithMarginLeft(0.5).
				WithMarginRight(0.5).
				WithScale(1.0).
				Do(ctx)
			return err
		}),
	}...)

	return tasks
}

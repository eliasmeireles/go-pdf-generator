package pdf

import (
	"context"
	"fmt"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
	"time"
)

func GetPDFFromURL(templateUrl string, ids []string, classes []string, additionalText string) ([]byte, error) {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	var pdfBuffer []byte
	err := chromedp.Run(ctx, printPdf(&pdfBuffer, templateUrl, ids, classes, additionalText))
	if err != nil {
		return nil, err
	}

	return pdfBuffer, nil
}

func printPdf(pdfBuffer *[]byte, url string, ids []string, classes []string, additionalText string) chromedp.Tasks {
	tasks := chromedp.Tasks{
		chromedp.Navigate(url),
	}

	// Inject additional text into the DOM
	if additionalText != "" {
		tasks = append(tasks, chromedp.Tasks{
			chromedp.Evaluate(fmt.Sprintf(`
			var div = document.createElement('div');
			div.style.position = 'fixed';
			div.style.bottom = '10px';
			div.style.right = '10px';
			div.style.backgroundColor = 'white';
			div.style.border = '1px solid black';
			div.style.padding = '5px';
			div.style.zIndex = '1000';
			div.textContent = '%s';
			document.body.appendChild(div);
		`, additionalText), nil),
		}...)
	}

	// Add visibility checks for each class in the classes slice
	for _, class := range classes {
		tasks = append(tasks, chromedp.Tasks{
			chromedp.Evaluate(fmt.Sprintf(`console.log("Checking visibility of .%s...");`, class), nil),
			chromedp.Evaluate(fmt.Sprintf(`console.log(document.querySelector(".%s") ? "Element exists" : "Element does not exist");`, class), nil),
			chromedp.Evaluate(fmt.Sprintf(`console.log(document.querySelector(".%s") && document.querySelector(".%s").offsetParent ? "Element is visible" : "Element is hidden");`, class, class), nil),
			chromedp.WaitVisible("."+class, chromedp.ByQuery),
		}...)
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
		chromedp.Sleep(2 * time.Second), // Adjust the duration as needed
		chromedp.ActionFunc(func(ctx context.Context) error {
			var err error
			*pdfBuffer, _, err = page.PrintToPDF().
				WithPrintBackground(true).
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

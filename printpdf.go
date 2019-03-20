package printpdf

import (
	"bytes"
	"context"
	"io"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
	"github.com/chromedp/chromedp/runner"
)

func NewReader(url string) (io.Reader, error) {
	ctx, cancel := context.WithCancel(context.Background())

	c, err := chromedp.New(ctx, chromedp.WithRunnerOptions(
		runner.Flag("headless", true),
		runner.Flag("disable-gpu", true),
	))
	if err != nil {
		return nil, err
	}

	cancel()

	var pdfReader io.Reader
	err = c.Run(ctx, chromedp.Tasks{
		chromedp.Navigate(url),
		chromedp.WaitReady("body"),
		chromedp.ActionFunc(func(ctx context.Context, h cdp.Executor) error {
			buf, err := page.PrintToPDF().Do(ctx, h)
			if err != nil {
				return err
			}
			pdfReader = bytes.NewBuffer(buf)
			return nil
		}),
	})
	if err != nil {
		return nil, err
	}
	return pdfReader, nil
}

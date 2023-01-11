package browser

import (
	"context"
	"log"
	"os"

	"github.com/chromedp/chromedp"
	"github.com/chromedp/chromedp/device"
)

type HeadlessBrowser struct {
	ctx    context.Context
	cancel context.CancelFunc
}

func NewHeadlessBrowser() *HeadlessBrowser {
	ctx, cancel := chromedp.NewContext(context.Background())
	return &HeadlessBrowser{ctx, cancel}
}

func (b *HeadlessBrowser) DontForgetToCleanUp() {
	b.cancel()
}

func (b *HeadlessBrowser) TakeFullScreenshot(url string, pathToSave string) {

	var buf []byte
	quality := 90

	if err := chromedp.Run(b.ctx, fullScreenshot(url, quality, &buf)); err != nil {
		log.Fatal(err)
	}
	saveInFile(&buf, pathToSave)
}

func (b *HeadlessBrowser) TakeElementScreenshot(url string, cssSelector string, pathToSave string) {
	var buf []byte

	if err := chromedp.Run(b.ctx, elementScreenshot(url, cssSelector, &buf)); err != nil {
		log.Fatal(err)
	}
	saveInFile(&buf, pathToSave)
}

func saveInFile(res *[]byte, pathToSave string) {
	if err := os.WriteFile(pathToSave, *res, 0o644); err != nil {
		log.Fatal(err)
	}
}

func elementScreenshot(url string, selector string, res *[]byte) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(url),
		chromedp.Screenshot(selector, res, chromedp.NodeVisible),
	}
}

func fullScreenshot(url string, quality int, res *[]byte) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(url),
		chromedp.FullScreenshot(res, quality),
		chromedp.Emulate(device.Reset),
	}
}

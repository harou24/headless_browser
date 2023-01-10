package tests

import (
	"testing"

	"github.com/harou24/go_boilerplate/browser"
)

const TEST_URL = `https://harou24.github.io/`

func TestTakingFullScreenshot(t *testing.T) {
	headlessBrowser := browser.NewHeadlessBrowser()
	defer headlessBrowser.DontForgetToCleanUp()
	headlessBrowser.TakeFullScreenshot(TEST_URL, `./test_full_screenshot.png`)
}

func TestTakingElementScreenshot(t *testing.T) {
	headlessBrowser := browser.NewHeadlessBrowser()
	defer headlessBrowser.DontForgetToCleanUp()
	headlessBrowser.TakeElementScreenshot(TEST_URL, `div.portfolio-btn`, `./test_element_screenshot.png`)
}

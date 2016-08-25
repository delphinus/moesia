package browser

import (
	"fmt"

	"github.com/sclevine/agouti"
)

const topURL = "https://as.its-kenpo.or.jp/service_category/index"
const pageWidth = 1280
const pageHeight = 1024

// Browser has WebDriver in property
type Browser struct {
	driver *agouti.WebDriver
	page   *agouti.Page
}

// New returns a new instance
func New() (self *Browser, err error) {
	self = &Browser{driver: agouti.PhantomJS()}
	if err = self.driver.Start(); err != nil {
		err = fmt.Errorf("Failed to start driver: %v", err)
		return
	}
	err = self.setPage()
	return
}

func (b *Browser) setPage() (err error) {
	capabilities := agouti.NewCapabilities()
	capabilities.Browser("safari")
	capabilities.Platform("MAC")
	if b.page, err = b.driver.NewPage(agouti.Desired(capabilities)); err != nil {
		err = fmt.Errorf("Failed to setPage(): %v", err)
		return
	}
	b.page.Size(pageWidth, pageHeight)
	return
}

// Start will start scraping
func (b *Browser) Start() (err error) {
	if err = b.page.Navigate(topURL); err != nil {
		err = fmt.Errorf("Failed to open topURL (%s): %v", topURL, err)
		return
	}
	if err = b.page.Screenshot("/tmp/test_ss.jpg"); err != nil {
		err = fmt.Errorf("Failed to save SS: %v", err)
		return
	}
	return
}

// End will finish the driver
func (b *Browser) End() (err error) {
	err = b.driver.Stop()
	return
}

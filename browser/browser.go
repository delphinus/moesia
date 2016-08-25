package browser

import (
	"fmt"

	"github.com/sclevine/agouti"
)

var topURL = "https://as.its-kenpo.or.jp/service_category/index"

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
	if self.page, err = self.driver.NewPage(agouti.Browser("phantomjs")); err != nil {
		err = fmt.Errorf("Failed to open page: %v", err)
		return
	}
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

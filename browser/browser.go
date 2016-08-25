package browser

import (
	"fmt"
	"io/ioutil"

	"github.com/sclevine/agouti"
)

const topURL = "https://as.its-kenpo.or.jp/service_category/index"
const firstLinkText = "直営・通年・夏季保養施設(空き照会)"
const pageWidth = 1280
const pageHeight = 1024

var hotelName = []string{"トスラブ箱根ビオーレ", "トスラブ箱根和奏林"}

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
	if err = b.page.FindByLink(firstLinkText).Click(); err != nil {
		err = fmt.Errorf("Failed to click '%s': %v", firstLinkText, err)
		return
	}
	for _, hotel := range hotelName {
		if err = b.page.FindByLink(hotel).Click(); err != nil {
			err = fmt.Errorf("Failed to open hotel '%s': %v", hotel, err)
			return
		}
		selections := b.page.AllByXPath(`//select[@id="apply_join_time"]/option`)
		var i = 0
		for {
			selection := selections.At(i)
			if selection == nil {
				if i == 0 {
					err = fmt.Errorf("Failed to find option in hotel '%s'", hotel)
					return
				}
				break
			}
			fmt.Println(selection.String())
			i++
		}
		b.page.Back()
	}
	return
}

// Screenshot will take screenshot
func (b *Browser) Screenshot() (filename string, err error) {
	file, err := ioutil.TempFile("", "moesia_")
	if err != nil {
		err = fmt.Errorf("ioutil.Tempfile failed: %v", err)
		return
	}
	filename = file.Name()
	if err = b.page.Screenshot(filename); err != nil {
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

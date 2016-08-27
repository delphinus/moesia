package browser

import (
	"fmt"
	"io/ioutil"

	"github.com/sclevine/agouti"
	"github.com/sclevine/agouti/api"
)

const topURL = "https://as.its-kenpo.or.jp/service_category/index"
const firstLinkText = "直営・通年・夏季保養施設(空き照会)"
const pageWidth = 1280
const pageHeight = 1024

var hotels = []string{"トスラブ箱根ビオーレ", "トスラブ箱根和奏林"}

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
	for _, hotel := range hotels {
		if err = b.page.FindByLink(hotel).Click(); err != nil {
			err = fmt.Errorf("Failed to open hotel '%s': %v", hotel, err)
			return
		}
		monthLinks := b.page.AllByXPath(fmt.Sprintf("//a[contains(., '%s')]", hotel))
		var monthLinkTexts []string
		if monthLinkTexts, err = b.getTexts(monthLinks); err != nil {
			return
		}
		for _, monthLinkText := range monthLinkTexts {
			if err = b.page.FindByLink(monthLinkText).Click(); err != nil {
				err = fmt.Errorf("Failed to click '%s' for hotel '%s': %v", monthLinkText, hotel, err)
				return
			}
			options := b.page.AllByXPath(`//select[@id="apply_join_time"]/option`)
			var optionTexts []string
			if optionTexts, err = b.getTexts(options); err != nil {
				return
			}
			for _, optionText := range optionTexts {
				fmt.Printf("%s: %s\n", hotel, optionText)
			}
			b.page.Back()
		}
		b.page.Back()
	}
	return
}

func (b *Browser) getTexts(multiSelection *agouti.MultiSelection) (texts []string, err error) {
	var elements []*api.Element
	elements, err = multiSelection.Elements()
	if err != nil {
		err = fmt.Errorf("Failed to get elements: %v", err)
		return
	}
	for _, element := range elements {
		var text string
		if text, err = element.GetText(); err != nil {
			err = fmt.Errorf("Failed to get text for element: %v", err)
			return
		}
		if len(text) > 0 {
			texts = append(texts, text)
		}
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

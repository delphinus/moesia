package browser

import (
	"fmt"
	"io/ioutil"

	"github.com/delphinus35/moesia/util"
	"github.com/delphinus35/moesia/vacancy"
	"github.com/sclevine/agouti"
	"github.com/sclevine/agouti/api"
)

var isTest = false

const topURL = "https://as.its-kenpo.or.jp/service_category/index"
const firstLinkText = "直営・通年・夏季保養施設(空き照会)"
const pageWidth = 1280
const pageHeight = 1024

var hotels = []string{"トスラブ箱根ビオーレ", "トスラブ箱根和奏林"}

// Browser has WebDriver in property
type Browser struct {
	driver *agouti.WebDriver
	page   *Page
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
	var page *agouti.Page
	if page, err = b.driver.NewPage(agouti.Desired(capabilities)); err != nil {
		err = fmt.Errorf("Failed to setPage(): %v", err)
		return
	}
	b.page = &Page{page}
	b.page.Size(pageWidth, pageHeight)
	return
}

// Process will do scraping
func (b *Browser) Process() (vacancies vacancy.Vacancies, err error) {
	if err = b.page._Navigate(topURL); err != nil {
		err = fmt.Errorf("Failed to open topURL (%s): %v", topURL, err)
		return
	}
	if err = b.page._FindByLink(firstLinkText).Click(); err != nil {
		err = fmt.Errorf("Failed to click '%s': %v", firstLinkText, err)
		return
	}
	for _, hotel := range hotels {
		if err = b.page._FindByLink(hotel).Click(); err != nil {
			err = fmt.Errorf("Failed to open hotel '%s': %v", hotel, err)
			return
		}
		monthLinks := b.page._AllByXPath(fmt.Sprintf("//a[contains(., '%s')]", hotel))
		var monthLinkTexts []string
		if monthLinkTexts, err = b.getTexts(monthLinks); err != nil {
			return
		}
		hotelVacancy := vacancy.Vacancy{Hotel: hotel}
		for _, monthLinkText := range monthLinkTexts {
			if err = b.page._FindByLink(monthLinkText).Click(); err != nil {
				err = fmt.Errorf("Failed to click '%s' for hotel '%s': %v", monthLinkText, hotel, err)
				return
			}
			options := b.page._AllByXPath(`//select[@id="apply_join_time"]/option`)
			var optionTexts []string
			if optionTexts, err = b.getTexts(options); err != nil {
				return
			}
			var date *util.Time
			for _, optionText := range optionTexts {
				date, err = util.MoesiaParseInLocation(optionText)
				hotelVacancy.Dates = append(hotelVacancy.Dates, date)
			}
			b.page.Back()
		}
		vacancies.List = append(vacancies.List, hotelVacancy)
		b.page.Back()
	}
	return
}

var getTextTexts []string

func (b *Browser) getTexts(multiSelection *agouti.MultiSelection) (texts []string, err error) {
	var elements []*api.Element
	if isTest {
		elements = []*api.Element{nil, nil, nil, nil}
	} else {
		elements, err = multiSelection.Elements()
		if err != nil {
			err = fmt.Errorf("Failed to get elements: %v", err)
			return
		}
	}
	for i, element := range elements {
		var text string
		if isTest {
			if i < len(getTextTexts) {
				text = getTextTexts[i]
			}
		} else if text, err = element.GetText(); err != nil {
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
	filename = fmt.Sprintf("%s.png", file.Name())
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

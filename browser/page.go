package browser

import (
	"github.com/sclevine/agouti"
)

// Page is a mocked struct for agouti.Page
type Page struct {
	*agouti.Page
}

type clicker interface {
	Click() error
}

type clickable struct{}

var clickError error

func (c *clickable) Click() error {
	return clickError
}

var navigateError error

func (page *Page) _Navigate(url string) error {
	if isTest {
		return navigateError
	}
	return page.Navigate(url)
}

func (page *Page) _FindByLink(text string) clicker {
	if isTest {
		return &clickable{}
	}
	return page.FindByLink(text)
}

func (page *Page) _AllByXPath(selector string) *agouti.MultiSelection {
	if isTest {
		return &agouti.MultiSelection{}
	}
	return page.AllByXPath(selector)
}

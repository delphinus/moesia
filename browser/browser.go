package browser

import (
	"fmt"

	"github.com/sclevine/agouti"
)

// Browser has WebDriver in property
type Browser struct {
	driver *agouti.WebDriver
}

// New returns a new instance
func New() (self *Browser, err error) {
	self = &Browser{driver: agouti.PhantomJS()}
	if err = self.driver.Start(); err != nil {
		err = fmt.Errorf("Failed to start driver: %v", err)
		return
	}
	defer self.driver.Stop()
	return
}

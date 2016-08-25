package moesia

import (
	"fmt"

	"github.com/delphinus35/moesia/browser"
	"github.com/urfave/cli"
)

// NewApp returns CLI app by urfave/cli
func NewApp() (app *cli.App) {
	app = cli.NewApp()
	app.Name = "moesia"
	app.Usage = "Explore ths site of ITS"
	app.Version = version
	app.Author = "delphinus"
	app.Email = "delphinus@remora.cx"
	app.Action = action
	return
}

func action(c *cli.Context) (err error) {
	b, err := browser.New()
	if err != nil {
		err = fmt.Errorf("Browser has occurred error: %v", err)
		return
	}
	if err != b.Start() {
		filename, _ := b.Screenshot()
		err = fmt.Errorf("Browser process has errors: %v, saved screenshot: %s", err, filename)
		return
	}
	if err != b.End() {
		err = fmt.Errorf("Browser finish process has errors: %v", err)
		return
	}
	return
}

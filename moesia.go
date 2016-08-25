package moesia

import (
	"fmt"
	"github.com/sclevine/agouti"
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
	driver := agouti.PhantomJS()
	if err = driver.Start(); err != nil {
		err = fmt.Errorf("Failed to start driver: %v", err)
		return
	}
	return
}

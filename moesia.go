package moesia

import (
	"fmt"
	"os"

	"github.com/delphinus35/moesia/browser"
	"github.com/delphinus35/moesia/config"
	"github.com/delphinus35/moesia/mail"
	"github.com/delphinus35/moesia/vacancy"
	"github.com/urfave/cli"
)

var cfg *config.Config

// NewApp returns CLI app by urfave/cli
func NewApp() (app *cli.App) {
	var err error
	if cfg, err = config.New(); err != nil {
		fmt.Fprintf(os.Stderr, "config file cannot be loaded: %v", err)
		os.Exit(1)
	}
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
	var vacancies vacancy.Vacancies
	if vacancies, err = b.Process(); err != nil {
		filename, _ := b.Screenshot()
		err = fmt.Errorf("Browser process has errors: %v, saved screenshot: %s", err, filename)
		return
	}
	m := mail.New(cfg)
	body, err := vacancies.MailBody()
	if err != nil {
		err = fmt.Errorf("failed to create mail body: %v", err)
		return
	}
	if err = m.Send(body); err != nil {
		err = fmt.Errorf("cannot send mail: %v", err)
		return
	}
	if err = b.End(); err != nil {
		err = fmt.Errorf("Browser finish process has errors: %v", err)
		return
	}
	return
}

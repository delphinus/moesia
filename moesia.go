package moesia

import (
	"encoding/json"
	"fmt"
	"os"
	"path"

	"github.com/delphinus35/moesia/browser"
	"github.com/delphinus35/moesia/vacancy"
	"github.com/urfave/cli"
)

type config struct {
}

// Config will have settings of moesia
var Config config

// NewApp returns CLI app by urfave/cli
func NewApp() (app *cli.App) {
	if err := setConfig(); err != nil {
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

func setConfig() error {
	configFilename := fmt.Sprintf("%s/.config/moesia/config.json", os.Getenv("HOME"))
	if _, err := os.Stat(configFilename); err != nil {
		if err := makeInitialConfigFile(configFilename); err != nil {
			return fmt.Errorf("failed to make initial config file: %v", err)
		}
	} else {
		if err := loadConfig(configFilename); err != nil {
			return fmt.Errorf("failed to load config file: %v", err)
		}
	}
	return nil
}

func makeInitialConfigFile(filename string) error {
	if err := os.MkdirAll(path.Dir(filename), 0777); err != nil {
		return fmt.Errorf("failed to mkdir: %v", err)
	}
	var file *os.File
	var err error
	if file, err = os.Create(filename); err != nil {
		return fmt.Errorf("failed to create %s: %v", filename, err)
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	if err := encoder.Encode(Config); err != nil {
		return fmt.Errorf("failed to encode config %v: %v", Config, err)
	}
	return nil
}

func loadConfig(filename string) error {
	var file *os.File
	var err error
	if file, err = os.Open(filename); err != nil {
		return fmt.Errorf("failed to open %s: %v", filename, err)
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&Config); err != nil {
		return fmt.Errorf("failed to decode config: %v", err)
	}
	return nil
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
	fmt.Print(vacancies.String())
	if err = b.End(); err != nil {
		err = fmt.Errorf("Browser finish process has errors: %v", err)
		return
	}
	return
}

package main

import (
	"os"

	"github.com/delphinus/moesia"
)

func main() {
	moesia.NewApp().Run(os.Args)
}

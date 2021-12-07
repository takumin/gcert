package main

import (
	"log"
	"os"

	"github.com/takumin/gcert/internal/app"
	"github.com/takumin/gcert/internal/config"
)

var (
	AppName  string = "gcert"
	Usage    string = "x509 certificate tool"
	Version  string = "unknown"
	Revision string = "unknown"
)

func main() {
	app := app.NewApp(
		AppName,
		Usage,
		Version,
		Revision,
		config.NewConfig(),
	)
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"

	"github.com/takumin/gcert/internal/command/completion"
	"github.com/takumin/gcert/internal/command/selfsign"
	"github.com/takumin/gcert/internal/config"
)

var (
	AppName  string = "gcert"
	Usage    string = "x509 certificate tool"
	Version  string = "unknown"
	Revision string = "unknown"
)

func main() {
	config := config.NewConfig()

	flags := []cli.Flag{}

	cmds := []*cli.Command{}
	cmds = append(cmds, completion.NewCommands(config, flags)...)
	cmds = append(cmds, selfsign.NewCommands(config, flags)...)

	app := &cli.App{
		Name:                 AppName,
		Usage:                Usage,
		Version:              fmt.Sprintf("%s (%s)", Version, Revision),
		Flags:                flags,
		Commands:             cmds,
		EnableBashCompletion: true,
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

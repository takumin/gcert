package app

import (
	"context"
	"fmt"

	"github.com/urfave/cli/v2"

	"github.com/takumin/gcert/internal/app/completion"
	"github.com/takumin/gcert/internal/app/selfsign"
	"github.com/takumin/gcert/internal/config"
)

type App interface {
	Run(args []string) error
	RunContext(ctx context.Context, args []string) error
}

type app struct {
	cli *cli.App
}

func NewApp(name, usage, version, revision string, c *config.Config) App {
	flags := []cli.Flag{}

	cmds := []*cli.Command{}
	cmds = append(cmds, completion.NewCommands(c, flags)...)
	cmds = append(cmds, selfsign.NewCommands(c, flags)...)

	return &app{
		cli: &cli.App{
			Name:                 name,
			Usage:                usage,
			Version:              fmt.Sprintf("%s (%s)", version, revision),
			Flags:                flags,
			Commands:             cmds,
			EnableBashCompletion: true,
		},
	}
}

func (a *app) Run(args []string) error {
	return a.cli.Run(args)
}

func (a *app) RunContext(ctx context.Context, args []string) error {
	return a.cli.RunContext(ctx, args)
}

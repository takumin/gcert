package completion

import (
	"github.com/urfave/cli/v2"

	"github.com/takumin/gcert/internal/command/completion/bash"
	"github.com/takumin/gcert/internal/command/completion/fish"
	"github.com/takumin/gcert/internal/command/completion/powershell"
	"github.com/takumin/gcert/internal/command/completion/zsh"
	"github.com/takumin/gcert/internal/config"
)

func NewCommands(c *config.Config, f []cli.Flag) []*cli.Command {
	cmds := []*cli.Command{}
	cmds = append(cmds, bash.NewCommands(c, f)...)
	cmds = append(cmds, fish.NewCommands(c, f)...)
	cmds = append(cmds, zsh.NewCommands(c, f)...)
	cmds = append(cmds, powershell.NewCommands(c, f)...)

	return []*cli.Command{
		{
			Name:        "completion",
			Usage:       "command completion",
			Subcommands: cmds,
			HideHelp:    true,
		},
	}
}

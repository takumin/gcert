package completion

import (
	"fmt"
	"html/template"
	"os"
	"strings"

	"github.com/urfave/cli/v2"

	"github.com/takumin/gcert/internal/config"
)

const bashCompletion = `
#!/bin/bash

_cli_bash_autocomplete() {
  if [[ "${COMP_WORDS[0]}" != "source" ]]; then
    local cur opts base
    COMPREPLY=()
    cur="${COMP_WORDS[COMP_CWORD]}"
    if [[ "$cur" == "-"* ]]; then
      opts=$( ${COMP_WORDS[@]:0:$COMP_CWORD} ${cur} --generate-bash-completion )
    else
      opts=$( ${COMP_WORDS[@]:0:$COMP_CWORD} --generate-bash-completion )
    fi
    COMPREPLY=( $(compgen -W "${opts}" -- ${cur}) )
    return 0
  fi
}

complete -o bashdefault -o default -o nospace -F _cli_bash_autocomplete {{.}}
`

const zshCompletion = `
#compdef {{.}}

_cli_zsh_autocomplete() {
  local -a opts
  local cur

  cur=${words[-1]}
  if [[ "$cur" == "-"* ]]; then
    opts=("${(@f)$(_CLI_ZSH_AUTOCOMPLETE_HACK=1 ${words[@]:0:#words[@]-1} ${cur} --generate-bash-completion)}")
  else
    opts=("${(@f)$(_CLI_ZSH_AUTOCOMPLETE_HACK=1 ${words[@]:0:#words[@]-1} --generate-bash-completion)}")
  fi

  if [[ "${opts[1]}" != "" ]]; then
    _describe 'values' opts
  else
    _files
  fi

  return
}

compdef _cli_zsh_autocomplete {{.}}
`

const powershellCompletion = `
$fn = $($MyInvocation.MyCommand.Name)
$name = $fn -replace "(.*)\.ps1$", '$1'
Register-ArgumentCompleter -Native -CommandName $name -ScriptBlock {
  param($commandName, $wordToComplete, $cursorPosition)
  $other = "$wordToComplete --generate-bash-completion"
  Invoke-Expression $other | ForEach-Object {
    [System.Management.Automation.CompletionResult]::new($_, $_, 'ParameterValue', $_)
  }
}
`

func NewCommands(c *config.Config, f []cli.Flag) []*cli.Command {
	flags := []cli.Flag{
		&cli.BoolFlag{
			Name:    "bash",
			Aliases: []string{"b"},
			Usage:   "show bash completion script",
		},
		&cli.BoolFlag{
			Name:    "fish",
			Aliases: []string{"f"},
			Usage:   "show fish completion script",
		},
		&cli.BoolFlag{
			Name:    "zsh",
			Aliases: []string{"z"},
			Usage:   "show zsh completion script",
		},
		&cli.BoolFlag{
			Name:    "powershell",
			Aliases: []string{"p"},
			Usage:   "show powershell completion script",
		},
	}
	return []*cli.Command{
		{
			Name:   "completion",
			Usage:  "command completion",
			Hidden: true,
			Flags:  flags,
			Action: func(ctx *cli.Context) error {
				if ctx.Bool("bash") {
					t, err := template.New("bashCompletion").Parse(strings.TrimSpace(bashCompletion) + "\n")
					if err != nil {
						return err
					}
					if err = t.Execute(os.Stdout, ctx.App.Name); err != nil {
						return err
					}
					return nil
				}
				if ctx.Bool("zsh") {
					t, err := template.New(ctx.App.Name).Parse(strings.TrimSpace(zshCompletion) + "\n")
					if err != nil {
						return err
					}
					if err = t.Execute(os.Stdout, ctx.App.Name); err != nil {
						return err
					}
					return nil
				}
				if ctx.Bool("fish") {
					fish, err := ctx.App.ToFishCompletion()
					if err != nil {
						return err
					}
					fmt.Println(fish)
					return nil
				}
				if ctx.Bool("powershell") {
					fmt.Println(strings.TrimSpace(powershellCompletion))
					return nil
				}
				return cli.ShowSubcommandHelp(ctx)
			},
		},
	}
}

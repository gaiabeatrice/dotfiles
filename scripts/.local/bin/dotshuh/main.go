package main

import (
	"log"
	"os/exec"

	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/huh/spinner"
)

type Action int

const (
	Cancel Action = iota
	BrewDump
	BrewInstall
	Stow
)

var theme *huh.Theme = huh.ThemeCatppuccin()

func main() {
	var action Action

	f := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[Action]().
				Value(&action).
				Options(
					huh.NewOption("Dump brew configuration", BrewDump),
					huh.NewOption("Install packages with brew", BrewInstall),
					huh.NewOption("Stow your configuration", Stow),
					huh.NewOption("Nevermind...", Cancel),
				).
				Title("What do you want to do today?"),
		),
	).WithTheme(theme)

	err := f.Run()
	if err != nil {
		log.Fatal(err)
	}

	switch action {
	case BrewDump:
		action := func() {
			execCommand("brew", []string{"bundle", "dump", "-f"})
		}
		spinner.New().Title("Running brew bundle dump -f").Action(action).Run()
	case BrewInstall:
		action := func() {
			execCommand("brew", []string{"bundle"})
		}
		spinner.New().Title("Running brew bundle").Action(action).Run()
	case Stow:
		action := func() {
			execCommand("stow", []string{"*/", "--no-folding"})
		}
		spinner.New().Title("Running stow */ --no-folding").Action(action).Run()
	}
}

func execCommand(name string, args []string) string {
	cmd := exec.Command(name, args...)

	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	return "success"
}

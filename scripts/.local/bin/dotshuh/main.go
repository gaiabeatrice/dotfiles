package main

import (
	"bufio"
	"fmt"
	"log"
	"os/exec"
	"strings"

	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
)

type Action int

const (
	Cancel Action = iota
	BrewDump
	BrewInstall
	Stow
)

var theme *huh.Theme = huh.ThemeCatppuccin()
var highlightColor = theme.Focused.FocusedButton.GetBackground()

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
		execCommand("brew", []string{"bundle", "dump", "-f"})
	case BrewInstall:
		execCommand("brew", []string{"bundle"})
	case Stow:
		execCommand("stow", []string{"*/", "--no-folding"})
	}
}

func execCommand(name string, args []string) string {
	cmd := exec.Command(name, args...)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println(err)
	}

	err = cmd.Start()

	runningCommandMessage :=
		lipgloss.NewStyle().
			Bold(true).
			Foreground(highlightColor).
			SetString("Running command:")

	fmt.Println(runningCommandMessage, name, strings.Join(args, " "))

	if err != nil {
		fmt.Println(err)
	}

	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		m := scanner.Text()
		fmt.Println(m)
	}
	cmd.Wait()

	return "success"
}

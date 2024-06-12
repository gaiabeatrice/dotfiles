package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var docStyle = lipgloss.NewStyle().Margin(1, 2)

type finishedMsg struct{ err error }

type item struct {
	title, desc string
}

type model struct {
	list list.Model
	err  error
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.desc }
func (i item) FilterValue() string { return i.title }

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		h, _ := docStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, 20)

	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "q", "ctrl+c":
			return m, tea.Quit

		case "enter":
			commandDescription := m.list.SelectedItem().(item).Description()
			return m, execCommand(commandDescription)
		}
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd

}

func (m model) View() string {
	return docStyle.Render(m.list.View())
}

func execCommand(command string) tea.Cmd {
	commandSlice := strings.Split(command, " ")
	mainCommand := commandSlice[0]
	sliceLen := len(commandSlice)
	args := commandSlice[1:sliceLen]

	c := exec.Command(mainCommand, args...)
	return tea.ExecProcess(c, func(err error) tea.Msg {
		return finishedMsg{err}
	})
}

func main() {
	items := []list.Item{
		item{title: "Dump brew configuration", desc: "brew bundle dump -f"},
		item{title: "Install packages with brew", desc: "brew bundle"},
		item{title: "Stow files", desc: "stow */ --no-folding"},
	}

	l := list.New(items, list.NewDefaultDelegate(), 0, 0)
	l.Title = "Choose which command to run"
	l.SetFilteringEnabled(false)
	l.SetShowPagination(false)

	m := model{list: l}

	if _, err := tea.NewProgram(m).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}

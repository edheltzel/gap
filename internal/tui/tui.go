// Package tui renders all gap terminal output through Bubble Tea.
package tui

import (
	"fmt"
	"io"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var body = lipgloss.NewStyle().Foreground(lipgloss.Color("252"))

// Print renders text through a one-shot Bubble Tea program and writes it to out.
func Print(out io.Writer, text string) error {
	if out == nil {
		return fmt.Errorf("tui: nil output")
	}
	p := tea.NewProgram(model{text: body.Render(text)}, tea.WithInput(nil), tea.WithOutput(out))
	_, err := p.Run()
	return err
}

type model struct {
	text string
}

func (m model) Init() tea.Cmd {
	return tea.Quit
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, tea.Quit
}

func (m model) View() string {
	return m.text
}

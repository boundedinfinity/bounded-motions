package tui

import (
	"go-motions/model"

	tea "github.com/charmbracelet/bubbletea"
)

func newFooter(root *model.KeyBinding) *footer {
	return &footer{
		Root:    root,
		Current: root,
	}
}

type footer struct {
	Root    *model.KeyBinding
	Current *model.KeyBinding
}

func (this footer) Init() tea.Cmd {
	return nil
}

func (this footer) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	return this, nil
}

func (this footer) View() string {
	path := this.Current.Path()
	var parts []string
	var text string

	for _, part := range path {
		parts = append(parts, part.Name)
	}

	return text
}

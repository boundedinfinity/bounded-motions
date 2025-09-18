package main

import (
	tea "github.com/charmbracelet/bubbletea"
)

func newForm(config *config) form {
	return form{
		config: config,
	}
}

type form struct {
	config *config
	labels []string
}

func (_ form) Init() tea.Cmd {
	return nil
}

func (this form) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	return this, tea.Batch(cmds...)
}

func (this form) View() string {

	return "form"

}

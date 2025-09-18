// https://github.com/charmbracelet/bubbletea/blob/main/examples/composable-views/main.go
package main

import (
	"go-motions/breadcrumb"
	"go-motions/filepicker"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type app struct {
	config     *config
	filepicker tea.Model
	form       tea.Model
	innerStyle lipgloss.Style
	outerStyle lipgloss.Style
	width      int
	height     int
	breadcrumb tea.Model
}

func newApp(config *config) app {
	return app{
		config:     config,
		form:       newForm(config),
		filepicker: newFilePicker(config),
		breadcrumb: breadcrumb.New(),
		outerStyle: lipgloss.NewStyle().MarginTop(1).MarginLeft(2),
		innerStyle: lipgloss.NewStyle().BorderStyle(lipgloss.NormalBorder()),
	}
}

func (this app) Init() tea.Cmd {
	return tea.Batch(
		this.filepicker.Init(),
		this.form.Init(),
	)
}

func (this app) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	var cmd tea.Cmd
	var quitting bool

	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			cmds = append(cmds, tea.Quit)
			quitting = true
		}
	case tea.WindowSizeMsg:
		this.height = msg.Height
		this.width = msg.Width
	case filepicker.DirSelectedMessage:
		cmds = append(cmds, breadcrumb.BreadCrumpUpdateCmd(strings.Split(msg.Path, "/")...))
	}

	if !quitting {
		this.filepicker, cmd = this.filepicker.Update(msg)
		cmds = append(cmds, cmd)

		this.breadcrumb, cmd = this.breadcrumb.Update(msg)
		cmds = append(cmds, cmd)
	}

	return this, tea.Batch(cmds...)
}

func (this app) View() string {
	var s string

	calcWidth := func(p int) int {
		return int(float32(this.width) * float32(p/100))
	}

	s = this.outerStyle.Render(

		lipgloss.JoinVertical(lipgloss.Top,
			this.innerStyle.Width(calcWidth(90)).Render(this.breadcrumb.View()),
			lipgloss.JoinHorizontal(lipgloss.Top,
				this.innerStyle.Width(calcWidth(45)).Render(this.filepicker.View()),
				this.innerStyle.Width(calcWidth(45)).Render(this.filepicker.View()),
			),
		),
	)

	return s
}

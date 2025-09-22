// Package tui https://github.com/charmbracelet/bubbletea/blob/main/examples/composable-views/main.go
package tui

import (
	"fmt"
	"go-motions/breadcrumb"
	"go-motions/model"
	"go-motions/simplelist"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type tui struct {
	config                model.ConfigJson
	root                  *model.KeyBinding
	current               *model.KeyBinding
	navigationBoxStyle    lipgloss.Style
	navigationPathStyle   lipgloss.Style
	navigationOptionStyle lipgloss.Style
	width                 int
	height                int
	bindingPath           *breadcrumb.Model
	bindingOptions        *simplelist.Model
	quitting              bool
	loaded                bool
}

func NewApp(config model.ConfigJson, root *model.KeyBinding) *tui {
	return &tui{
		config:                config,
		root:                  root,
		current:               root,
		bindingPath:           breadcrumb.New(),
		bindingOptions:        simplelist.New(),
		navigationPathStyle:   newStyle(config.Style.NavigationPath),
		navigationOptionStyle: newStyle(config.Style.NavigationOption),
		navigationBoxStyle:    newStyle(config.Style.NavigationBox).BorderStyle(lipgloss.NormalBorder()),
	}
}

func newStyle(jstyle model.ConfigJsonStyle) lipgloss.Style {
	style := lipgloss.NewStyle()

	if jstyle.ForegroundColor != "" {
		style.Foreground(lipgloss.Color(jstyle.ForegroundColor))
	}

	style.Margin(jstyle.Margin).Padding(jstyle.Padding)

	return style
}

func (this tui) Init() tea.Cmd {
	this.bindingPath.Update(newBindingPath(this.current)())
	this.bindingOptions.Update(newBindingOption(this.current))
	return nil
}

func (this tui) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch key := msg.String(); key {
		case "ctrl+c", "q":
			cmds = append(cmds, tea.Quit)
			this.quitting = true
		default:
			if found, ok := this.current.MatchChild(key); ok {
				this.current = found
				_, cmd := this.bindingPath.Update(newBindingPath(found)())
				cmds = append(cmds, cmd)
				_, cmd = this.bindingOptions.Update(newBindingOption(found))
				cmds = append(cmds, cmd)
			}
		}
	case tea.WindowSizeMsg:
		// calcWidth := func(p int) int {
		// 	return int(float32(this.width) * float32(p/100))
		// }

		this.height = msg.Height
		this.width = msg.Width
		// this.navigationBoxStyle.Width(calcWidth(this.width))
		this.navigationBoxStyle.Width(this.width)
		this.loaded = true
	}

	return this, tea.Batch(cmds...)
}

func (this tui) View() string {
	var s string

	if this.quitting {
		return s
	}

	if !this.loaded {
		return "loading..."
	}

	s = this.navigationBoxStyle.Render(
		lipgloss.JoinHorizontal(
			lipgloss.Left,
			this.navigationPathStyle.Render(this.bindingPath.View()),
			this.navigationOptionStyle.Render(this.bindingOptions.View()),
		),
	)

	return s
}

// ////////////////////////////////////////////////////////////////////////////////////
// Utilities
// ////////////////////////////////////////////////////////////////////////////////////
// https://www.w3schools.com/colors/colors_picker.asp

func newBindingPath(binding *model.KeyBinding) tea.Cmd {
	var items []fmt.Stringer

	for _, part := range binding.Path() {
		items = append(items, &bindingPath{*part})
	}

	return breadcrumb.BreadCrumpUpdateCmd(items...)
}

type bindingPath struct {
	binding model.KeyBinding
}

func (this bindingPath) String() string {
	return fmt.Sprintf("%s (%s)", this.binding.Key, this.binding.Name)
}

func newBindingOption(binding *model.KeyBinding) simplelist.ListItemMsg {
	var items []simplelist.ListItem

	for _, child := range binding.Children {
		items = append(items, bindingOption{*child})
	}

	return simplelist.NewMsg(items)
}

type bindingOption struct {
	binding model.KeyBinding
}

func (this bindingOption) String() string {
	return fmt.Sprintf("%s (%s)", this.binding.Key, this.binding.Name)
}

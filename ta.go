// https://github.com/charmbracelet/bubbletea/blob/master/examples/split-editors/main.go
package main

import (
	"fmt"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	initialInputs = 2
	maxInputs     = 6
	minInputs     = 1
	helpHeight    = 5
)

type keymap = struct {
	next, prev, add, remove, quit key.Binding
}

type textAreaView struct {
	width    int
	height   int
	keymap   keymap
	help     help.Model
	textArea textarea.Model
	focus    int
}

func newTextArea() textAreaView {
	endOfBufferStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("235"))

	ta := textarea.New()
	ta.Prompt = ""
	ta.Placeholder = "Type something"
	ta.ShowLineNumbers = true
	ta.Cursor.Style = lipgloss.NewStyle().
		Foreground(lipgloss.Color("212"))
	ta.FocusedStyle.Placeholder = lipgloss.NewStyle().
		Foreground(lipgloss.Color("99"))
	ta.BlurredStyle.Placeholder = lipgloss.NewStyle().
		Foreground(lipgloss.Color("238"))
	ta.FocusedStyle.CursorLine = lipgloss.NewStyle().
		Background(lipgloss.Color("57")).
		Foreground(lipgloss.Color("230"))
	ta.FocusedStyle.Base = lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("238"))
	ta.BlurredStyle.Base = lipgloss.NewStyle().
		Border(lipgloss.HiddenBorder())
	ta.FocusedStyle.EndOfBuffer = endOfBufferStyle
	ta.BlurredStyle.EndOfBuffer = endOfBufferStyle
	ta.KeyMap.DeleteWordBackward.SetEnabled(false)
	ta.KeyMap.LineNext = key.NewBinding(key.WithKeys("down"))
	ta.KeyMap.LinePrevious = key.NewBinding(key.WithKeys("up"))
	ta.Blur()

	m := textAreaView{
		textArea: ta,
		help:     help.New(),
		keymap: keymap{
			next: key.NewBinding(
				key.WithKeys("tab"),
				key.WithHelp("tab", "next"),
			),
			prev: key.NewBinding(
				key.WithKeys("shift+tab"),
				key.WithHelp("shift+tab", "prev"),
			),
			add: key.NewBinding(
				key.WithKeys("ctrl+n"),
				key.WithHelp("ctrl+n", "add an editor"),
			),
			remove: key.NewBinding(
				key.WithKeys("ctrl+w"),
				key.WithHelp("ctrl+w", "remove an editor"),
			),
			quit: key.NewBinding(
				key.WithKeys("esc", "ctrl+c"),
				key.WithHelp("esc", "quit"),
			),
		},
	}

	m.textArea.Focus()
	m.updateKeybindings()
	return m
}

func (_ textAreaView) Init() tea.Cmd {
	return textarea.Blink
}

func (this textAreaView) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		fmt.Printf("[unkown key: %v]", msg)
		switch {
		case key.Matches(msg, this.keymap.quit):
			this.textArea.Blur()
			return this, tea.Quit
		case key.Matches(msg, this.keymap.next):
			this.textArea.Blur()
			this.focus++
			cmd := this.textArea.Focus()
			cmds = append(cmds, cmd)
		case key.Matches(msg, this.keymap.prev):
			this.textArea.Blur()
			this.focus--
			cmd := this.textArea.Focus()
			cmds = append(cmds, cmd)
		case key.Matches(msg, this.keymap.add):
			// nothing
		case key.Matches(msg, this.keymap.remove):
			// nothing
		}
	case tea.WindowSizeMsg:
		this.height = msg.Height
		this.width = msg.Width
	}

	this.updateKeybindings()
	this.sizeInputs()

	// Update all textareas
	this.textArea.Update(msg)

	return this, tea.Batch(cmds...)
}

func (this *textAreaView) sizeInputs() {
	this.textArea.SetWidth(this.width)
	this.textArea.SetHeight(this.height - helpHeight)
}

func (this *textAreaView) updateKeybindings() {
	// this.keymap.add.SetEnabled(len(this.textArea) < maxInputs)
	// this.keymap.remove.SetEnabled(len(this.textArea) > minInputs)
}

func (this textAreaView) View() string {
	help := this.help.ShortHelpView([]key.Binding{
		this.keymap.next,
		this.keymap.prev,
		this.keymap.add,
		this.keymap.remove,
		this.keymap.quit,
	})

	return this.textArea.View() + "\n\n" + help
	// return lipgloss.JoinHorizontal(lipgloss.Top, views...) + "\n\n" + help
}

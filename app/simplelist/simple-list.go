package simplelist

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

func New() *Model {
	return NewWithConfig(ModelConfig{
		Sep: "\n",
	})
}

func NewWithConfig(config ModelConfig) *Model {
	m := &Model{
		Items:  []ListItem{},
		Config: config,
	}

	if m.Config.Sep == "" {
		m.Config.Sep = "\n"
	}

	return m
}

type ModelConfig struct {
	Sep string
}

type Model struct {
	Items  []ListItem
	Config ModelConfig
}

func (_ Model) Init() tea.Cmd {
	return nil
}

func (this *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case ListItemMsg:
		this.Items = msg.Items
	}

	return this, nil
}

func (this Model) View() string {
	var items []string

	for _, item := range this.Items {
		items = append(items, fmt.Sprintf("- %s", item.String()))
	}

	return strings.Join(items, this.Config.Sep)
}

type ListItem fmt.Stringer

type ListItemMsg struct {
	Items []ListItem
}

func NewMsg(items []ListItem) ListItemMsg {
	return ListItemMsg{Items: items}
}

package breadcrumb

import (
	tea "github.com/charmbracelet/bubbletea"
)

type BreadCrumbUpdateMessage struct {
	Items []string
}

func BreadCrumpUpdateCmd(items ...string) tea.Cmd {
	return func() tea.Msg {
		return BreadCrumbUpdateMessage{
			Items: items,
		}
	}
}

func BreadCrumpClearCmd() tea.Cmd {
	return BreadCrumpUpdateCmd()
}

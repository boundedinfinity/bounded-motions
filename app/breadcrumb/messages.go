package breadcrumb

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type BreadCrumbUpdateMessage struct {
	Items []fmt.Stringer
}

func BreadCrumpUpdateCmd(items ...fmt.Stringer) tea.Cmd {
	return func() tea.Msg {
		return BreadCrumbUpdateMessage{
			Items: items,
		}
	}
}

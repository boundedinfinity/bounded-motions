package filepicker

import tea "github.com/charmbracelet/bubbletea"

type DirSelectedMessage struct {
	Path string
}

func DirSelectedCmd(path string) tea.Cmd {
	return func() tea.Msg {
		return DirSelectedMessage{
			Path: path,
		}
	}
}

type FileSelectedMessage struct {
	Path string
}

func FileSelectedCmd(path string) tea.Cmd {
	return func() tea.Msg {
		return DirSelectedMessage{
			Path: path,
		}
	}
}

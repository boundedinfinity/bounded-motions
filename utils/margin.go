package utils

import "github.com/charmbracelet/lipgloss"

func MarginHorizontal(s lipgloss.Style, m int) {
	s.MarginLeft(m)
	s.MarginRight(m)
}

func MarginVertical(s lipgloss.Style, m int) {
	s.MarginTop(m)
	s.MarginBottom(m)
}

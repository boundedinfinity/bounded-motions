package breadcrumb

import (
	"go-motions/utils"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	items          []string
	delimiter      string
	itemStyle      lipgloss.Style
	delimiterStyle lipgloss.Style
	frameStyle     lipgloss.Style
}

func New() Model {
	return Model{
		items:          []string{},
		delimiter:      ">",
		itemStyle:      defaultItemStyle(),
		delimiterStyle: defaultDelimterStyle(),
		frameStyle:     defaultFrameStyle(),
	}
}

func (this Model) Items(items ...string) Model {
	return utils.SetAndReturn(this, &this.items, items)
}

func (this Model) Delimiter(delimiter string) Model {
	return utils.SetAndReturn(this, &this.delimiter, delimiter)
}

func (this Model) ItemStyle(style lipgloss.Style) Model {
	return utils.SetAndReturn(this, &this.itemStyle, style)
}

func (this Model) DelimiterStyle(style lipgloss.Style) Model {
	return utils.SetAndReturn(this, &this.delimiterStyle, style)
}

func (this Model) FrameStyle(style lipgloss.Style) Model {
	return utils.SetAndReturn(this, &this.frameStyle, style)
}

func (_ Model) Init() tea.Cmd {
	return nil
}

func (this Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case BreadCrumbUpdateMessage:
		this.items = msg.Items
	}

	return this, nil
}

func (this Model) View() string {
	var items []string

	for _, item := range this.items {
		items = append(items, this.itemStyle.Render(item))
	}

	delimiter := this.delimiterStyle.Render(this.delimiter)

	return strings.Join(items, delimiter)
}

// https://www.w3schools.com/colors/colors_picker.asp
var (
	_FG_COLOR = "#333399"
)

func defaultItemStyle() lipgloss.Style {
	return lipgloss.NewStyle().Foreground(lipgloss.Color(_FG_COLOR))
}

func defaultDelimterStyle() lipgloss.Style {
	return lipgloss.NewStyle().Foreground(lipgloss.Color(_FG_COLOR)).PaddingLeft(1).PaddingRight(1)
}

func defaultFrameStyle() lipgloss.Style {
	return lipgloss.NewStyle().Foreground(lipgloss.Color(_FG_COLOR))
}

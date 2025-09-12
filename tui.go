package bounded_motions

import (
	"github.com/rivo/tview"
)

// ////////////////////////////////////////////////////////////////////////////
// BoundedMotions
// ////////////////////////////////////////////////////////////////////////////

func New() *BoundedMotions {
	return &BoundedMotions{
		tui: &Tui{},
	}
}

type BoundedMotions struct {
	tui *Tui
}

// ////////////////////////////////////////////////////////////////////////////
// Tui
// ////////////////////////////////////////////////////////////////////////////

type Tui struct {
	app    *tview.Application
	header *tview.TextView
	footer *tview.TextView
}

func (this *Tui) Run() error {
	newPrimitive := func(text string) *tview.TextView {
		return tview.NewTextView().
			SetTextAlign(tview.AlignLeft).
			SetText(text)
	}

	this.app = tview.NewApplication()
	this.header = newPrimitive("Header")
	this.footer = newPrimitive("Footer")

	grid := tview.NewGrid().
		SetRows(3, 3).
		SetColumns(0).
		SetBorders(true).
		//AddItem(p tview.Primitive, row int, column int, rowSpan int, colSpan int, minGridHeight int, minGridWidth int, focus bool)
		AddItem(this.header, 0, 0, 1, 3, 0, 0, false).
		AddItem(this.footer, 1, 0, 1, 3, 0, 0, false)

	return tview.NewApplication().SetRoot(grid, true).Run()
}

// ////////////////////////////////////////////////////////////////////////////
// Run
// ////////////////////////////////////////////////////////////////////////////

func Run() {
	bm := New()

	if err := bm.tui.Run(); err != nil {
		panic(err)
	}
}

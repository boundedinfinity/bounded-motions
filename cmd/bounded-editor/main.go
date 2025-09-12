// Demo code for the TextArea primitive.
package main

import (
	"fmt"
	"strings"

	"github.com/boundedinfinity/motions/motions"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func main() {
	ms := motions.Create()
	app := tview.NewApplication()

	textArea := tview.NewTextArea().
		SetWrap(false).
		SetPlaceholder("Enter text here...")
	textArea.SetTitle("Text Area").SetBorder(true)

	result0 := ms.Initial()
	motionView := tview.NewTextView().
		SetDynamicColors(true).
		SetText(fmt.Sprintf("  %v : %v",
			strings.Join(result0.Path, " > "),
			strings.Join(result0.Options, ", "),
		))
	position := tview.NewTextView().
		SetDynamicColors(true).
		SetTextAlign(tview.AlignRight)
	pages := tview.NewPages()

	updateInfos := func() {
		fromRow, fromColumn, toRow, toColumn := textArea.GetCursor()
		if fromRow == toRow && fromColumn == toColumn {
			position.SetText(fmt.Sprintf("Row: [yellow]%d[white], Column: [yellow]%d ", fromRow, fromColumn))
		} else {
			position.SetText(fmt.Sprintf("[red]From[white] Row: [yellow]%d[white], Column: [yellow]%d[white] - [red]To[white] Row: [yellow]%d[white], To Column: [yellow]%d ", fromRow, fromColumn, toRow, toColumn))
		}
	}

	updateInfos()

	// AddItem(p Primitive, row, column, rowSpan, colSpan, minGridHeight, minGridWidth int, focus bool) *Grid {

	mainView := tview.NewGrid().
		SetRows(0, 3).
		AddItem(textArea, 0, 0, 1, 3, 0, 0, true).
		AddItem(motionView, 1, 0, 1, 1, 0, 0, false).
		AddItem(position, 1, 1, 1, 1, 0, 0, false)

	tview.NewTextView()

	pages.AddAndSwitchToPage("main", mainView, true)

	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		// helpInfo.SetText(fmt.Sprintf("rune: %v, key:%v, mod: %v, name: %v", string(event.Rune()), tcell.KeyNames[event.Key()], event.Modifiers(), event.Name()))
		result := ms.HandleEvent(motions.MotionEvent{Key: event})
		text := fmt.Sprintf("  %v : %v",
			strings.Join(result.Path, " > "),
			strings.Join(result.Options, ", "),
		)

		if result.Message != "" {
			text += fmt.Sprintf(` [red]%s`, result.Message)
		}

		motionView.SetText(text)

		if event.Rune() == 'q' {
			textArea.Select(5, 20)
			return nil
		}

		return event
	})

	if err := app.SetRoot(pages, true).EnableMouse(true).EnablePaste(true).Run(); err != nil {
		panic(err)
	}
}

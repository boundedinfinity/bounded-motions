package main

import (
	"fmt"
	"os"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func main() {
	config, err := loadConfig("./config.json")
	handleErr(err)

	fmt.Printf("%+v", config)

	newPrimitive := func(text string) *tview.TextView {
		return tview.NewTextView().
			SetTextAlign(tview.AlignBottom).
			SetText(text)
	}

	menu := newPrimitive("Menu")
	// main := newPrimitive("Main content")
	main := tview.NewCheckbox().SetLabel("Test")

	sideBar := newPrimitive("Side Bar")

	grid := tview.NewGrid().
		SetRows(3, 0, 3).
		SetColumns(30, 0, 30).
		SetBorders(true).
		AddItem(newPrimitive("Header"), 0, 0, 1, 3, 0, 0, false).
		AddItem(newPrimitive("Footer"), 2, 0, 1, 3, 0, 0, false)

	// Layout for screens narrower than 100 cells (menu and side bar are hidden).
	grid.AddItem(menu, 0, 0, 0, 0, 0, 0, false).
		AddItem(main, 1, 0, 1, 3, 0, 0, true).
		AddItem(sideBar, 0, 0, 0, 0, 0, 0, false)

	// Layout for screens wider than 100 cells.
	grid.AddItem(menu, 1, 0, 1, 1, 0, 100, false).
		AddItem(main, 1, 1, 1, 1, 0, 100, false).
		AddItem(sideBar, 1, 2, 1, 1, 0, 100, false)

	app := tview.NewApplication().SetRoot(grid, true).EnableMouse(true)

	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		sideBar.SetText(event.Name())
		return event
	})

	if err := app.Run(); err != nil {
		panic(err)
	}
}

func handleErr(err error) {
	if err != nil {
		fmt.Println("Error while running program:", err)
		os.Exit(1)
	}
}

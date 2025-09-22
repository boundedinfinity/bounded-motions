package main

import (
	"fmt"
	"go-motions/model"
	"go-motions/tui"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	var config model.ConfigJson
	if err := model.LoadConfig("./../config.json", &config); err != nil {
		handleErr(err)
	}

	keyBinding := new(model.KeyBinding)
	if err := model.LoadKeyBindings(config, keyBinding); err != nil {
		handleErr(err)
	}

	tui := tui.NewApp(config, keyBinding)
	program := tea.NewProgram(tui)

	if _, err := program.Run(); err != nil {
		panic(err)
	}
}

func handleErr(err error) {
	if err != nil {
		fmt.Println("Error while running program:", err)
		os.Exit(1)
	}
}

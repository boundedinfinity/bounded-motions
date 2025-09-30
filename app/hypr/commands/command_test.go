package commands_test

import (
	"go-motions/hypr/commands"
	"testing"
)

func TestCommand(t *testing.T) {
	exec := commands.Exec{Command: "ghostty"}

	if result := exec.String(); result == "" {
		t.Errorf("is empty")
	}

	// if result := hypr.Commands.Dispatch.Window.Pass(hypr.Windows.Id.Class("abc")).String(); result == "" {
	// 	t.Errorf("is empty")
	// }

	// if result := hypr.Commands.Dispatch.Window.Shortcut(hypr.ModKeys.Super, "a", hypr.Windows.Id.Title("abc")).String(); result == "" {
	// 	t.Errorf("is empty")
	// }

	// if result := hypr.Commands.Dispatch.Window.KeyState(hypr.ModKeys.Super, "a", hypr.KeyStates.Down, hypr.Windows.Id.Class("abc")).String(); result == "" {
	// 	t.Errorf("is empty")
	// }

	// if result := hypr.Commands.Dispatch.Window.KillActive().String(); result == "" {
	// 	t.Errorf("is empty")
	// }
}

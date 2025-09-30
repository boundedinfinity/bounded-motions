package commands_test

import (
	"fmt"
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

func TestParseMonitors(t *testing.T) {
	result, err := commands.ParseResult(commands.Monitors{}, `Monitor HDMI-A-1 (ID 0):
3840x2160@60.00000 at 0x-2160
description: Samsung Electric Company LF32TU87 HCPR601907
make: Samsung Electric Company
model: LF32TU87
physical size (mm): 700x400
serial: HCPR601907
active workspace: 1 (1)
special workspace: 0 ()
reserved: 0 26 0 0
scale: 1.00
transform: 0
focused: no
dpmsStatus: 1
vrr: false
solitary: 0
solitaryBlockedBy: windowed mode,missing candidate
activelyTearing: false
tearingBlockedBy: next frame is not torn,user settings,missing candidate
directScanoutTo: 0
directScanoutBlockedBy: user settings,missing candidate
disabled: false
currentFormat: XRGB8888
mirrorOf: none
availableModes: 3840x2160@60.00Hz 3840x2160@60.00Hz 3840x2160@59.94Hz 3840x2160@50.00Hz 3840x2160@30.00Hz 3840x2160@30.00Hz 3840x2160@29.97Hz 1920x1200@60.00Hz 1920x1080@60.00Hz 1920x1080@60.00Hz 1920x1080@59.94Hz 1920x1080@50.00Hz 1920x1080@50.00Hz 1600x1200@60.00Hz 1680x1050@59.88Hz 1600x900@60.00Hz 1280x1024@75.03Hz 1280x1024@60.02Hz 1440x900@59.90Hz 1280x800@59.91Hz 1152x864@75.00Hz 1280x720@60.00Hz 1280x720@60.00Hz 1280x720@59.94Hz 1280x720@50.00Hz 1024x768@75.03Hz 1024x768@70.07Hz 1024x768@60.00Hz 832x624@74.55Hz 800x600@75.00Hz 800x600@72.19Hz 800x600@60.32Hz 800x600@56.25Hz 720x576@50.00Hz 720x480@60.00Hz 720x480@59.94Hz 640x480@75.00Hz 640x480@72.81Hz 640x480@66.67Hz 640x480@60.00Hz 640x480@59.94Hz 720x400@70.08Hz`)

	fmt.Print(result)
	fmt.Print(err)
}

package commands

// /////////////////////////////////////////////////////////////////////////////

var _ HyprCommand = &CursorPos{}

type CursorPos struct{}

func (this CursorPos) String() string {
	return "cursorpos"
}

func (_ CursorPos) hyprCommand() HyprCommandType {
	return HyprCommandTypes.CursorPos
}

package commands

import "fmt"

// /////////////////////////////////////////////////////////////////////////////
// https://wiki.hypr.land/Configuring/Using-hyprctl/#setcursor

var _ HyprCommand = &SetCursor{}

type SetCursor struct {
	Fontname string
	Size     int
}

func (this SetCursor) String() string {
	return fmt.Sprintf("setcursor %s %d", this.Fontname, this.Size)
}

func (_ SetCursor) hyprCommand() HyprCommandType {
	return HyprCommandTypes.SetError
}

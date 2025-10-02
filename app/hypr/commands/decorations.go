package commands

import (
	"fmt"
	"go-motions/hypr/model"
)

// /////////////////////////////////////////////////////////////////////////////
// https://wiki.hypr.land/Configuring/Using-hyprctl/#info

var _ HyprCommand = &Decorations{}

type Decorations struct {
	Window model.HyprWindowId
}

func (this Decorations) String() string {
	return fmt.Sprintf("decorations %s", this.Window)
}

func (_ Decorations) hyprCommand() HyprCommandType {
	return HyprCommandTypes.Decorations
}

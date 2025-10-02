package commands

import "fmt"

// /////////////////////////////////////////////////////////////////////////////
// https://wiki.hypr.land/Configuring/Using-hyprctl/#notify

var _ HyprCommand = &DismissNotify{}

type DismissNotify struct {
	Count int
}

func (this DismissNotify) String() string {
	return fmt.Sprintf("dismissnotify %d", this.Count)
}

func (_ DismissNotify) hyprCommand() HyprCommandType {
	return HyprCommandTypes.DismissNotify
}

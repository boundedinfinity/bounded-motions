package commands

import "fmt"

// /////////////////////////////////////////////////////////////////////////////
// https://wiki.hypr.land/Configuring/Using-hyprctl/#switchxkblayout

var _ HyprCommand = &SwitchXkbLayout{}

type SwitchXkbLayout struct {
	Device  string
	Command string
}

func (this SwitchXkbLayout) String() string {
	return fmt.Sprintf("switchxkblayout %s %s", this.Device, this.Command)
}

func (_ SwitchXkbLayout) hyprCommand() HyprCommandType {
	return HyprCommandTypes.SwitchXkbLayout
}

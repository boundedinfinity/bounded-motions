package commands

import "fmt"

// /////////////////////////////////////////////////////////////////////////////
// https://wiki.hypr.land/Configuring/Using-hyprctl/#seterror

var _ HyprCommand = &SetError{}

type SetError struct {
	Err     error
	Disable bool
}

func (this SetError) String() string {
	if this.Disable {
		return "seterror disabled"
	}

	return fmt.Sprintf("seterror %s", this.Err)
}

func (_ SetError) hyprCommand() HyprCommandType {
	return HyprCommandTypes.SetError
}

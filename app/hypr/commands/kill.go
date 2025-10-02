package commands

// /////////////////////////////////////////////////////////////////////////////
// https://wiki.hypr.land/Configuring/Using-hyprctl/#kill

var _ HyprCommand = &Kill{}

type Kill struct {
	Command string
}

func (this Kill) String() string {
	return "kill"
}

func (_ Kill) hyprCommand() HyprCommandType {
	return HyprCommandTypes.Kill
}

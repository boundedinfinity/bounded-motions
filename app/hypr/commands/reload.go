package commands

// /////////////////////////////////////////////////////////////////////////////
// https://wiki.hypr.land/Configuring/Using-hyprctl/#reload

var _ HyprCommand = &Reload{}

type Reload struct {
	Command string
}

func (this Reload) String() string {
	return "reload"
}

func (_ Reload) hyprCommand() HyprCommandType {
	return HyprCommandTypes.Version
}

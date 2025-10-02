package commands

// /////////////////////////////////////////////////////////////////////////////
// https://wiki.hypr.land/Configuring/Using-hyprctl/#info

var _ HyprCommand = &Binds{}

type Binds struct{}

func (this Binds) String() string {
	return "binds"
}

func (_ Binds) hyprCommand() HyprCommandType {
	return HyprCommandTypes.Binds
}

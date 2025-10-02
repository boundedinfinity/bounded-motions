package commands

// /////////////////////////////////////////////////////////////////////////////
// https://wiki.hypr.land/Configuring/Using-hyprctl/#info

var _ HyprCommand = &Splash{}

type Splash struct{}

func (this Splash) String() string {
	return "splash"
}

func (_ Splash) hyprCommand() HyprCommandType {
	return HyprCommandTypes.Splash
}

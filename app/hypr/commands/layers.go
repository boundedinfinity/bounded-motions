package commands

// /////////////////////////////////////////////////////////////////////////////
// https://wiki.hypr.land/Configuring/Using-hyprctl/#info

var _ HyprCommand = &Layers{}

type Layers struct{}

func (this Layers) String() string {
	return "layers"
}

func (_ Layers) hyprCommand() HyprCommandType {
	return HyprCommandTypes.Layers
}

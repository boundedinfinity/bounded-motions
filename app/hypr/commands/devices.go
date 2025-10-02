package commands

// /////////////////////////////////////////////////////////////////////////////
// https://wiki.hypr.land/Configuring/Using-hyprctl/#info

var _ HyprCommand = &Devices{}

type Devices struct{}

func (this Devices) String() string {
	return "devices"
}

func (_ Devices) hyprCommand() HyprCommandType {
	return HyprCommandTypes.Devices
}

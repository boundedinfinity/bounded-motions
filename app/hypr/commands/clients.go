package commands

// /////////////////////////////////////////////////////////////////////////////
// https://wiki.hypr.land/Configuring/Using-hyprctl/#info

var _ HyprCommand = &Clients{}

type Clients struct{}

func (this Clients) String() string {
	return "clients"
}

func (_ Clients) hyprCommand() HyprCommandType {
	return HyprCommandTypes.Clients
}

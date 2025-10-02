package commands

// /////////////////////////////////////////////////////////////////////////////

var _ HyprCommand = &Animations{}

type Animations struct{}

func (this Animations) String() string {
	return "animations"
}

func (_ Animations) hyprCommand() HyprCommandType {
	return HyprCommandTypes.Animations
}

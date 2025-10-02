package commands

// /////////////////////////////////////////////////////////////////////////////
// https://wiki.hypr.land/Configuring/Using-hyprctl/#info

var _ HyprCommand = &WorkspaceRules{}

type WorkspaceRules struct{}

func (this WorkspaceRules) String() string {
	return "workspacerules"
}

func (_ WorkspaceRules) hyprCommand() HyprCommandType {
	return HyprCommandTypes.WorkspaceRules
}

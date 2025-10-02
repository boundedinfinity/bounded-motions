package commands

// /////////////////////////////////////////////////////////////////////////////
// https://wiki.hypr.land/Configuring/Using-hyprctl/#info

var _ HyprCommand = &ActiveWorkspace{}

type ActiveWorkspace struct{}

func (this ActiveWorkspace) String() string {
	return "activeworkspace"
}

func (_ ActiveWorkspace) hyprCommand() HyprCommandType {
	return HyprCommandTypes.ActiveWorkspace
}

// /////////////////////////////////////////////////////////////////////////////

var _ HyperCommandResult = &ActiveWorkspaceResult{}

type ActiveWorkspaceResult struct {
	Id              string
	MonitorName     string
	MonitorId       string
	WindowCount     int
	HasFullScreen   int
	LastWindow      string
	LastWindowTitle string
	IsPersistent    bool
}

func (this ActiveWorkspaceResult) String() string {
	return "active workspace: [" + this.Id + "]"
}

func (this *ActiveWorkspaceResult) hyperCommandResult() {}

func parseActiveWorkspace(input string) (*ActiveWorkspaceResult, error) {
	var result ActiveWorkspaceResult
	var err error

	workspacesMatch := workspacesLineRe.FindAllStringSubmatch(input, -1)
	monitorIdMatch := workspacesMonitorRe.FindAllStringSubmatch(input, -1)
	windowsMatch := workspacesWindowsRe.FindAllStringSubmatch(input, -1)
	fullscreenMatch := workspacesFullscreenRe.FindAllStringSubmatch(input, -1)
	lastWindowMatch := workspacesLastWindowRe.FindAllStringSubmatch(input, -1)
	lastTitleMatch := workspacesLastTitleRe.FindAllStringSubmatch(input, -1)
	persistentMatch := workspacesPersistentRe.FindAllStringSubmatch(input, -1)

	result.Id = getSubmatch(workspacesMatch[0], 1)
	result.MonitorName = getSubmatch(workspacesMatch[0], 2)
	result.MonitorId = getSubmatch(monitorIdMatch[0], 1)
	result.WindowCount = getSubmatchInt(windowsMatch[0], 1)
	result.HasFullScreen = getSubmatchInt(fullscreenMatch[0], 1)
	result.LastWindow = getSubmatch(lastWindowMatch[0], 1)
	result.LastWindowTitle = getSubmatch(lastTitleMatch[0], 1)
	result.IsPersistent = getSubmatchBool(persistentMatch[0], 1)

	return &result, err
}

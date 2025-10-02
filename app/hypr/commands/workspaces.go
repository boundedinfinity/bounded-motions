package commands

import (
	"regexp"
	"strings"
)

// /////////////////////////////////////////////////////////////////////////////
// https://wiki.hypr.land/Configuring/Using-hyprctl/#info

var _ HyprCommand = &Monitors{}

type Workspaces struct{}

func (this Workspaces) String() string {
	return "workspaces"
}

func (_ Workspaces) hyprCommand() HyprCommandType {
	return HyprCommandTypes.Workspaces
}

// /////////////////////////////////////////////////////////////////////////////

var _ HyperCommandResult = &WorkspacesResult{}

type WorkspacesResult []*WorkspaceResult

func (this WorkspacesResult) String() string {
	var result []string

	for _, workspace := range this {
		result = append(result, workspace.Id)
	}

	return "workspaces: [" + strings.Join(result, ", ") + "]"
}

func (_ WorkspacesResult) hyperCommandResult() {}

type WorkspaceResult struct {
	Id              string
	MonitorName     string
	MonitorId       string
	WindowCount     int
	HasFullScreen   int
	LastWindow      string
	LastWindowTitle string
	IsPersistent    bool
}

var (
	workspacesLineRe       = regexp.MustCompile(`\((\d+)\) on monitor (.+):`)
	workspacesMonitorRe    = regexp.MustCompile(`monitorID:(.*)\n`)
	workspacesWindowsRe    = regexp.MustCompile(`windows:(.*)\n`)
	workspacesFullscreenRe = regexp.MustCompile(`hasfullscreen:(.*)\n`)
	workspacesLastWindowRe = regexp.MustCompile(`lastwindow:(.*)\n`)
	workspacesLastTitleRe  = regexp.MustCompile(`lastwindowtitle:(.*)\n`)
	workspacesPersistentRe = regexp.MustCompile(`ispersistent:(.*)\n`)
)

func parseWorkspaces(input string) (WorkspacesResult, error) {
	var results WorkspacesResult
	var err error

	workspacesMatch := workspacesLineRe.FindAllStringSubmatch(input, -1)
	monitorIdMatch := workspacesMonitorRe.FindAllStringSubmatch(input, -1)
	windowsMatch := workspacesWindowsRe.FindAllStringSubmatch(input, -1)
	fullscreenMatch := workspacesFullscreenRe.FindAllStringSubmatch(input, -1)
	lastWindowMatch := workspacesLastWindowRe.FindAllStringSubmatch(input, -1)
	lastTitleMatch := workspacesLastTitleRe.FindAllStringSubmatch(input, -1)
	persistentMatch := workspacesPersistentRe.FindAllStringSubmatch(input, -1)

	for i := range workspacesMatch {
		var result WorkspaceResult

		result.Id = getSubmatch(workspacesMatch[i], 1)
		result.MonitorName = getSubmatch(workspacesMatch[i], 2)
		result.MonitorId = getSubmatch(monitorIdMatch[i], 1)
		result.WindowCount = getSubmatchInt(windowsMatch[i], 1)
		result.HasFullScreen = getSubmatchInt(fullscreenMatch[i], 1)
		result.LastWindow = getSubmatch(lastWindowMatch[i], 1)
		result.LastWindowTitle = getSubmatch(lastTitleMatch[i], 1)
		result.IsPersistent = getSubmatchBool(persistentMatch[i], 1)

		results = append(results, &result)
	}

	return results, err
}

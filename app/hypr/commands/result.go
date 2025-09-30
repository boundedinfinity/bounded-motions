package commands

import (
	"regexp"
	"strings"
)

type HyperCommandResult interface {
	String() string
	hyperCommandResult() // discriminator
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

// /////////////////////////////////////////////////////////////////////////////

var _ HyperCommandResult = &VersionResult{}

type VersionResult struct {
	Version      string
	Commit       string
	Branch       string
	Date         string
	Tag          string
	Dependencies map[string]string
	Flags        []string
}

func (this VersionResult) String() string {
	return this.Version
}

func (_ VersionResult) hyperCommandResult() {}

var (
	versionVersion = regexp.MustCompile(`([\d\.]+) built from branch (.*) at commit (.+)\s+\(`)
	versionDate    = regexp.MustCompile(`Date:(.*)\n`)
)

func parseVersion(input string) (*VersionResult, error) {
	var result VersionResult

	versionMatch := versionVersion.FindAllStringSubmatch(input, -1)
	dateMatch := versionDate.FindAllStringSubmatch(input, -1)

	result.Version = getSubmatch(versionMatch[0], 1)
	result.Branch = getSubmatch(versionMatch[0], 2)
	result.Commit = getSubmatch(versionMatch[0], 3)
	result.Date = getSubmatch(dateMatch[0], 1)

	return &result, nil
}

// /////////////////////////////////////////////////////////////////////////////

var _ HyperCommandResult = &MonitorsResult{}

type MonitorsResult []*MonitorResult

func (this MonitorsResult) String() string {
	var result []string

	for _, monitor := range this {
		result = append(result, monitor.Name)
	}

	return "monitors: [" + strings.Join(result, ", ") + "]"
}

func (_ MonitorsResult) hyperCommandResult() {}

type MonitorResult struct {
	Name             string
	Id               string
	Resolution       MonitorResolution
	Position         MonitorPosistion
	Description      string
	Make             string
	Model            string
	Serial           string
	Size             MonitorSize
	ActiveWorkspace  string
	SpecialWorkspace string
}

type MonitorResolution struct {
	Width       string
	Height      string
	RefreshRate string
}

type MonitorPosistion struct {
	X string
	Y string
}

type MonitorSize struct {
	Width  string
	Height string
}

var (
	monitorRe          = regexp.MustCompile(`Monitor\s(.+)\s\(ID (.+)\).*\n`)
	resolutionRe       = regexp.MustCompile(`(\d+)x(\d+)@([\d\.]+)\sat\s([-\d]+)x([-\d]+)\n`)
	descriptionRe      = regexp.MustCompile(`description:(.*)\n`)
	makeRe             = regexp.MustCompile(`make:(.*)\n`)
	modelRe            = regexp.MustCompile(`model:(.*)\n`)
	physicalSizeRe     = regexp.MustCompile(`physical size \(mm\):\s+(\d+)x(\d+).*\n`)
	serialRe           = regexp.MustCompile(`serial:(.*)\n`)
	activeWorkspaceRe  = regexp.MustCompile(`active workspace:(.*)\(.*\n`)
	specialWorkspaceRe = regexp.MustCompile(`special workspace:(.*)\(.*\n`)
)

func parseMonitors(input string) (MonitorsResult, error) {
	var results MonitorsResult
	var err error

	monitorMatch := monitorRe.FindAllStringSubmatch(input, -1)
	resolutionMatch := resolutionRe.FindAllStringSubmatch(input, -1)
	descriptionMatch := descriptionRe.FindAllStringSubmatch(input, -1)
	makeMatch := makeRe.FindAllStringSubmatch(input, -1)
	modelMatch := modelRe.FindAllStringSubmatch(input, -1)
	physicalSizeMatch := physicalSizeRe.FindAllStringSubmatch(input, -1)
	serialMatch := serialRe.FindAllStringSubmatch(input, -1)
	activeWorkspaceMatch := activeWorkspaceRe.FindAllStringSubmatch(input, -1)
	speicalWorkspaceMatch := specialWorkspaceRe.FindAllStringSubmatch(input, -1)

	for i := range monitorMatch {
		var result MonitorResult

		result.Name = getSubmatch(monitorMatch[i], 1)
		result.Id = getSubmatch(monitorMatch[i], 2)
		result.Resolution.Width = getSubmatch(resolutionMatch[i], 1)
		result.Resolution.Height = getSubmatch(resolutionMatch[i], 2)
		result.Resolution.RefreshRate = getSubmatch(resolutionMatch[i], 3)
		result.Position.X = getSubmatch(resolutionMatch[i], 4)
		result.Position.Y = getSubmatch(resolutionMatch[i], 5)
		result.Description = getSubmatch(descriptionMatch[i], 1)
		result.Make = getSubmatch(makeMatch[i], 1)
		result.Model = getSubmatch(modelMatch[i], 1)
		result.Size.Height = getSubmatch(physicalSizeMatch[i], 1)
		result.Size.Width = getSubmatch(physicalSizeMatch[i], 2)
		result.Serial = getSubmatch(serialMatch[i], 1)
		result.ActiveWorkspace = getSubmatch(activeWorkspaceMatch[i], 1)
		result.SpecialWorkspace = getSubmatch(speicalWorkspaceMatch[i], 1)

		results = append(results, &result)
	}

	return results, err
}

// /////////////////////////////////////////////////////////////////////////////

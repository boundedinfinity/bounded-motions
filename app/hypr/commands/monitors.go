package commands

import (
	"regexp"
	"strings"
)

// /////////////////////////////////////////////////////////////////////////////
// https://wiki.hypr.land/Configuring/Using-hyprctl/#info

var _ HyprCommand = &Monitors{}

type Monitors struct{}

func (this Monitors) String() string {
	return "monitors"
}

func (_ Monitors) hyprCommand() HyprCommandType {
	return HyprCommandTypes.Monitors
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
	X int
	Y int
}

type MonitorSize struct {
	Width  int
	Height int
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
		result.Position.X = getSubmatchInt(resolutionMatch[i], 4)
		result.Position.Y = getSubmatchInt(resolutionMatch[i], 5)
		result.Description = getSubmatch(descriptionMatch[i], 1)
		result.Make = getSubmatch(makeMatch[i], 1)
		result.Model = getSubmatch(modelMatch[i], 1)
		result.Size.Height = getSubmatchInt(physicalSizeMatch[i], 1)
		result.Size.Width = getSubmatchInt(physicalSizeMatch[i], 2)
		result.Serial = getSubmatch(serialMatch[i], 1)
		result.ActiveWorkspace = getSubmatch(activeWorkspaceMatch[i], 1)
		result.SpecialWorkspace = getSubmatch(speicalWorkspaceMatch[i], 1)

		results = append(results, &result)
	}

	return results, err
}

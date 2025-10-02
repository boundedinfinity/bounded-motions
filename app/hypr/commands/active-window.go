package commands

import (
	"fmt"
	"regexp"
)

// /////////////////////////////////////////////////////////////////////////////
// https://wiki.hypr.land/Configuring/Using-hyprctl/#info

var _ HyprCommand = &ActiveWindow{}

type ActiveWindow struct{}

func (this ActiveWindow) String() string {
	return "activewindow"
}

func (_ ActiveWindow) hyprCommand() HyprCommandType {
	return HyprCommandTypes.ActiveWindow
}

// /////////////////////////////////////////////////////////////////////////////

var _ HyperCommandResult = &ActiveWindowResult{}

type ActiveWindowResult struct {
	Id               string
	Title            string
	Mapped           bool
	Hidden           bool
	At               MonitorPosistion
	Size             MonitorSize
	WorkspaceId      string
	Floating         bool
	Pseudo           bool
	MonitorId        string
	Class            string
	InitialClass     string
	InitialTitle     string
	Pid              int
	Xwayland         int
	Pinned           bool
	Fullscreen       bool
	FullscreenClient int
	Grouped          string
	Tags             string
	Swallowing       bool
	FocusHistoryId   int
	InhibitingIdle   int
	XdgTag           string
	XdgDescription   string
}

func (this ActiveWindowResult) String() string {
	return fmt.Sprintf("active window [%s]", this.Id)
}

func (_ ActiveWindowResult) hyperCommandResult() {}

var (
	windowLineRe             = regexp.MustCompile(`Window (\w+) ->`)
	windowMappedRe           = regexp.MustCompile(`mapped: (.*)\n`)
	windowHiddenRe           = regexp.MustCompile(`hidden: (.*)\n`)
	windowAtRe               = regexp.MustCompile(`at: (\d+),(\d+)\n`)
	windowSizeRe             = regexp.MustCompile(`size: (\d+),(\d+)\n`)
	windowWorkspaceRe        = regexp.MustCompile(`workspace: (.*)\(.*\n`)
	windowFloatingRe         = regexp.MustCompile(`floating: (\d+)\n`)
	windowPseudoRe           = regexp.MustCompile(`pseudo: (\d+)\n`)
	windowMonitorRe          = regexp.MustCompile(`monitor: (\d+)\n`)
	windowClassRe            = regexp.MustCompile(`class: (.+)\n`)
	windowTitleRe            = regexp.MustCompile(`title: (.+)\n`)
	windowInitialClassRe     = regexp.MustCompile(`initialClass: (.+)\n`)
	windowInitialTitleRe     = regexp.MustCompile(`initialTitle: (.+)\n`)
	windowPidRe              = regexp.MustCompile(`pid: (\d+)\n`)
	windowXwaylandRe         = regexp.MustCompile(`xwayland: (\d+)\n`)
	windowPinnedRe           = regexp.MustCompile(`pinned: (\d+)\n`)
	windowFullscreenRe       = regexp.MustCompile(`fullscreen: (\d+)\n`)
	windowFullscreenClientRe = regexp.MustCompile(`fullscreenClient: (\d+)\n`)
	windowGroupedRe          = regexp.MustCompile(`grouped: (\d+)\n`)
	windowTagsRe             = regexp.MustCompile(`tags: (.*)\n`)
	windowSwallowingRe       = regexp.MustCompile(`swallowing: (\d+)\n`)
	windowFocusHistoryIdRe   = regexp.MustCompile(`focusHistoryID: (\d+)\n`)
	windowInhibitingIdleRe   = regexp.MustCompile(`inhibitingIdle: (\d+)\n`)
	windowXdgTagRe           = regexp.MustCompile(`xdgTag: (.*)\n`)
	windowXdgDescriptionRe   = regexp.MustCompile(`xdgDescription: (.*)\n`)
)

func parseActiveWindow(input string, i int) (*ActiveWindowResult, error) {
	var result ActiveWindowResult
	var err error

	windowLineMatch := windowLineRe.FindAllStringSubmatch(input, -1)
	windowMappedMatch := windowMappedRe.FindAllStringSubmatch(input, -1)
	windowHiddenMatch := windowHiddenRe.FindAllStringSubmatch(input, -1)
	windowAtMatch := windowAtRe.FindAllStringSubmatch(input, -1)
	windowSizeMatch := windowSizeRe.FindAllStringSubmatch(input, -1)
	windowWorkspaceMatch := windowWorkspaceRe.FindAllStringSubmatch(input, -1)
	windowFloatingMatch := windowFloatingRe.FindAllStringSubmatch(input, -1)
	windowPseudoMatch := windowPseudoRe.FindAllStringSubmatch(input, -1)
	windowMonitorMatch := windowMonitorRe.FindAllStringSubmatch(input, -1)
	windowClassMatch := windowClassRe.FindAllStringSubmatch(input, -1)
	windowTitleMatch := windowTitleRe.FindAllStringSubmatch(input, -1)
	windowInitialClassMatch := windowInitialClassRe.FindAllStringSubmatch(input, -1)
	windowInitialTitleMatch := windowInitialTitleRe.FindAllStringSubmatch(input, -1)
	windowPidMatch := windowPidRe.FindAllStringSubmatch(input, -1)
	windowXwaylandMatch := windowXwaylandRe.FindAllStringSubmatch(input, -1)
	windowPineedMatch := windowPinnedRe.FindAllStringSubmatch(input, -1)
	windowFullscreenMatch := windowFullscreenRe.FindAllStringSubmatch(input, -1)
	windowFullscreenClientMatch := windowFullscreenClientRe.FindAllStringSubmatch(input, -1)
	windowGroupMatch := windowGroupedRe.FindAllStringSubmatch(input, -1)
	windowTagsMatch := windowTagsRe.FindAllStringSubmatch(input, -1)
	windowSwallowingMatch := windowSwallowingRe.FindAllStringSubmatch(input, -1)
	windowFocusHistoryIdMatch := windowFocusHistoryIdRe.FindAllStringSubmatch(input, -1)
	windowInhibitingIdleMatch := windowInhibitingIdleRe.FindAllStringSubmatch(input, -1)
	windowXdgTagMatch := windowXdgTagRe.FindAllStringSubmatch(input, -1)
	windowXdgDescriptionmatch := windowXdgDescriptionRe.FindAllStringSubmatch(input, -1)

	result.Id = getSubmatch(windowLineMatch[i], 1)
	result.Title = getSubmatch(windowTitleMatch[i], 1)
	result.Mapped = getSubmatchBool(windowMappedMatch[i], 1)
	result.Hidden = getSubmatchBool(windowHiddenMatch[i], 1)
	result.At.X = getSubmatchInt(windowAtMatch[i], 1)
	result.At.Y = getSubmatchInt(windowAtMatch[i], 2)
	result.Size.Height = getSubmatchInt(windowSizeMatch[i], 1)
	result.Size.Width = getSubmatchInt(windowSizeMatch[i], 2)
	result.WorkspaceId = getSubmatch(windowWorkspaceMatch[i], 1)
	result.Floating = getSubmatchBool(windowFloatingMatch[i], 1)
	result.Pseudo = getSubmatchBool(windowPseudoMatch[i], 1)
	result.MonitorId = getSubmatch(windowMonitorMatch[i], 1)
	result.Class = getSubmatch(windowClassMatch[i], 1)
	result.InitialClass = getSubmatch(windowInitialClassMatch[i], 1)
	result.InitialTitle = getSubmatch(windowInitialTitleMatch[i], 1)
	result.Pid = getSubmatchInt(windowPidMatch[i], 1)
	result.Xwayland = getSubmatchInt(windowXwaylandMatch[i], 1)
	result.Pinned = getSubmatchBool(windowPineedMatch[i], 1)
	result.Fullscreen = getSubmatchBool(windowFullscreenMatch[i], 1)
	result.FullscreenClient = getSubmatchInt(windowFullscreenClientMatch[i], 1)
	result.Grouped = getSubmatch(windowGroupMatch[i], 1)
	result.Tags = getSubmatch(windowTagsMatch[i], 1)
	result.Swallowing = getSubmatchBool(windowSwallowingMatch[i], 1)
	result.FocusHistoryId = getSubmatchInt(windowFocusHistoryIdMatch[i], 1)
	result.InhibitingIdle = getSubmatchInt(windowInhibitingIdleMatch[i], 1)
	result.XdgTag = getSubmatch(windowXdgTagMatch[i], 1)
	result.XdgDescription = getSubmatch(windowXdgDescriptionmatch[i], 1)

	return &result, err
}

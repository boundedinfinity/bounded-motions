// Package commands implements functionality to send and receive commands to Hyprland
package commands

import (
	"errors"
	"fmt"
	"go-motions/hypr/events"
	"go-motions/hypr/model"

	"strconv"
	"strings"
	"time"
)

type HyprCommand interface {
	String() string
	hyprCommand() HyprCommandType // discriminator
}

type HyperCommandResult interface {
	String() string
	hyperCommandResult() // discriminator
}

func ParseResult(cmd HyprCommand, input string) (HyperCommandResult, error) {
	var result HyperCommandResult
	var err error

	switch cmd.String() {
	case string(HyprCommandTypes.Version):
		result, err = parseVersion(input)
	}

	return result, err
}

// /////////////////////////////////////////////////////////////////////////////
// https://wiki.hypr.land/Configuring/Using-hyprctl/
// /////////////////////////////////////////////////////////////////////////////

type HyprCommandType string

var HyprCommandTypes = hyprCommandTypes{
	Error:                          "error",
	Unhandled:                      "unhandled",
	Version:                        "version",
	Monitors:                       "monitors",
	Workspaces:                     "workspaces",
	ActiveWorkspace:                "activieworkspace",
	WorkspaceRules:                 "workspacerules",
	Clients:                        "clients",
	Devices:                        "devices",
	Decorations:                    "decorations",
	Binds:                          "binds",
	ActiveWindow:                   "activewindow",
	Layers:                         "layers",
	Splash:                         "splash",
	GetOption:                      "getoption",
	CursorPos:                      "cursorpos",
	Animations:                     "animations",
	Instances:                      "instances",
	Layouts:                        "layouts",
	ConfigErrors:                   "configerrors",
	RollingLog:                     "rollinglog",
	Locked:                         "locked",
	Descriptions:                   "descriptions",
	Submap:                         "submap",
	Keyword:                        "keyword",
	Output:                         "output",
	SwitchXkbLayout:                "switchxkblayout",
	SetError:                       "seterror",
	GetProp:                        "getprop",
	Notify:                         "notify",
	DismissNotify:                  "dismissnotify",
	Kill:                           "kill",
	Exec:                           "dispatch exec",
	Execr:                          "dispatch execr",
	Pass:                           "dispatch pass",
	SendShortcut:                   "dispatch sendshortcut",
	SendKeyState:                   "dispatch sendkeystate",
	KillActive:                     "dispatch killactive",
	ForceKillActive:                "dispatch forcekillactive",
	CloseWindow:                    "dispatch closewindow",
	KillWindow:                     "dispatch killwindow",
	Signal:                         "dispatch signal",
	SignalWindow:                   "dispatch signalwindow",
	Workspace:                      "dispatch workspace",
	MoveToWorkspace:                "dispatch movetoworkspace",
	MoveToWorkspaceSilent:          "dispatch movetoworkspacesilent",
	ToggleFloating:                 "dispatch togglefloating",
	SetFloating:                    "dispatch setfloating",
	SetTiled:                       "dispatch settiled",
	Fullscreen:                     "dispatch fullscreen",
	FullscreenState:                "dispatch fullscreenstate",
	Dpms:                           "dispatch dpms",
	Pin:                            "dispatch pin",
	MoveFocus:                      "dispatch movefocus",
	MoveWindow:                     "dispatch movewindow",
	SwapWindow:                     "dispatch swapwindow",
	CenterWindow:                   "dispatch centerwindow",
	ResizeActive:                   "dispatch resizeactive",
	MoveActive:                     "dispatch moveactive",
	ResizeWindowPixel:              "dispatch resizewindowpixel",
	MoveWindowPixel:                "dispatch movewindowpixel",
	CycleNext:                      "dispatch cyclenext",
	SwapNext:                       "dispatch swapnext",
	TagWindow:                      "dispatch tagwindow",
	FocusWindow:                    "dispatch focuswindow",
	FocusMonitor:                   "dispatch focusmonitor",
	SplitRatio:                     "dispatch splitratio",
	MoveCursorToCorner:             "dispatch movecursortocorner",
	MoveCursor:                     "dispatch movecursor",
	RenameWorkspace:                "dispatch renameworkspace",
	Exit:                           "dispatch exit",
	ForceRendererReload:            "dispatch forcerendererreload",
	MoveCurrentWorkspaceToMonitor:  "dispatch movecurrentworkspacetomonitor",
	FocusWorkspaceOnCurrentMonitor: "dispatch focusworkspaceoncurrentmonitor",
	MoveWorkspaceToMonitor:         "dispatch moveworkspacetomonitor",
	SwapActiveWorkspaces:           "dispatch swapactiveworkspaces",
	BringActiveToTop:               "dispatch bringactivetotop",
	AlterZorder:                    "dispatch alterzorder",
	ToggleSpecialWorkspace:         "dispatch togglespecialworkspace",
	FocusUrgentOrLast:              "dispatch focusurgentorlast",
	ToggleGroup:                    "dispatch togglegroup",
	ChangeGroupActive:              "dispatch changegroupactive",
	FocusCurrentOrLast:             "dispatch focuscurrentorlast",
	LockGroups:                     "dispatch lockgroups",
	LockActiveGroup:                "dispatch lockactivegroup",
	MoveIntoGroup:                  "dispatch moveintogroup",
	MoveOutOfGroup:                 "dispatch moveoutofgroup",
	MoveWindowOrGroup:              "dispatch movewindoworgroup",
	MoveGroupWindow:                "dispatch movegroupwindow",
	DenyWindowFromGroup:            "dispatch denywindowfromgroup",
	SetIgnoreGroupLock:             "dispatch setignoregrouplock",
	Global:                         "dispatch global",
	submap:                         "dispatch submap",
	Event:                          "dispatch event",
	SetProp:                        "dispatch setprop",
	ToggleSwallow:                  "dispatch toggleswallow",
}

type hyprCommandTypes struct {
	Error                          HyprCommandType
	Unhandled                      HyprCommandType
	Version                        HyprCommandType
	Monitors                       HyprCommandType
	Workspaces                     HyprCommandType
	ActiveWorkspace                HyprCommandType
	WorkspaceRules                 HyprCommandType
	Clients                        HyprCommandType
	Devices                        HyprCommandType
	Decorations                    HyprCommandType
	Binds                          HyprCommandType
	ActiveWindow                   HyprCommandType
	Layers                         HyprCommandType
	Splash                         HyprCommandType
	GetOption                      HyprCommandType
	CursorPos                      HyprCommandType
	Animations                     HyprCommandType
	Instances                      HyprCommandType
	Layouts                        HyprCommandType
	ConfigErrors                   HyprCommandType
	RollingLog                     HyprCommandType
	Locked                         HyprCommandType
	Descriptions                   HyprCommandType
	Submap                         HyprCommandType
	Keyword                        HyprCommandType
	reload                         HyprCommandType
	Kill                           HyprCommandType
	setcursor                      HyprCommandType
	Output                         HyprCommandType
	SwitchXkbLayout                HyprCommandType
	SetError                       HyprCommandType
	GetProp                        HyprCommandType
	Notify                         HyprCommandType
	DismissNotify                  HyprCommandType
	Exec                           HyprCommandType
	Execr                          HyprCommandType
	Pass                           HyprCommandType
	SendShortcut                   HyprCommandType
	SendKeyState                   HyprCommandType
	KillActive                     HyprCommandType
	ForceKillActive                HyprCommandType
	CloseWindow                    HyprCommandType
	KillWindow                     HyprCommandType
	Signal                         HyprCommandType
	SignalWindow                   HyprCommandType
	Workspace                      HyprCommandType
	MoveToWorkspace                HyprCommandType
	MoveToWorkspaceSilent          HyprCommandType
	ToggleFloating                 HyprCommandType
	SetFloating                    HyprCommandType
	SetTiled                       HyprCommandType
	Fullscreen                     HyprCommandType
	FullscreenState                HyprCommandType
	Dpms                           HyprCommandType
	Pin                            HyprCommandType
	MoveFocus                      HyprCommandType
	MoveWindow                     HyprCommandType
	SwapWindow                     HyprCommandType
	CenterWindow                   HyprCommandType
	ResizeActive                   HyprCommandType
	MoveActive                     HyprCommandType
	ResizeWindowPixel              HyprCommandType
	MoveWindowPixel                HyprCommandType
	CycleNext                      HyprCommandType
	SwapNext                       HyprCommandType
	TagWindow                      HyprCommandType
	FocusWindow                    HyprCommandType
	FocusMonitor                   HyprCommandType
	SplitRatio                     HyprCommandType
	MoveCursorToCorner             HyprCommandType
	MoveCursor                     HyprCommandType
	RenameWorkspace                HyprCommandType
	Exit                           HyprCommandType
	ForceRendererReload            HyprCommandType
	MoveCurrentWorkspaceToMonitor  HyprCommandType
	FocusWorkspaceOnCurrentMonitor HyprCommandType
	MoveWorkspaceToMonitor         HyprCommandType
	SwapActiveWorkspaces           HyprCommandType
	BringActiveToTop               HyprCommandType
	AlterZorder                    HyprCommandType
	ToggleSpecialWorkspace         HyprCommandType
	FocusUrgentOrLast              HyprCommandType
	ToggleGroup                    HyprCommandType
	ChangeGroupActive              HyprCommandType
	FocusCurrentOrLast             HyprCommandType
	LockGroups                     HyprCommandType
	LockActiveGroup                HyprCommandType
	MoveIntoGroup                  HyprCommandType
	MoveOutOfGroup                 HyprCommandType
	MoveWindowOrGroup              HyprCommandType
	MoveGroupWindow                HyprCommandType
	DenyWindowFromGroup            HyprCommandType
	SetIgnoreGroupLock             HyprCommandType
	Global                         HyprCommandType
	submap                         HyprCommandType
	Event                          HyprCommandType
	SetProp                        HyprCommandType
	ToggleSwallow                  HyprCommandType
}

// /////////////////////////////////////////////////////////////////////////////
// https://wiki.hypr.land/Configuring/Using-hyprctl/#reload

var _ HyprCommand = &Reload{}

type Reload struct {
	Command string
}

func (this Reload) String() string {
	return "reload"
}

func (_ Reload) hyprCommand() HyprCommandType {
	return HyprCommandTypes.Version
}

// /////////////////////////////////////////////////////////////////////////////
// https://wiki.hypr.land/Configuring/Using-hyprctl/#kill

var _ HyprCommand = &Kill{}

type Kill struct {
	Command string
}

func (this Kill) String() string {
	return "kill"
}

func (_ Kill) hyprCommand() HyprCommandType {
	return HyprCommandTypes.Kill
}

// /////////////////////////////////////////////////////////////////////////////
// https://wiki.hypr.land/Configuring/Using-hyprctl/#setcursor

var _ HyprCommand = &SetCursor{}

type SetCursor struct {
	Fontname string
	Size     int
}

func (this SetCursor) String() string {
	return fmt.Sprintf("setcursor %s %d", this.Fontname, this.Size)
}

func (_ SetCursor) hyprCommand() HyprCommandType {
	return HyprCommandTypes.SetError
}

// /////////////////////////////////////////////////////////////////////////////
// https://wiki.hypr.land/Configuring/Using-hyprctl/#switchxkblayout

var _ HyprCommand = &SwitchXkbLayout{}

type SwitchXkbLayout struct {
	Device  string
	Command string
}

func (this SwitchXkbLayout) String() string {
	return fmt.Sprintf("switchxkblayout %s %s", this.Device, this.Command)
}

func (_ SwitchXkbLayout) hyprCommand() HyprCommandType {
	return HyprCommandTypes.SwitchXkbLayout
}

// /////////////////////////////////////////////////////////////////////////////
// https://wiki.hypr.land/Configuring/Using-hyprctl/#seterror

var _ HyprCommand = &SetError{}

type SetError struct {
	Err     error
	Disable bool
}

func (this SetError) String() string {
	if this.Disable {
		return "seterror disabled"
	}

	return fmt.Sprintf("seterror %s", this.Err)
}

func (_ SetError) hyprCommand() HyprCommandType {
	return HyprCommandTypes.SetError
}

// /////////////////////////////////////////////////////////////////////////////
// https://wiki.hypr.land/Configuring/Using-hyprctl/#notify

var _ HyprCommand = &Notify{}

type Notify struct {
	Icon     model.HyprNotifyIcon
	Duration time.Duration
	Color    string
	Message  string
}

func (this Notify) String() string {
	return fmt.Sprintf("notify %s %d %s %s", this.Icon, this.Duration.Milliseconds(), this.Color, this.Message)
}

func (_ Notify) hyprCommand() HyprCommandType {
	return HyprCommandTypes.Notify
}

// /////////////////////////////////////////////////////////////////////////////
// https://wiki.hypr.land/Configuring/Using-hyprctl/#notify

var _ HyprCommand = &DismissNotify{}

type DismissNotify struct {
	Count int
}

func (this DismissNotify) String() string {
	return fmt.Sprintf("dismissnotify %d", this.Count)
}

func (_ DismissNotify) hyprCommand() HyprCommandType {
	return HyprCommandTypes.DismissNotify
}

// /////////////////////////////////////////////////////////////////////////////
// Info
// /////////////////////////////////////////////////////////////////////////////
// https://wiki.hypr.land/Configuring/Using-hyprctl/#info

var _ HyprCommand = &Version{}

type Version struct{}

func (this Version) String() string {
	return "version"
}

func (_ Version) hyprCommand() HyprCommandType {
	return HyprCommandTypes.Version
}

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

func parseVersion(input string) (*VersionResult, error) {
	info := VersionResult{Dependencies: make(map[string]string)}
	lines := strings.Split(input, "\n")

	for _, line := range lines {
		line = strings.TrimSpace(line)

		if line == "" {
			continue
		}

		if strings.HasPrefix("Hyprland", line) {
			parts := strings.Split(line, "built from")

			info.Version = parts[0]
			info.Version = strings.ReplaceAll(info.Version, "Hyperland", "")
			info.Version = strings.TrimSpace(info.Version)

			parts = strings.Split(parts[1], "(version:")
			info.Commit = parts[0]
			info.Commit = strings.ReplaceAll(info.Commit, "commit", "")
			info.Commit = strings.TrimSpace(info.Commit)

			continue
		}

		if strings.HasPrefix(line, "Date:") {
			info.Date = line
			info.Date = strings.ReplaceAll(info.Date, "Date:", "")
			info.Date = strings.TrimSpace(info.Date)
			continue
		}

		if strings.HasPrefix(line, "Tag:") {
			info.Tag = strings.ReplaceAll(line, "Tag:", "")
			info.Tag = strings.TrimSpace(info.Tag)
			continue
		}
	}

	return &info, nil
}

// /////////////////////////////////////////////////////////////////////////////

var _ HyprCommand = &Monitors{}

type Monitors struct{}

func (this Monitors) String() string {
	return "monitors"
}

func (_ Monitors) hyprCommand() HyprCommandType {
	return HyprCommandTypes.Monitors
}

// /////////////////////////////////////////////////////////////////////////////

var _ HyprCommand = &Monitors{}

type Workspaces struct{}

func (this Workspaces) String() string {
	return "workspaces"
}

func (_ Workspaces) hyprCommand() HyprCommandType {
	return HyprCommandTypes.Workspaces
}

// /////////////////////////////////////////////////////////////////////////////

var _ HyprCommand = &ActiveWorkspace{}

type ActiveWorkspace struct{}

func (this ActiveWorkspace) String() string {
	return "activeworkspace"
}

func (_ ActiveWorkspace) hyprCommand() HyprCommandType {
	return HyprCommandTypes.ActiveWorkspace
}

// /////////////////////////////////////////////////////////////////////////////

var _ HyprCommand = &WorkspaceRules{}

type WorkspaceRules struct{}

func (this WorkspaceRules) String() string {
	return "workspacerules"
}

func (_ WorkspaceRules) hyprCommand() HyprCommandType {
	return HyprCommandTypes.WorkspaceRules
}

// /////////////////////////////////////////////////////////////////////////////

var _ HyprCommand = &Clients{}

type Clients struct{}

func (this Clients) String() string {
	return "clients"
}

func (_ Clients) hyprCommand() HyprCommandType {
	return HyprCommandTypes.Clients
}

// /////////////////////////////////////////////////////////////////////////////

var _ HyprCommand = &Devices{}

type Devices struct{}

func (this Devices) String() string {
	return "devices"
}

func (_ Devices) hyprCommand() HyprCommandType {
	return HyprCommandTypes.Devices
}

// /////////////////////////////////////////////////////////////////////////////

var _ HyprCommand = &Decorations{}

type Decorations struct {
	Window model.HyprWindowId
}

func (this Decorations) String() string {
	return fmt.Sprintf("decorations %s", this.Window)
}

func (_ Decorations) hyprCommand() HyprCommandType {
	return HyprCommandTypes.Decorations
}

// /////////////////////////////////////////////////////////////////////////////

var _ HyprCommand = &Binds{}

type Binds struct{}

func (this Binds) String() string {
	return "binds"
}

func (_ Binds) hyprCommand() HyprCommandType {
	return HyprCommandTypes.Binds
}

// /////////////////////////////////////////////////////////////////////////////

var _ HyprCommand = &ActiveWindow{}

type ActiveWindow struct{}

func (this ActiveWindow) String() string {
	return "activewindow"
}

func (_ ActiveWindow) hyprCommand() HyprCommandType {
	return HyprCommandTypes.ActiveWindow
}

// /////////////////////////////////////////////////////////////////////////////

var _ HyprCommand = &Layers{}

type Layers struct{}

func (this Layers) String() string {
	return "layers"
}

func (_ Layers) hyprCommand() HyprCommandType {
	return HyprCommandTypes.Layers
}

// /////////////////////////////////////////////////////////////////////////////

var _ HyprCommand = &Splash{}

type Splash struct{}

func (this Splash) String() string {
	return "splash"
}

func (_ Splash) hyprCommand() HyprCommandType {
	return HyprCommandTypes.Splash
}

// /////////////////////////////////////////////////////////////////////////////

var _ HyprCommand = &CursorPos{}

type CursorPos struct{}

func (this CursorPos) String() string {
	return "cursorpos"
}

func (_ CursorPos) hyprCommand() HyprCommandType {
	return HyprCommandTypes.CursorPos
}

// /////////////////////////////////////////////////////////////////////////////

var _ HyprCommand = &Animations{}

type Animations struct{}

func (this Animations) String() string {
	return "animations"
}

func (_ Animations) hyprCommand() HyprCommandType {
	return HyprCommandTypes.Animations
}

// /////////////////////////////////////////////////////////////////////////////

var _ HyprCommand = &Instances{}

type Instances struct{}

func (this Instances) String() string {
	return "instances"
}

func (_ Instances) hyprCommand() HyprCommandType {
	return HyprCommandTypes.Instances
}

// /////////////////////////////////////////////////////////////////////////////

var _ HyprCommand = &Layouts{}

type Layouts struct{}

func (this Layouts) String() string {
	return "layouts"
}

func (_ Layouts) hyprCommand() HyprCommandType {
	return HyprCommandTypes.Layouts
}

// /////////////////////////////////////////////////////////////////////////////

var _ HyprCommand = &ConfigErrors{}

type ConfigErrors struct{}

func (this ConfigErrors) String() string {
	return "configerrors"
}

func (_ ConfigErrors) hyprCommand() HyprCommandType {
	return HyprCommandTypes.ConfigErrors
}

// /////////////////////////////////////////////////////////////////////////////

var _ HyprCommand = &RollingLog{}

type RollingLog struct{}

func (this RollingLog) String() string {
	return "rollinglog"
}

func (_ RollingLog) hyprCommand() HyprCommandType {
	return HyprCommandTypes.RollingLog
}

// /////////////////////////////////////////////////////////////////////////////

var _ HyprCommand = &Locked{}

type Locked struct{}

func (this Locked) String() string {
	return "locked"
}

func (_ Locked) hyprCommand() HyprCommandType {
	return HyprCommandTypes.Locked
}

// /////////////////////////////////////////////////////////////////////////////

var _ HyprCommand = &Descriptions{}

type Descriptions struct{}

func (this Descriptions) String() string {
	return "descriptions"
}

func (_ Descriptions) hyprCommand() HyprCommandType {
	return HyprCommandTypes.Descriptions
}

// /////////////////////////////////////////////////////////////////////////////

var _ HyprCommand = &Submap{}

type Submap struct{}

func (this Submap) String() string {
	return "submap"
}

func (_ Submap) hyprCommand() HyprCommandType {
	return HyprCommandTypes.Submap
}

// func (this *infoCommandBuilder) GetOption() HyprCommand {
// 	return &hyprCommand{line: "getoption"}
// }

// /////////////////////////////////////////////////////////////////////////////
// Dispatchers
//
// https://wiki.hypr.land/Configuring/Using-hyprctl/#dispatch
// https://wiki.hypr.land/Configuring/Dispatchers/
//
// /////////////////////////////////////////////////////////////////////////////

var _ HyprCommand = &Exec{}

type Exec struct {
	Command string
}

func (this Exec) String() string {
	return fmt.Sprintf("dispatch command %s", this.Command)
}

func (_ Exec) hyprCommand() HyprCommandType {
	return HyprCommandTypes.Exec
}

// /////////////////////////////////////////////////////////////////////////////

var _ HyprCommand = &Execr{}

type Execr struct {
	Command string
}

func (this Execr) String() string {
	return fmt.Sprintf("dispatch command %s", this.Command)
}

func (_ Execr) hyprCommand() HyprCommandType {
	return HyprCommandTypes.Execr
}

// /////////////////////////////////////////////////////////////////////////////

var _ HyprCommand = &Pass{}

type Pass struct {
	Window model.HyprWindowId
}

func (this Pass) String() string {
	return fmt.Sprintf("dispatch pass %s", this.Window)
}

func (_ Pass) hyprCommand() HyprCommandType {
	return HyprCommandTypes.Pass
}

// /////////////////////////////////////////////////////////////////////////////

var _ HyprCommand = &SendShortcut{}

type SendShortcut struct {
	Window model.HyprWindowId
}

func (this SendShortcut) String() string {
	return fmt.Sprintf("dispatch sendshortcut %s", this.Window)
}

func (_ SendShortcut) hyprCommand() HyprCommandType {
	return HyprCommandTypes.SendShortcut
}

// /////////////////////////////////////////////////////////////////////////////

var _ HyprCommand = &SendKeyState{}

type SendKeyState struct {
	ModKey model.ModKey
	Key    string
	State  model.KeyState
	Window model.HyprWindowId
}

func (this SendKeyState) String() string {
	return fmt.Sprintf("dispatch sendkeystate %s, %s, %s, %s", this.ModKey, this.Key, this.State, this.Window)
}

func (_ SendKeyState) hyprCommand() HyprCommandType {
	return HyprCommandTypes.SendKeyState
}

// /////////////////////////////////////////////////////////////////////////////

var _ HyprCommand = &KillActive{}

type KillActive struct {
}

func (this KillActive) String() string {
	return "dispatch killactive"
}

func (_ KillActive) hyprCommand() HyprCommandType {
	return HyprCommandTypes.KillActive
}

// /////////////////////////////////////////////////////////////////////////////

var _ HyprCommand = &ForceKillActive{}

type ForceKillActive struct {
}

func (this ForceKillActive) String() string {
	return "dispatch killactive"
}

func (_ ForceKillActive) hyprCommand() HyprCommandType {
	return HyprCommandTypes.ForceKillActive
}

// /////////////////////////////////////////////////////////////////////////////

var _ HyprCommand = &CloseWindow{}

type CloseWindow struct {
	Window model.HyprWindowId
}

func (this CloseWindow) String() string {
	return fmt.Sprintf("dispatch closewindow %s", this.Window)
}

func (_ CloseWindow) hyprCommand() HyprCommandType {
	return HyprCommandTypes.CloseWindow
}

// /////////////////////////////////////////////////////////////////////////////

var _ HyprCommand = &KillWindow{}

type KillWindow struct {
	Window model.HyprWindowId
}

func (this KillWindow) String() string {
	return fmt.Sprintf("dispatch killwindow %s", this.Window)
}

func (_ KillWindow) hyprCommand() HyprCommandType {
	return HyprCommandTypes.KillWindow
}

// /////////////////////////////////////////////////////////////////////////////

var _ HyprCommand = &Signal{}

type Signal struct{}

func (this Signal) String() string {
	return "dispatch signal"
}

func (_ Signal) hyprCommand() HyprCommandType {
	return HyprCommandTypes.Signal
}

// /////////////////////////////////////////////////////////////////////////////

var _ HyprCommand = &SignalWindow{}

type SignalWindow struct {
	Window model.HyprWindowId
}

func (this SignalWindow) String() string {
	return fmt.Sprintf("dispatch signalwindow %s", this.Window)
}

func (_ SignalWindow) hyprCommand() HyprCommandType {
	return HyprCommandTypes.SignalWindow
}

// /////////////////////////////////////////////////////////////////////////////

var _ HyprCommand = &Workspace{}

type Workspace struct {
	Workspace model.HyprWorkspaceId
}

func (this Workspace) String() string {
	return fmt.Sprintf("dispatch workspace %s", this.Workspace)
}

func (_ Workspace) hyprCommand() HyprCommandType {
	return HyprCommandTypes.Workspace
}

// /////////////////////////////////////////////////////////////////////////////

var _ HyprCommand = &MoveToWorkspace{}

type MoveToWorkspace struct {
	Workspace model.HyprWorkspaceId
	Window    model.HyprWindowId
}

func (this MoveToWorkspace) String() string {
	if this.Window != nil {
		return fmt.Sprintf("dispatch movetoworkspace %s,%s", this.Workspace, this.Workspace)
	}

	return fmt.Sprintf("dispatch movetoworkspace %s", this.Workspace)
}

func (_ MoveToWorkspace) hyprCommand() HyprCommandType {
	return HyprCommandTypes.MoveToWorkspace
}

// /////////////////////////////////////////////////////////////////////////////

var _ HyprCommand = &MoveToWorkspaceSilent{}

type MoveToWorkspaceSilent struct {
	Workspace model.HyprWorkspaceId
	Window    model.HyprWindowId
}

func (this MoveToWorkspaceSilent) String() string {
	if this.Window != nil {
		return fmt.Sprintf("dispatch movetoworkspacesilent %s,%s", this.Workspace, this.Workspace)
	}

	return fmt.Sprintf("dispatch movetoworkspacesilent %s", this.Workspace)
}

func (_ MoveToWorkspaceSilent) hyprCommand() HyprCommandType {
	return HyprCommandTypes.MoveToWorkspaceSilent
}

// /////////////////////////////////////////////////////////////////////////////

var _ HyprCommand = &ToggleFloating{}

type ToggleFloating struct {
	Window model.HyprWindowId
}

func (this ToggleFloating) String() string {
	if this.Window != nil {
		return fmt.Sprintf("dispatch togglefloating %s", this.Window)
	}

	return "dispatch togglefloating"
}

func (_ ToggleFloating) hyprCommand() HyprCommandType {
	return HyprCommandTypes.ToggleFloating
}

// /////////////////////////////////////////////////////////////////////////////

var _ HyprCommand = &SetFloating{}

type SetFloating struct {
	Window model.HyprWindowId
}

func (this SetFloating) String() string {
	if this.Window != nil {
		return fmt.Sprintf("dispatch setfloating %s", this.Window)
	}

	return "dispatch setfloating"
}

func (_ SetFloating) hyprCommand() HyprCommandType {
	return HyprCommandTypes.SetFloating
}

// /////////////////////////////////////////////////////////////////////////////

var _ HyprCommand = &SetTiled{}

type SetTiled struct {
	Window model.HyprWindowId
}

func (this SetTiled) String() string {
	if this.Window != nil {
		return fmt.Sprintf("dispatch settiled %s", this.Window)
	}

	return "dispatch settiled"
}

func (_ SetTiled) hyprCommand() HyprCommandType {
	return HyprCommandTypes.SetTiled
}

// /////////////////////////////////////////////////////////////////////////////

var _ HyprCommand = &FullScreen{}

type FullScreen struct {
	Mode FullScreenMode
}

func (this FullScreen) String() string {
	return fmt.Sprintf("dispatch fullscreen %d", this.Mode)
}

func (_ FullScreen) hyprCommand() HyprCommandType {
	return HyprCommandTypes.Fullscreen
}

type FullScreenMode int

func (this FullScreenMode) String() string {
	return strconv.Itoa(int(this))
}

var FullScreenModes = fullscreenModes{
	FullScreen: 0,
	Maximize:   1,
}

type fullscreenModes struct {
	FullScreen FullScreenMode
	Maximize   FullScreenMode
}

// /////////////////////////////////////////////////////////////////////////////

var _ HyprCommand = &FullScreenState{}

type FullScreenState struct {
	State FullscreenState2
}

func (this FullScreenState) String() string {
	return fmt.Sprintf("dispatch fullscreenstate %d", this.State)
}

func (_ FullScreenState) hyprCommand() HyprCommandType {
	return HyprCommandTypes.FullscreenState
}

type FullscreenState2 int

func (this FullscreenState2) String() string {
	return strconv.Itoa(int(this))
}

var FullScreenStates = fullscreenStates{
	Current:               -1,
	None:                  0,
	Maximize:              1,
	Fullscreen:            2,
	MaximizeAndFullscreen: 3,
}

type fullscreenStates struct {
	Current               FullscreenState2
	None                  FullscreenState2
	Maximize              FullscreenState2
	Fullscreen            FullscreenState2
	MaximizeAndFullscreen FullscreenState2
}

// /////////////////////////////////////////////////////////////////////////////

var _ HyprCommand = &Dpms{}

type Dpms struct {
	Status DpmsStatus
}

func (this Dpms) String() string {
	return fmt.Sprintf("dispatch dpms %s", this.Status)
}

func (_ Dpms) hyprCommand() HyprCommandType {
	return HyprCommandTypes.Dpms
}

type DpmsStatus string

func (this DpmsStatus) String() string {
	return string(this)
}

var DpmsStatuses = dpmsStatuses{
	On:     "on",
	Off:    "off",
	Toggle: "toggle",
}

type dpmsStatuses struct {
	On     DpmsStatus
	Off    DpmsStatus
	Toggle DpmsStatus
}

func (this dpmsStatuses) Parse(input string) (DpmsStatus, error) {
	var found DpmsStatus
	var err error

	switch strings.ToLower(input) {
	case "on":
		found = this.On
	case "off":
		found = this.Off
	case "toggle":
		found = this.Toggle
	default:
		err = fmt.Errorf("%s is invalid DPMS status", input)
	}

	return found, err
}

func (this dpmsStatuses) All() []DpmsStatus {
	return []DpmsStatus{this.On, this.Off, this.Toggle}
}

// /////////////////////////////////////////////////////////////////////////////

var _ HyprCommand = &Pin{}

type Pin struct {
	Window model.HyprWindowId
}

func (this Pin) String() string {
	if this.Window != nil {
		return fmt.Sprintf("dispatch settiled %s", this.Window)
	}

	return "dispatch settiled"
}

func (_ Pin) hyprCommand() HyprCommandType {
	return HyprCommandTypes.Pin
}

// /////////////////////////////////////////////////////////////////////////////

var _ HyprCommand = &MoveFocus{}

type MoveFocus struct {
	Direction model.HyprDirection
}

func (this MoveFocus) String() string {
	return fmt.Sprintf("dispatch movefocus %s", this.Direction)
}

func (_ MoveFocus) hyprCommand() HyprCommandType {
	return HyprCommandTypes.MoveFocus
}

// /////////////////////////////////////////////////////////////////////////////

var _ HyprCommand = &MoveWindow{}

type MoveWindow struct {
	Direction model.HyprDirection
	Monitor   model.MonitorId
	Silent    bool
}

func (this MoveWindow) String() string {
	text := "dispatch movewindow "

	switch {
	case this.Monitor == nil && this.Direction != model.HyprDirections.None:
		text += this.Direction.String()
	case this.Monitor != nil && this.Direction == model.HyprDirections.None:
		text += this.Monitor.String()
	default:
		panic(errors.New("TODO"))
	}

	if this.Silent {
		text += " silent"
	}

	return text
}

func (_ MoveWindow) hyprCommand() HyprCommandType {
	return HyprCommandTypes.MoveWindow
}

// /////////////////////////////////////////////////////////////////////////////

var _ HyprCommand = &SwapWindow{}

type SwapWindow struct {
	Direction model.HyprDirection
	Monitor   model.MonitorId
}

func (this SwapWindow) String() string {
	text := "dispatch swapwindow "

	switch {
	case this.Monitor == nil && this.Direction != model.HyprDirections.None:
		text += this.Direction.String()
	case this.Monitor != nil && this.Direction == model.HyprDirections.None:
		text += this.Monitor.String()
	default:
		panic(errors.New("TODO"))
	}

	return text
}

func (_ SwapWindow) hyprCommand() HyprCommandType {
	return HyprCommandTypes.SwapWindow
}

// /////////////////////////////////////////////////////////////////////////////

var _ HyprCommand = &CenterWindow{}

type CenterWindow struct {
	RespectMonitorReservedArea bool
}

func (this CenterWindow) String() string {
	text := "dispatch centerwindow"

	if this.RespectMonitorReservedArea {
		text += " 1"
	}

	return text
}

func (_ CenterWindow) hyprCommand() HyprCommandType {
	return HyprCommandTypes.CenterWindow
}

// /////////////////////////////////////////////////////////////////////////////

var _ HyprCommand = &ResizeActive{}

type ResizeActive struct {
	Params model.ResizeParam
}

func (this ResizeActive) String() string {
	text := fmt.Sprintf("dispatch resizeactive  %s", this.Params)
	return text
}

func (_ ResizeActive) hyprCommand() HyprCommandType {
	return HyprCommandTypes.ResizeActive
}

// /////////////////////////////////////////////////////////////////////////////

var _ HyprCommand = &MoveActive{}

type MoveActive struct {
	Params model.ResizeParam
}

func (this MoveActive) String() string {
	text := fmt.Sprintf("dispatch moveactive  %s", this.Params)
	return text
}

func (_ MoveActive) hyprCommand() HyprCommandType {
	return HyprCommandTypes.MoveActive
}

// /////////////////////////////////////////////////////////////////////////////

var _ HyprCommand = &ResizeWindowPixel{}

type ResizeWindowPixel struct {
	Params model.ResizeParam
	Window model.HyprWindowId
}

func (this ResizeWindowPixel) String() string {
	text := fmt.Sprintf("dispatch resizewindowpixel  %s,%s", this.Params, this.Window)
	return text
}

func (_ ResizeWindowPixel) hyprCommand() HyprCommandType {
	return HyprCommandTypes.ResizeWindowPixel
}

// /////////////////////////////////////////////////////////////////////////////

var _ HyprCommand = &MoveWindowPixel{}

type MoveWindowPixel struct {
	Params model.ResizeParam
	Window model.HyprWindowId
}

func (this MoveWindowPixel) String() string {
	text := fmt.Sprintf("dispatch movewindowpixel  %s,%s", this.Params, this.Window)
	return text
}

func (_ MoveWindowPixel) hyprCommand() HyprCommandType {
	return HyprCommandTypes.MoveWindowPixel
}

// /////////////////////////////////////////////////////////////////////////////

var _ HyprCommand = &CycleNext{}

type CycleNext struct {
	Target CycleNextTarget
}

func (this CycleNext) String() string {
	text := fmt.Sprintf("dispatch cyclenext  %s", this.Target)
	return text
}

func (_ CycleNext) hyprCommand() HyprCommandType {
	return HyprCommandTypes.CycleNext
}

type CycleNextTarget string

func (this CycleNextTarget) String() string {
	return string(this)
}

var CycleNextTargets = cycleNextTargets{
	Next:            "",
	Prev:            "prev",
	Titled:          "tiled",
	Floating:        "floating",
	PrevTiled:       "prev tilted",
	Visible:         "visible",
	VisblePrevTiled: "visible prev titled",
	Hist:            "hist",
}

type cycleNextTargets struct {
	Next            CycleNextTarget
	Prev            CycleNextTarget
	Titled          CycleNextTarget
	Floating        CycleNextTarget
	PrevTiled       CycleNextTarget
	Visible         CycleNextTarget
	VisblePrevTiled CycleNextTarget
	Hist            CycleNextTarget
}

// /////////////////////////////////////////////////////////////////////////////

var _ HyprCommand = &SwapNext{}

type SwapNext struct {
	Target SwapNextTarget
}

func (this SwapNext) String() string {
	text := fmt.Sprintf("dispatch swapnext  %s", this.Target)
	return text
}

func (_ SwapNext) hyprCommand() HyprCommandType {
	return HyprCommandTypes.SwapNext
}

type SwapNextTarget string

func (this SwapNextTarget) String() string {
	return string(this)
}

var SwapNextTargets = swapNextTargets{
	Next: "",
	Prev: "prev",
}

type swapNextTargets struct {
	Next SwapNextTarget
	Prev SwapNextTarget
}

// /////////////////////////////////////////////////////////////////////////////

var _ HyprCommand = &TagWindow{}

type TagWindow struct {
	Window model.HyprWindowId
	Tag    string
}

func (this TagWindow) String() string {
	text := fmt.Sprintf("dispatch tagwindow %s", this.Tag)

	if this.Window != nil {
		text += " " + this.Window.String()
	}

	return text
}

func (_ TagWindow) hyprCommand() HyprCommandType {
	return HyprCommandTypes.TagWindow
}

// /////////////////////////////////////////////////////////////////////////////

var _ HyprCommand = &FocusWindow{}

type FocusWindow struct {
	Window model.HyprWindowId
}

func (this FocusWindow) String() string {
	text := fmt.Sprintf("dispatch focuswindow %s", this.Window)

	return text
}

func (_ FocusWindow) hyprCommand() HyprCommandType {
	return HyprCommandTypes.FocusWindow
}

// /////////////////////////////////////////////////////////////////////////////

var _ HyprCommand = &FocusMonitor{}

type FocusMonitor struct {
	Monitor model.MonitorId
}

func (this FocusMonitor) String() string {
	text := fmt.Sprintf("dispatch focusmonitor %s", this.Monitor)

	return text
}

func (_ FocusMonitor) hyprCommand() HyprCommandType {
	return HyprCommandTypes.FocusMonitor
}

// /////////////////////////////////////////////////////////////////////////////

var _ HyprCommand = &SplitRatio{}

type SplitRatio struct {
	Ratio float32
}

func (this SplitRatio) String() string {
	text := fmt.Sprintf("dispatch splitratio %f", this.Ratio)

	return text
}

func (_ SplitRatio) hyprCommand() HyprCommandType {
	return HyprCommandTypes.SplitRatio
}

// /////////////////////////////////////////////////////////////////////////////

var _ HyprCommand = &MoveCursorToCorner{}

type MoveCursorToCorner struct {
	Target MoveCursorToCornerTarget
}

func (this MoveCursorToCorner) String() string {
	text := fmt.Sprintf("dispatch movecursortocorner %s", this.Target)

	return text
}

func (_ MoveCursorToCorner) hyprCommand() HyprCommandType {
	return HyprCommandTypes.MoveCursorToCorner
}

type MoveCursorToCornerTarget int

func (this MoveCursorToCornerTarget) String() string {
	return strconv.Itoa(int(this))
}

var MoveCursorToCornerTargets = moveCursorToCornerTargets{
	BottomLeft:  0,
	BottomRight: 1,
	TopRight:    2,
	TopLeft:     3,
}

type moveCursorToCornerTargets struct {
	BottomLeft  MoveCursorToCornerTarget
	BottomRight MoveCursorToCornerTarget
	TopRight    MoveCursorToCornerTarget
	TopLeft     MoveCursorToCornerTarget
}

// /////////////////////////////////////////////////////////////////////////////

var _ HyprCommand = &MoveCursor{}

type MoveCursor struct {
	X int
	Y int
}

func (this MoveCursor) String() string {
	text := fmt.Sprintf("dispatch movecursor %d %d", this.X, this.Y)
	return text
}

func (_ MoveCursor) hyprCommand() HyprCommandType {
	return HyprCommandTypes.MoveCursor
}

// /////////////////////////////////////////////////////////////////////////////

var _ HyprCommand = &RenameWorkspace{}

type RenameWorkspace struct {
	Workspace model.HyprWorkspaceId
	Name      string
}

func (this RenameWorkspace) String() string {
	text := fmt.Sprintf("dispatch renameworkspace %s %s", this.Workspace, this.Name)
	return text
}

func (_ RenameWorkspace) hyprCommand() HyprCommandType {
	return HyprCommandTypes.RenameWorkspace
}

// /////////////////////////////////////////////////////////////////////////////

var _ HyprCommand = &Exit{}

type Exit struct {
	Workspace model.HyprWorkspaceId
	Name      string
}

func (this Exit) String() string {
	return "dispatch exit"
}

func (_ Exit) hyprCommand() HyprCommandType {
	return HyprCommandTypes.Exit
}

// /////////////////////////////////////////////////////////////////////////////

var _ HyprCommand = &ForceRendererReload{}

type ForceRendererReload struct {
	Workspace model.HyprWorkspaceId
	Name      string
}

func (this ForceRendererReload) String() string {
	return "dispatch forcerendererreload"
}

func (_ ForceRendererReload) hyprCommand() HyprCommandType {
	return HyprCommandTypes.ForceRendererReload
}

// /////////////////////////////////////////////////////////////////////////////

var _ HyprCommand = &MoveCurrentWorkspaceToMonitor{}

type MoveCurrentWorkspaceToMonitor struct {
	Monitor model.MonitorId
}

func (this MoveCurrentWorkspaceToMonitor) String() string {
	text := fmt.Sprintf("dispatch movecurrentworkspacetomonitor %s", this.Monitor)
	return text
}

func (_ MoveCurrentWorkspaceToMonitor) hyprCommand() HyprCommandType {
	return HyprCommandTypes.MoveCurrentWorkspaceToMonitor
}

// /////////////////////////////////////////////////////////////////////////////

var _ HyprCommand = &FocusWorkspaceOnCurrentMonitor{}

type FocusWorkspaceOnCurrentMonitor struct {
	Workspace model.HyprWorkspaceId
}

func (this FocusWorkspaceOnCurrentMonitor) String() string {
	text := fmt.Sprintf("dispatch focusworkspaceoncurrentmonitor %s", this.Workspace)
	return text
}

func (_ FocusWorkspaceOnCurrentMonitor) hyprCommand() HyprCommandType {
	return HyprCommandTypes.FocusWorkspaceOnCurrentMonitor
}

// /////////////////////////////////////////////////////////////////////////////

var _ HyprCommand = &MoveWorkspaceToMonitor{}

type MoveWorkspaceToMonitor struct {
	Workspace model.HyprWorkspaceId
	Monitor   model.MonitorId
}

func (this MoveWorkspaceToMonitor) String() string {
	text := fmt.Sprintf("dispatch moveworkspacetomonitor %s %s", this.Workspace, this.Monitor)
	return text
}

func (_ MoveWorkspaceToMonitor) hyprCommand() HyprCommandType {
	return HyprCommandTypes.MoveWorkspaceToMonitor
}

// /////////////////////////////////////////////////////////////////////////////

var _ HyprCommand = &SwapActiveWorkspaces{}

type SwapActiveWorkspaces struct {
	Monitor1 model.MonitorId
	Monitor2 model.MonitorId
}

func (this SwapActiveWorkspaces) String() string {
	text := fmt.Sprintf("dispatch swapactiveworkspaces %s %s", this.Monitor1, this.Monitor2)
	return text
}

func (_ SwapActiveWorkspaces) hyprCommand() HyprCommandType {
	return HyprCommandTypes.SwapActiveWorkspaces
}

// /////////////////////////////////////////////////////////////////////////////

var _ HyprCommand = &AlterZorder{}

type AlterZorder struct {
	Zheight Zheight
	Window  model.HyprWindowId
}

func (this AlterZorder) String() string {
	text := fmt.Sprintf("dispatch alterzorder %s%s", this.Zheight, model.Windows.AppendIf(this.Window))
	return text
}

func (_ AlterZorder) hyprCommand() HyprCommandType {
	return HyprCommandTypes.AlterZorder
}

type Zheight string

func (this Zheight) String() string {
	return string(this)
}

var Zheights = zheights{
	Top:    "top",
	Bottom: "bottom",
}

type zheights struct {
	Top    Zheight
	Bottom Zheight
}

// /////////////////////////////////////////////////////////////////////////////

var _ HyprCommand = &ToggleSpecialWorkspace{}

type ToggleSpecialWorkspace struct {
	Workspace model.HyprWorkspaceId
}

func (this ToggleSpecialWorkspace) String() string {
	text := fmt.Sprintf("dispatch togglespecialworkspace %s", model.Workspaces.AppendIf(this.Workspace))

	return text
}

func (_ ToggleSpecialWorkspace) hyprCommand() HyprCommandType {
	return HyprCommandTypes.ToggleSpecialWorkspace
}

// /////////////////////////////////////////////////////////////////////////////

var _ HyprCommand = &FocusUrgentOrLast{}

type FocusUrgentOrLast struct{}

func (this FocusUrgentOrLast) String() string {
	return "dispatch focusurgentorlast"
}

func (_ FocusUrgentOrLast) hyprCommand() HyprCommandType {
	return HyprCommandTypes.FocusUrgentOrLast
}

// /////////////////////////////////////////////////////////////////////////////

var _ HyprCommand = &ToggleGroup{}

type ToggleGroup struct{}

func (this ToggleGroup) String() string {
	return "dispatch togglegroup"
}

func (_ ToggleGroup) hyprCommand() HyprCommandType {
	return HyprCommandTypes.ToggleGroup
}

// /////////////////////////////////////////////////////////////////////////////

var _ HyprCommand = &ChangeGroupActive{}

type ChangeGroupActive struct {
	Target ChangeGroupActiveTarget
}

func (this ChangeGroupActive) String() string {
	text := fmt.Sprintf("dispatch changegroupactive %s", this.Target)
	return text
}

func (_ ChangeGroupActive) hyprCommand() HyprCommandType {
	return HyprCommandTypes.ChangeGroupActive
}

type ChangeGroupActiveTarget interface {
	String() string
	changeGroupActiveTarget()
}

type changeGroupActiveTarget struct {
	value string
}

func (this changeGroupActiveTarget) String() string {
	return string(this.value)
}

func (_ changeGroupActiveTarget) changeGroupActiveTarget() {}

var ChangeGroupActiveTargets = changeGroupActiveTargets{
	Back:    &changeGroupActiveTarget{value: "b"},
	Forword: &changeGroupActiveTarget{value: "f"},
	Index:   func(index int) ChangeGroupActiveTarget { return &changeGroupActiveTarget{value: strconv.Itoa(index)} },
}

type changeGroupActiveTargets struct {
	Back    ChangeGroupActiveTarget
	Forword ChangeGroupActiveTarget
	Index   func(int) ChangeGroupActiveTarget
}

// /////////////////////////////////////////////////////////////////////////////

var _ HyprCommand = &FocusCurrentOrLast{}

type FocusCurrentOrLast struct{}

func (this FocusCurrentOrLast) String() string {
	return "dispatch focuscurrentorlast"
}

func (_ FocusCurrentOrLast) hyprCommand() HyprCommandType {
	return HyprCommandTypes.FocusCurrentOrLast
}

// /////////////////////////////////////////////////////////////////////////////

var _ HyprCommand = &LockGroups{}

type LockGroups struct {
	Target LockGroupsTarget
}

func (this LockGroups) String() string {
	text := fmt.Sprintf("dispatch lockgroups %s", this.Target)
	return text
}

func (_ LockGroups) hyprCommand() HyprCommandType {
	return HyprCommandTypes.LockGroups
}

type LockGroupsTarget interface {
	String() string
	lockGroupsTarget()
}

type lockGroupsTarget struct {
	value string
}

func (this lockGroupsTarget) String() string {
	return string(this.value)
}

func (_ lockGroupsTarget) lockGroupsTarget() {}

var LockGroupsTargets = lockGroupsTargets{
	Lock:   &lockGroupsTarget{value: "lock"},
	Unlock: &lockGroupsTarget{value: "unlock"},
	Toggle: &lockGroupsTarget{value: "toggle"},
}

type lockGroupsTargets struct {
	Lock   LockGroupsTarget
	Unlock LockGroupsTarget
	Toggle LockGroupsTarget
}

// /////////////////////////////////////////////////////////////////////////////

var _ HyprCommand = &LockActiveGroup{}

type LockActiveGroup struct {
	Target ChangeGroupActiveTarget
}

func (this LockActiveGroup) String() string {
	text := fmt.Sprintf("dispatch lockactivegroup %s", this.Target)
	return text
}

func (_ LockActiveGroup) hyprCommand() HyprCommandType {
	return HyprCommandTypes.LockActiveGroup
}

// /////////////////////////////////////////////////////////////////////////////

var _ HyprCommand = &MoveIntoGroup{}

type MoveIntoGroup struct {
	Direction model.HyprDirection
}

func (this MoveIntoGroup) String() string {
	text := fmt.Sprintf("dispatch moveintogroup %s", this.Direction)
	return text
}

func (_ MoveIntoGroup) hyprCommand() HyprCommandType {
	return HyprCommandTypes.MoveIntoGroup
}

// /////////////////////////////////////////////////////////////////////////////

var _ HyprCommand = &MoveOutOfGroup{}

type MoveOutOfGroup struct {
	Window model.HyprWindowId
}

func (this MoveOutOfGroup) String() string {
	text := fmt.Sprintf("dispatch moveoutofgroup %s", model.Windows.AppendIf(this.Window))
	return text
}

func (_ MoveOutOfGroup) hyprCommand() HyprCommandType {
	return HyprCommandTypes.MoveOutOfGroup
}

// /////////////////////////////////////////////////////////////////////////////

var _ HyprCommand = &MoveWindowOrGroup{}

type MoveWindowOrGroup struct {
	Direction model.HyprDirection
}

func (this MoveWindowOrGroup) String() string {
	text := fmt.Sprintf("dispatch movewindoworgroup %s", this.Direction)
	return text
}

func (_ MoveWindowOrGroup) hyprCommand() HyprCommandType {
	return HyprCommandTypes.MoveWindowOrGroup
}

// /////////////////////////////////////////////////////////////////////////////

var _ HyprCommand = &MoveGroupWindow{}

type MoveGroupWindow struct {
	Target MoveGroupWindowTarget
}

func (this MoveGroupWindow) String() string {
	text := fmt.Sprintf("dispatch movegroupwindow %s", this.Target)
	return text
}

func (_ MoveGroupWindow) hyprCommand() HyprCommandType {
	return HyprCommandTypes.MoveGroupWindow
}

type MoveGroupWindowTarget interface {
	String() string
	moveGroupWindowTarget()
}

type moveGroupWindowTarget struct {
	value string
}

func (this moveGroupWindowTarget) String() string {
	return string(this.value)
}

func (_ moveGroupWindowTarget) moveGroupWindowTarget() {}

var MoveGroupWindowTargets = moveGroupWindowTargets{
	Back:    &moveGroupWindowTarget{value: "b"},
	Forword: &moveGroupWindowTarget{value: "f"},
}

type moveGroupWindowTargets struct {
	Back    MoveGroupWindowTarget
	Forword MoveGroupWindowTarget
}

// /////////////////////////////////////////////////////////////////////////////

var _ HyprCommand = &DenyWindowFromGroup{}

type DenyWindowFromGroup struct {
	Target DenyWindowFromGroupTarget
}

func (this DenyWindowFromGroup) String() string {
	text := fmt.Sprintf("dispatch denywindowfromgroup %s", this.Target)
	return text
}

func (_ DenyWindowFromGroup) hyprCommand() HyprCommandType {
	return HyprCommandTypes.DenyWindowFromGroup
}

type DenyWindowFromGroupTarget interface {
	String() string
	denyWindowFromGroupTarget()
}

type denyWindowFromGroupTarget struct {
	value string
}

func (this denyWindowFromGroupTarget) String() string {
	return string(this.value)
}

func (_ denyWindowFromGroupTarget) denyWindowFromGroupTarget() {}

var DenyWindowFromGroupTargets = denyWindowFromGroupTargets{
	On:     &denyWindowFromGroupTarget{value: "on"},
	Off:    &denyWindowFromGroupTarget{value: "off"},
	Toggle: &denyWindowFromGroupTarget{value: "toggle"},
}

type denyWindowFromGroupTargets struct {
	On     DenyWindowFromGroupTarget
	Off    DenyWindowFromGroupTarget
	Toggle DenyWindowFromGroupTarget
}

// /////////////////////////////////////////////////////////////////////////////

var _ HyprCommand = &SetIgnoreGroupLock{}

type SetIgnoreGroupLock struct {
	Target DenyWindowFromGroupTarget
}

func (this SetIgnoreGroupLock) String() string {
	text := fmt.Sprintf("dispatch setignoregrouplock %s", this.Target)
	return text
}

func (_ SetIgnoreGroupLock) hyprCommand() HyprCommandType {
	return HyprCommandTypes.SetIgnoreGroupLock
}

// /////////////////////////////////////////////////////////////////////////////

var _ HyprCommand = &Global{}

type Global struct {
	Name string
}

func (this Global) String() string {
	text := fmt.Sprintf("dispatch global %s", this.Name)

	return text
}

func (_ Global) hyprCommand() HyprCommandType {
	return HyprCommandTypes.Global
}

// /////////////////////////////////////////////////////////////////////////////

var _ HyprCommand = &DispatchSubmap{}

type DispatchSubmap struct {
	Target DenyWindowFromGroupTarget
}

func (this DispatchSubmap) String() string {
	text := fmt.Sprintf("dispatch submap %s", this.Target)
	return text
}

func (_ DispatchSubmap) hyprCommand() HyprCommandType {
	return HyprCommandTypes.Submap
}

type DispatchSubmapTarget interface {
	String() string
	submapTarget()
}

type dispatchSubmapTarget struct {
	value string
}

func (this dispatchSubmapTarget) String() string {
	return string(this.value)
}

func (_ dispatchSubmapTarget) submapTarget() {}

var DispatchSubmapTargets = dispatchSubmapTargets{
	Reset: &dispatchSubmapTarget{value: "reset"},
	Name:  func(name string) DispatchSubmapTarget { return &dispatchSubmapTarget{value: name} },
}

type dispatchSubmapTargets struct {
	Reset DispatchSubmapTarget
	Name  func(string) DispatchSubmapTarget
}

// /////////////////////////////////////////////////////////////////////////////

var _ HyprCommand = &Event{}

type Event struct {
	Event events.HyprEvent
}

func (this Event) String() string {
	text := fmt.Sprintf("dispatch event %s", this.Event)
	return text
}

func (_ Event) hyprCommand() HyprCommandType {
	return HyprCommandTypes.Event
}

// /////////////////////////////////////////////////////////////////////////////

var _ HyprCommand = &SetProp{}

type SetProp struct {
	Prop string
}

func (this SetProp) String() string {
	text := fmt.Sprintf("dispatch setprop %s", this.Prop)
	return text
}

func (_ SetProp) hyprCommand() HyprCommandType {
	return HyprCommandTypes.SetProp
}

// /////////////////////////////////////////////////////////////////////////////

var _ HyprCommand = &ToggleSwallow{}

type ToggleSwallow struct{}

func (this ToggleSwallow) String() string {
	return "dispatch toggleswallow"
}

func (_ ToggleSwallow) hyprCommand() HyprCommandType {
	return HyprCommandTypes.ToggleSwallow
}

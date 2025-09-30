package model

import (
	"fmt"
	"strconv"
	"strings"
)

// /////////////////////////////////////////////////////////////////////////////
// Window ID
// /////////////////////////////////////////////////////////////////////////////

var Windows = windows{}

type windows struct {
	Id windowIds
}

type windowIds struct{}

func (_ windowIds) Class(regex string) HyprWindowId {
	return &windowId{line: fmt.Sprintf("class:%s", regex)}
}

func (_ windowIds) InitialClass(regex string) HyprWindowId {
	return &windowId{line: fmt.Sprintf("initialclass:%s", regex)}
}

func (_ windowIds) Title(regex string) HyprWindowId {
	return &windowId{line: fmt.Sprintf("title:%s", regex)}
}

func (_ windowIds) InitialTitle(regex string) HyprWindowId {
	return &windowId{line: fmt.Sprintf("initialtitle:%s", regex)}
}

func (_ windowIds) Tag(regex string) HyprWindowId {
	return &windowId{line: fmt.Sprintf("tag:%s", regex)}
}

func (_ windowIds) Pid(pid int64) HyprWindowId {
	return &windowId{line: fmt.Sprintf("pid:%d", pid)}
}

func (_ windowIds) Address(address string) HyprWindowId {
	return &windowId{line: fmt.Sprintf("address:%s", address)}
}

func (_ windowIds) ActiveWindow(address string) HyprWindowId {
	return &windowId{line: "activewindow"}
}

func (_ windowIds) Floating(address string) HyprWindowId {
	return &windowId{line: "floating"}
}

func (_ windowIds) Tiled(address string) HyprWindowId {
	return &windowId{line: "tiled"}
}

func (_ windows) AppendIf(window HyprWindowId) string {
	if window == nil {
		return ""
	}

	return " " + window.String()
}

type HyprWindowId interface {
	String() string
	windowId() // discriminator
}

type windowId struct {
	line string
}

func (this windowId) String() string {
	return this.line
}

func (_ windowId) windowId() {}

// /////////////////////////////////////////////////////////////////////////////
//  Monitor ID
// /////////////////////////////////////////////////////////////////////////////

type MonitorId interface {
	String() string
	monitorId()
}

// /////////////////////////////////////////////////////////////////////////////
//  Workspace ID
// /////////////////////////////////////////////////////////////////////////////

var Workspaces = workspaces{}

type workspaces struct {
	Id workspaceIds
}

func (_ workspaces) AppendIf(workspace HyprWorkspaceId) string {
	if workspace != nil {
		return " " + workspace.String()
	}
	return ""
}

type workspaceIds struct{}

func (this workspaceIds) Id(id int) HyprWorkspaceId {
	return &hyprWorkSpaceId{line: strconv.Itoa(id)}
}

func (this workspaceIds) Name(name string) HyprWorkspaceId {
	return &hyprWorkSpaceId{line: fmt.Sprintf("name:%s", name)}
}

func (this workspaceIds) Special(name string) HyprWorkspaceId {
	return &hyprWorkSpaceId{line: fmt.Sprintf("special:%s", name)}
}

type HyprWorkspaceId interface {
	String() string
	workspaceId() // discriminator
}

type hyprWorkSpaceId struct {
	line string
}

func (this hyprWorkSpaceId) String() string {
	return this.line
}

func (_ hyprWorkSpaceId) workspaceId() {}

// /////////////////////////////////////////////////////////////////////////////
// Notify Icon
// /////////////////////////////////////////////////////////////////////////////

type HyprNotifyIcon int

func (this HyprNotifyIcon) String() string {
	return strconv.Itoa(int(this))
}

var NotifyIcons = notifyIcons{
	NONE:     -1,
	WARNING:  0,
	INFO:     1,
	HINT:     2,
	ERROR:    3,
	CONFUSED: 4,
	OK:       5,
}

type notifyIcons struct {
	NONE     HyprNotifyIcon
	WARNING  HyprNotifyIcon
	INFO     HyprNotifyIcon
	HINT     HyprNotifyIcon
	ERROR    HyprNotifyIcon
	CONFUSED HyprNotifyIcon
	OK       HyprNotifyIcon
}

// /////////////////////////////////////////////////////////////////////////////
// Direction
// /////////////////////////////////////////////////////////////////////////////

// type HyprDirection2 interface {
// 	String() string
// 	hyprDirection()
// }

// type hyprDirection2 struct {
// 	direction rune
// }

// func (this hyprDirection2) String() string {
// 	return string(this.direction)
// }

// func (this hyprDirection2) hyprDirection() {}

type HyprDirection string

func (this HyprDirection) String() string {
	return string(this)
}

var HyprDirections = hyprDirections{
	None:  "-",
	Up:    "u",
	Down:  "d",
	Left:  "l",
	Right: "r",
}

type hyprDirections struct {
	None  HyprDirection
	Up    HyprDirection
	Down  HyprDirection
	Left  HyprDirection
	Right HyprDirection
}

func (this hyprDirections) Parse(input string) (HyprDirection, error) {
	var found HyprDirection
	var err error

	switch strings.ToLower(input) {
	case "u":
		found = this.Up
	case "d":
		found = this.Down
	case "l":
		found = this.Left
	case "r":
		found = this.Right
	default:
		err = fmt.Errorf("%s is invalid direction", input)
	}

	return found, err
}

func (this hyprDirections) All() []HyprDirection {
	return []HyprDirection{this.Up, this.Down, this.Left, this.Right}
}

// /////////////////////////////////////////////////////////////////////////////
// Resize Params
// /////////////////////////////////////////////////////////////////////////////

type ResizeParam interface {
	String() string
	isResizeParam() // descriminator
}

var ResizeParams = resizeParams{}

type resizeParams struct {
}

func (_ resizeParams) Relative(x, y int) ResizeParam {
	return &realativePixel{x, y}
}

func (_ resizeParams) Exact(x, y int) ResizeParam {
	return &exactPixel{x, y}
}

type realativePixel struct {
	x, y int
}

func (this realativePixel) String() string {
	return fmt.Sprintf("%d %d", this.x, this.y)
}

func (this realativePixel) isResizeParam() {}

type exactPixel struct {
	x, y int
}

func (this exactPixel) String() string {
	return fmt.Sprintf("exact %d %d", this.x, this.y)
}

func (this exactPixel) isResizeParam() {}

// /////////////////////////////////////////////////////////////////////////////
// Mod Key Param
// /////////////////////////////////////////////////////////////////////////////

type ModKey string

func (this ModKey) String() string {
	return string(this)
}

var ModKeys = modkeys{
	Super:    "SuperAll",
	SuperAlt: "SUPER_ALT",
}

type modkeys struct {
	Super    ModKey
	SuperAlt ModKey
}

// /////////////////////////////////////////////////////////////////////////////
// Key State Param
// /////////////////////////////////////////////////////////////////////////////

type KeyState string

func (this KeyState) String() string {
	return string(this)
}

var KeyStates = keyStates{
	Up:     "up",
	Down:   "down",
	Repeat: "repeat",
}

type keyStates struct {
	Up     KeyState
	Down   KeyState
	Repeat KeyState
}

package events

import (
	"errors"
	"go-motions/hypr/model"
	"strings"
)

type HyprEvent interface {
	// String() string
	hyprEvent() // discriminator
}

var (
	ErrEventDetails = errors.New("invalid event details")
)

func Parse(line string) (HyprEvent, error) {
	var event HyprEvent
	var err error

	typ, details, ok := splitOnEvent(line)

	if !ok {
		return event, &HyprEventError{
			Text: line,
			Err:  errors.New("invalid event"),
		}
	}

	var parts []string

	switch HyprEventType(typ) {
	case HyprEventTypes.ActiveWindow:
		if parts, err = splitOnDetail(details, 2); err == nil {
			event = &ActiveWindow{
				Class: model.Windows.Id.Class(parts[0]),
				Title: model.Windows.Id.Title(parts[1]),
			}
		}
	case HyprEventTypes.ActiveWindowV2:
		event = &ActiveWindowV2{
			Address: model.Windows.Id.Address(details),
		}
	case HyprEventTypes.OpenWindow:
		if parts, err = splitOnDetail(details, 3); err == nil {
			event = &OpenWindow{
				Address:       model.Windows.Id.Address(parts[0]),
				WorkspaceName: parts[1],
				Class:         model.Windows.Id.Class(parts[2]),
				Title:         model.Windows.Id.Title(parts[3]),
			}
		}
	case HyprEventTypes.CloseWindow:
		event = &CloseWindow{Address: details}
	case HyprEventTypes.MoveWindowV2:
		if parts, err = splitOnDetail(details, 3); err == nil {
			event = &MoveWindow{
				Address:       parts[0],
				WorkspaceId:   parts[1],
				WorkspaceName: parts[2],
			}
		}
	case HyprEventTypes.WindowTitleV2:
		if parts, err = splitOnDetail(details, 2); err == nil {
			event = &WindowTitle{
				Address: parts[0],
				Title:   parts[1],
			}
		}
	case HyprEventTypes.ChangeFloatingMode:
		if parts, err = splitOnDetail(details, 2); err == nil {
			switch parts[1] {
			case "0":
				event = &ChangeFloatingMode{
					Address:    parts[0],
					IsFloating: false,
				}
			case "1":
				event = &ChangeFloatingMode{
					Address:    parts[0],
					IsFloating: true,
				}
			default:
				err = ErrEventDetails
			}
		}
	case HyprEventTypes.Bell:
		event = &Bell{Address: details}
	case HyprEventTypes.Pin:
		if parts, err = splitOnDetail(details, 2); err == nil {
			temp := &Pin{Address: parts[0]}
			switch parts[1] {
			case "0":
				temp.IsPinned = false
			case "1":
				temp.IsPinned = true
			default:
				err = ErrEventDetails
			}
			event = temp
		}
	case HyprEventTypes.Minimized:
		if parts, err = splitOnDetail(details, 2); err == nil {
			temp := &Minimized{Address: parts[0]}
			switch parts[1] {
			case "0":
				temp.IsMinimized = false
			case "1":
				temp.IsMinimized = true
			default:
				err = ErrEventDetails
			}
			event = temp
		}
	case HyprEventTypes.Urgent:
		event = &WindowUrgent{Address: details}
	case HyprEventTypes.WorkspaceV2:
		event = &Workspace{Name: details}
	case HyprEventTypes.WorkspaceV2:
		if parts, err = splitOnDetail(details, 2); err == nil {
			event = &WorkspaceV2{
				Id:   parts[0],
				Name: parts[1],
			}
		}
	case HyprEventTypes.MoveWorkspace:
		if parts, err = splitOnDetail(details, 2); err == nil {
			event = &MoveWorkspace{
				Name:        parts[0],
				MonitorName: parts[1],
			}
		}
	case HyprEventTypes.MoveWorkspaceV2:
		if parts, err = splitOnDetail(details, 3); err == nil {
			event = &MoveWorkspaceV2{
				Id:          parts[0],
				Name:        parts[1],
				MonitorName: parts[2],
			}
		}
	case HyprEventTypes.RenameWorkspace:
		if parts, err = splitOnDetail(details, 2); err == nil {
			event = &RenameWorkspace{
				Id:      parts[0],
				NewName: parts[1],
			}
		}
	case HyprEventTypes.CreateWorkspace:
		event = &CreateWorkspace{Name: details}
	case HyprEventTypes.CreateWorkspaceV2:
		if parts, err = splitOnDetail(details, 2); err == nil {
			event = &CreateWorkspaceV2{
				Id:   parts[0],
				Name: parts[1],
			}
		}
	case HyprEventTypes.DestroyWorkspace:
		event = &DestroyWorkspace{Name: details}
	case HyprEventTypes.DestroyWorkspaceV2:
		if parts, err = splitOnDetail(details, 2); err == nil {
			event = &DestroyWorkspaceV2{
				Id:   parts[0],
				Name: parts[1],
			}
		}
	case HyprEventTypes.FocusedMon:
		if parts, err = splitOnDetail(details, 2); err == nil {
			event = &MonitorFocused{
				Name:          parts[0],
				WorkspaceName: parts[1],
			}
		}
	case HyprEventTypes.FocusedMonV2:
		if parts, err = splitOnDetail(details, 2); err == nil {
			event = &MonitorFocusedV2{
				Name:        parts[0],
				WorkspaceId: parts[1],
			}
		}
	case HyprEventTypes.FullScreen:
		switch details {
		case "0":
			event = &Fullscreen{IsFullScreen: false}
		case "1":
			event = &Fullscreen{IsFullScreen: true}
		default:
			err = ErrEventDetails
		}
	case HyprEventTypes.MonitorRemovedV2:
		if parts, err = splitOnDetail(details, 3); err == nil {
			event = &MonitorRemoved{
				Id:          parts[0],
				Name:        parts[1],
				Description: parts[2],
			}
		}
	case HyprEventTypes.MonitorAddedV2:
		if parts, err = splitOnDetail(details, 3); err == nil {
			event = &MonitorAdded{
				Id:          parts[0],
				Name:        parts[1],
				Description: parts[2],
			}
		}
	case HyprEventTypes.Screencast:
		if parts, err = splitOnDetail(details, 2); err == nil {
			temp := &Screencast{}
			isValid := true

			switch parts[0] {
			case "0":
				temp.IsCast = false
			case "1":
				temp.IsCast = true
			default:
				isValid = false
			}

			switch parts[1] {
			case "0":
				temp.IsMonitorShare = true
				temp.IsWindowShare = false
			case "1":
				temp.IsMonitorShare = false
				temp.IsWindowShare = true
			default:
				isValid = false
			}

			if isValid {
				event = temp
			} else {
				err = ErrEventDetails
			}
		}
	case HyprEventTypes.ToggleGroup:
		if parts, err = splitOnDetail(details, 2); err == nil {
			temp := &ToggleGroup{WindowAddresses: parts[1:]}
			switch parts[1] {
			case "0":
				temp.IsDestroyed = true
			case "1":
				temp.IsDestroyed = false
			default:
				err = ErrEventDetails
			}
			event = temp
		}
	case HyprEventTypes.MoveIntoGroup:
		event = &MoveIntoGroup{WindowAddress: details}
	case HyprEventTypes.MoveOutOfGroup:
		event = &MoveOutOfGroup{WindowAddress: details}
	case HyprEventTypes.IgnoreGroupLock:
		switch details {
		case "0":
			event = &IgnoreGroupLock{Ignored: false}
		case "1":
			event = &IgnoreGroupLock{Ignored: true}
		default:
			err = ErrEventDetails
		}
	case HyprEventTypes.LockGroups:
		switch details {
		case "0":
			event = &LockGroups{IsLocked: false}
		case "1":
			event = &LockGroups{IsLocked: true}
		default:
			err = ErrEventDetails
		}
	case HyprEventTypes.ConfigreLoaded:
		event = &ConfigReloaded{}
	case HyprEventTypes.CloseLayer:
		event = &CloseLayer{Namespace: details}
	case HyprEventTypes.OpenLayer:
		event = &OpenLayer{Namespace: details}
	case HyprEventTypes.Submap:
		event = &Submap{Name: details}
	case HyprEventTypes.ActiveLayout:
		if parts, err = splitOnDetail(details, 2); err == nil {
			event = &ActiveLayout{
				KeyboardName: parts[0],
				Name:         parts[1],
			}
		}
	case HyprEventTypes.ActiveSpecial:
		if parts, err = splitOnDetail(details, 2); err == nil {
			event = &ActiveSpecial{
				WorkspaceName: parts[0],
				MonitorName:   parts[1],
			}
		}
	case HyprEventTypes.ActiveSpecialV2:
		if parts, err = splitOnDetail(details, 3); err == nil {
			event = &ActiveSpecialV2{
				WorkspaceId:   parts[0],
				WorkspaceName: parts[1],
				MonitorName:   parts[2],
			}
		}
	default:
		event = HyprEventUnhandled{
			Line: strings.TrimSpace(line),
		}
	}

	return event, err
}

func splitOnEvent(line string) (string, string, bool) {
	parts := strings.Split(strings.TrimSpace(line), ">>")

	if len(parts) != 2 {
		return "", "", false
	}

	return parts[0], parts[1], true
}

func splitOnDetail(line string, count int) ([]string, error) {
	var parts []string
	var err error

	parts = strings.Split(strings.TrimSpace(line), ",")

	if len(parts) != count {
		err = errors.New("invalid event details")
	}

	return parts, err
}

// /////////////////////////////////////////////////////////////////////////////
// Error
// /////////////////////////////////////////////////////////////////////////////

var _ HyprEvent = &HyprEventError{}

type HyprEventError struct {
	Text string
	Err  error
}

func (this HyprEventError) Error() string {
	return this.Err.Error()
}

func (_ HyprEventError) hyprEvent() {}

// /////////////////////////////////////////////////////////////////////////////
// Unhandled
// /////////////////////////////////////////////////////////////////////////////

var _ HyprEvent = &HyprEventUnhandled{}

type HyprEventUnhandled struct {
	Line string
}

func (_ HyprEventUnhandled) hyprEvent() {}

// /////////////////////////////////////////////////////////////////////////////
// Event Types
// /////////////////////////////////////////////////////////////////////////////

// https://wiki.hypr.land/IPC/
// vscode regex: 	^(\w+)\t.*\t([\w,\d/]+)\n
// vscode replace: 	$1 HyprEventType = "$1" // $2\n

type HyprEventType string

var HyprEventTypes = hyprEventTypes{
	Error:              "error",
	Unhandled:          "unhandled",
	Workspace:          "workspace",          // WORKSPACENAME
	WorkspaceV2:        "workspacev2",        // WORKSPACEID,WORKSPACENAME
	FocusedMon:         "focusedmon",         // MONNAME,WORKSPACENAME
	FocusedMonV2:       "focusedmonv2",       // MONNAME,WORKSPACEID
	ActiveWindow:       "activewindow",       // WINDOWCLASS,WINDOWTITLE
	ActiveWindowV2:     "activewindowv2",     // WINDOWADDRESS
	FullScreen:         "fullscreen",         // 0/1 (exit fullscreen / enter fullscreen)
	MonitorRemoved:     "monitorremoved",     // MONITORNAME
	MonitorRemovedV2:   "monitorremovedv2",   // MONITORID,MONITORNAME,MONITORDESCRIPTION
	MonitorAdded:       "monitoradded",       // MONITORNAME
	MonitorAddedV2:     "monitoraddedv2",     // MONITORID,MONITORNAME,MONITORDESCRIPTION
	CreateWorkspace:    "createworkspace",    // WORKSPACENAME
	CreateWorkspaceV2:  "createworkspacev2",  // WORKSPACEID,WORKSPACENAME
	DestroyWorkspace:   "destroyworkspace",   // WORKSPACENAME
	DestroyWorkspaceV2: "destroyworkspacev2", // WORKSPACEID,WORKSPACENAME
	MoveWorkspace:      "moveworkspace",      // WORKSPACENAME,MONNAME
	MoveWorkspaceV2:    "moveworkspacev2",    // WORKSPACEID,WORKSPACENAME,MONNAME
	RenameWorkspace:    "renameworkspace",    // WORKSPACEID,NEWNAME
	ActiveSpecial:      "activespecial",      // WORKSPACENAME,MONNAME
	ActiveSpecialV2:    "activespecialv2",    // WORKSPACEID,WORKSPACENAME,MONNAME
	ActiveLayout:       "activelayout",       // KEYBOARDNAME,LAYOUTNAME
	OpenWindow:         "openwindow",         // WINDOWADDRESS,WORKSPACENAME,WINDOWCLASS,WINDOWTITLE
	CloseWindow:        "closewindow",        // WINDOWADDRESS
	MoveWindow:         "movewindow",         // WINDOWADDRESS,WORKSPACENAME
	MoveWindowV2:       "movewindowv2",       // WINDOWADDRESS,WORKSPACEID,WORKSPACENAME
	OpenLayer:          "openlayer",          // NAMESPACE
	CloseLayer:         "closelayer",         // NAMESPACE
	Submap:             "submap",             // SUBMAPNAME
	ChangeFloatingMode: "changefloatingmode", // WINDOWADDRESS,FLOATING
	Urgent:             "urgent",             // WINDOWADDRESS
	Screencast:         "screencast",         // STATE,OWNER
	WindowTitle:        "windowtitle",        // WINDOWADDRESS
	WindowTitleV2:      "windowtitlev2",      // WINDOWADDRESS,WINDOWTITLE
	ToggleGroup:        "togglegroup",        //	0/1,WINDOWADDRESS(ES)
	MoveIntoGroup:      "moveintogroup",      // WINDOWADDRESS
	MoveOutOfGroup:     "moveoutofgroup",     // WINDOWADDRESS
	IgnoreGroupLock:    "ignoregrouplock",    // 0/1
	LockGroups:         "lockgroups",         // 0/1
	ConfigreLoaded:     "configreloaded",     // empty
	Pin:                "pin",                // WINDOWADDRESS,PINSTATE
	Minimized:          "minimized",          // WINDOWADDRESS,0/1
	Bell:               "bell",               // WINDOWADDRESS
}

type hyprEventTypes struct {
	Error              HyprEventType
	Unhandled          HyprEventType
	Workspace          HyprEventType // WORKSPACENAME
	WorkspaceV2        HyprEventType // WORKSPACEID,WORKSPACENAME
	FocusedMon         HyprEventType // MONNAME,WORKSPACENAME
	FocusedMonV2       HyprEventType // MONNAME,WORKSPACEID
	ActiveWindow       HyprEventType // WINDOWCLASS,WINDOWTITLE
	ActiveWindowV2     HyprEventType // WINDOWADDRESS
	FullScreen         HyprEventType // 0/1 (exit fullscreen / enter fullscreen)
	MonitorRemoved     HyprEventType // MONITORNAME
	MonitorRemovedV2   HyprEventType // MONITORID,MONITORNAME,MONITORDESCRIPTION
	MonitorAdded       HyprEventType // MONITORNAME
	MonitorAddedV2     HyprEventType // MONITORID,MONITORNAME,MONITORDESCRIPTION
	CreateWorkspace    HyprEventType // WORKSPACENAME
	CreateWorkspaceV2  HyprEventType // WORKSPACEID,WORKSPACENAME
	DestroyWorkspace   HyprEventType // WORKSPACENAME
	DestroyWorkspaceV2 HyprEventType // WORKSPACEID,WORKSPACENAME
	MoveWorkspace      HyprEventType // WORKSPACENAME,MONNAME
	MoveWorkspaceV2    HyprEventType // WORKSPACEID,WORKSPACENAME,MONNAME
	RenameWorkspace    HyprEventType // WORKSPACEID,NEWNAME
	ActiveSpecial      HyprEventType // WORKSPACENAME,MONNAME
	ActiveSpecialV2    HyprEventType // WORKSPACEID,WORKSPACENAME,MONNAME
	ActiveLayout       HyprEventType // KEYBOARDNAME,LAYOUTNAME
	OpenWindow         HyprEventType // WINDOWADDRESS,WORKSPACENAME,WINDOWCLASS,WINDOWTITLE
	CloseWindow        HyprEventType // WINDOWADDRESS
	MoveWindow         HyprEventType // WINDOWADDRESS,WORKSPACENAME
	MoveWindowV2       HyprEventType // WINDOWADDRESS,WORKSPACEID,WORKSPACENAME
	OpenLayer          HyprEventType // NAMESPACE
	CloseLayer         HyprEventType // NAMESPACE
	Submap             HyprEventType // SUBMAPNAME
	ChangeFloatingMode HyprEventType // WINDOWADDRESS,FLOATING
	Urgent             HyprEventType // WINDOWADDRESS
	Screencast         HyprEventType // STATE,OWNER
	WindowTitle        HyprEventType // WINDOWADDRESS
	WindowTitleV2      HyprEventType // WINDOWADDRESS,WINDOWTITLE
	ToggleGroup        HyprEventType //	0/1,WINDOWADDRESS(ES)
	MoveIntoGroup      HyprEventType // WINDOWADDRESS
	MoveOutOfGroup     HyprEventType // WINDOWADDRESS
	IgnoreGroupLock    HyprEventType // 0/1
	LockGroups         HyprEventType // 0/1
	ConfigreLoaded     HyprEventType // empty
	Pin                HyprEventType // WINDOWADDRESS,PINSTATE
	Minimized          HyprEventType // WINDOWADDRESS,0/1
	Bell               HyprEventType // WINDOWADDRESS
}

// /////////////////////////////////////////////////////////////////////////////
// Events
// /////////////////////////////////////////////////////////////////////////////

var _ HyprEvent = &Workspace{}

type Workspace struct {
	Name string
}

func (_ Workspace) hyprEvent() {}

// /////////////////////////////////////////////////////////////////////////////

var _ HyprEvent = &WorkspaceV2{}

type WorkspaceV2 struct {
	Id   string
	Name string
}

func (_ WorkspaceV2) hyprEvent() {}

// /////////////////////////////////////////////////////////////////////////////

var _ HyprEvent = &ActiveWindow{}

type ActiveWindow struct {
	Class model.HyprWindowId
	Title model.HyprWindowId
}

func (_ ActiveWindow) hyprEvent() {}

// /////////////////////////////////////////////////////////////////////////////

var _ HyprEvent = &ActiveWindowV2{}

type ActiveWindowV2 struct {
	Address model.HyprWindowId
}

func (_ ActiveWindowV2) hyprEvent() {}

// /////////////////////////////////////////////////////////////////////////////

var _ HyprEvent = &Screencast{}

type Screencast struct {
	IsCast         bool
	IsMonitorShare bool
	IsWindowShare  bool
}

func (_ Screencast) hyprEvent() {}

// /////////////////////////////////////////////////////////////////////////////

var _ HyprEvent = &Fullscreen{}

type Fullscreen struct {
	IsFullScreen bool
}

func (_ Fullscreen) hyprEvent() {}

// /////////////////////////////////////////////////////////////////////////////

var _ HyprEvent = &MonitorAdded{}

type MonitorAdded struct {
	Id          string
	Name        string
	Description string
}

func (_ MonitorAdded) hyprEvent() {}

// /////////////////////////////////////////////////////////////////////////////

var _ HyprEvent = &MonitorRemoved{}

type MonitorRemoved struct {
	Id          string
	Name        string
	Description string
}

func (_ MonitorRemoved) hyprEvent() {}

// /////////////////////////////////////////////////////////////////////////////

var _ HyprEvent = &MonitorFocused{}

type MonitorFocused struct {
	Name          string
	WorkspaceName string
}

func (_ MonitorFocused) hyprEvent() {}

// /////////////////////////////////////////////////////////////////////////////

var _ HyprEvent = &MonitorFocusedV2{}

type MonitorFocusedV2 struct {
	Name        string
	WorkspaceId string
}

func (_ MonitorFocusedV2) hyprEvent() {}

// /////////////////////////////////////////////////////////////////////////////

var _ HyprEvent = &WindowActive{}

type WindowActive struct {
	Address string
	Title   string
	Class   string
}

func (_ WindowActive) hyprEvent() {}

func (this WindowActive) IsReady() bool {
	return this.Address != "" && this.Title != "" && this.Class != ""
}

// /////////////////////////////////////////////////////////////////////////////

var _ HyprEvent = &OpenWindow{}

type OpenWindow struct {
	Address       model.HyprWindowId
	WorkspaceName string
	Class         model.HyprWindowId
	Title         model.HyprWindowId
}

func (_ OpenWindow) hyprEvent() {}

// /////////////////////////////////////////////////////////////////////////////

var _ HyprEvent = &CloseWindow{}

type CloseWindow struct {
	Address string
}

func (_ CloseWindow) hyprEvent() {}

// /////////////////////////////////////////////////////////////////////////////

var _ HyprEvent = &WindowUrgent{}

type WindowUrgent struct {
	Address string
}

func (_ WindowUrgent) hyprEvent() {}

// /////////////////////////////////////////////////////////////////////////////

var _ HyprEvent = &MoveWindow{}

type MoveWindow struct {
	Address       string
	WorkspaceId   string
	WorkspaceName string
}

func (_ MoveWindow) hyprEvent() {}

// /////////////////////////////////////////////////////////////////////////////

var _ HyprEvent = &ChangeFloatingMode{}

type ChangeFloatingMode struct {
	Address    string
	IsFloating bool
}

func (_ ChangeFloatingMode) hyprEvent() {}

// /////////////////////////////////////////////////////////////////////////////

var _ HyprEvent = &WindowTitle{}

type WindowTitle struct {
	Address string
	Title   string
}

func (_ WindowTitle) hyprEvent() {}

// /////////////////////////////////////////////////////////////////////////////

var _ HyprEvent = &Pin{}

type Pin struct {
	Address  string
	IsPinned bool
}

func (_ Pin) hyprEvent() {}

// /////////////////////////////////////////////////////////////////////////////

var _ HyprEvent = &Minimized{}

type Minimized struct {
	Address     string
	IsMinimized bool
}

func (_ Minimized) hyprEvent() {}

// /////////////////////////////////////////////////////////////////////////////

var _ HyprEvent = &Bell{}

type Bell struct {
	Address string
}

func (_ Bell) hyprEvent() {}

// /////////////////////////////////////////////////////////////////////////////

var _ HyprEvent = &ToggleGroup{}

type ToggleGroup struct {
	IsDestroyed     bool
	WindowAddresses []string
}

func (_ ToggleGroup) hyprEvent() {}

// /////////////////////////////////////////////////////////////////////////////

var _ HyprEvent = &MoveIntoGroup{}

type MoveIntoGroup struct {
	WindowAddress string
}

func (_ MoveIntoGroup) hyprEvent() {}

// /////////////////////////////////////////////////////////////////////////////

var _ HyprEvent = &MoveOutOfGroup{}

type MoveOutOfGroup struct {
	WindowAddress string
}

func (_ MoveOutOfGroup) hyprEvent() {}

// /////////////////////////////////////////////////////////////////////////////

var _ HyprEvent = &IgnoreGroupLock{}

type IgnoreGroupLock struct {
	Ignored bool
}

func (_ IgnoreGroupLock) hyprEvent() {}

// /////////////////////////////////////////////////////////////////////////////

var _ HyprEvent = &LockGroups{}

type LockGroups struct {
	IsLocked bool
}

func (_ LockGroups) hyprEvent() {}

// /////////////////////////////////////////////////////////////////////////////

var _ HyprEvent = &ConfigReloaded{}

type ConfigReloaded struct {
}

func (_ ConfigReloaded) hyprEvent() {}

// /////////////////////////////////////////////////////////////////////////////

var _ HyprEvent = &MoveWorkspace{}

type MoveWorkspace struct {
	Name        string
	MonitorName string
}

func (_ MoveWorkspace) hyprEvent() {}

// /////////////////////////////////////////////////////////////////////////////

var _ HyprEvent = &MoveWorkspaceV2{}

type MoveWorkspaceV2 struct {
	Id          string
	Name        string
	MonitorName string
}

func (_ MoveWorkspaceV2) hyprEvent() {}

// /////////////////////////////////////////////////////////////////////////////

var _ HyprEvent = &RenameWorkspace{}

type RenameWorkspace struct {
	Id      string
	NewName string
}

func (_ RenameWorkspace) hyprEvent() {}

// /////////////////////////////////////////////////////////////////////////////

var _ HyprEvent = &CreateWorkspace{}

type CreateWorkspace struct {
	Name string
}

func (_ CreateWorkspace) hyprEvent() {}

// /////////////////////////////////////////////////////////////////////////////

var _ HyprEvent = &CreateWorkspaceV2{}

type CreateWorkspaceV2 struct {
	Id   string
	Name string
}

func (_ CreateWorkspaceV2) hyprEvent() {}

// /////////////////////////////////////////////////////////////////////////////

var _ HyprEvent = &DestroyWorkspace{}

type DestroyWorkspace struct {
	Name string
}

func (_ DestroyWorkspace) hyprEvent() {}

// /////////////////////////////////////////////////////////////////////////////

var _ HyprEvent = &DestroyWorkspaceV2{}

type DestroyWorkspaceV2 struct {
	Id   string
	Name string
}

func (_ DestroyWorkspaceV2) hyprEvent() {}

// /////////////////////////////////////////////////////////////////////////////

var _ HyprEvent = &OpenLayer{}

type OpenLayer struct {
	Namespace string
}

func (_ OpenLayer) hyprEvent() {}

// /////////////////////////////////////////////////////////////////////////////

var _ HyprEvent = &CloseLayer{}

type CloseLayer struct {
	Namespace string
}

func (_ CloseLayer) hyprEvent() {}

// /////////////////////////////////////////////////////////////////////////////

var _ HyprEvent = &Submap{}

type Submap struct {
	Name string
}

func (_ Submap) hyprEvent() {}

// /////////////////////////////////////////////////////////////////////////////

var _ HyprEvent = &ActiveSpecial{}

type ActiveSpecial struct {
	WorkspaceName string
	MonitorName   string
}

func (_ ActiveSpecial) hyprEvent() {}

// /////////////////////////////////////////////////////////////////////////////

var _ HyprEvent = &ActiveSpecialV2{}

type ActiveSpecialV2 struct {
	WorkspaceId   string
	WorkspaceName string
	MonitorName   string
}

func (_ ActiveSpecialV2) hyprEvent() {}

// /////////////////////////////////////////////////////////////////////////////

var _ HyprEvent = &ActiveLayout{}

type ActiveLayout struct {
	KeyboardName string
	Name         string
}

func (_ ActiveLayout) hyprEvent() {}

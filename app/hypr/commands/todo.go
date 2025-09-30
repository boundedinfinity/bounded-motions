package commands

import (
	"bufio"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// /////////////////////////////////////////////////////////////////////////////
// Relative Direction
// /////////////////////////////////////////////////////////////////////////////

type RDirection string

func (this RDirection) String() string {
	return string(this)
}

var RDirections = rdirections{
	Abs:  "",
	Up:   "+",
	Down: "-",
}

type rdirections struct {
	Abs  RDirection
	Up   RDirection
	Down RDirection
}

// /////////////////////////////////////////////////////////////////////////////
// Monitor Result
// /////////////////////////////////////////////////////////////////////////////

type SizeInPixel struct {
	Width  int
	Height int
}

type MonitorsResult struct {
	Monitors []MonitorResult
}

type MonitorResult struct {
	Name             string
	Id               int
	Resolution       string
	RefreshRate      string
	Position         string
	Description      string
	Make             string
	Model            string
	PhysicalSizeMM   SizeInPixel
	Serial           string
	ActiveWorkspace  string
	SpecialWorkspace string
	Reserved         string
	Scale            string
	Transform        bool
	Focused          bool
	DpmsStatus       string
	Vrr              bool
	Solitary         string
	SolitaryBlocked  string
	ActivelyTearing  bool
	TearingBlocked   string
	DirectScanoutTo  string
	DirectScanoutBlk string
	Disabled         string
	CurrentFormat    string
	MirrorOf         string
	AvailableModes   []string
}

var monitorHeaderRegex = regexp.MustCompile(`Monitor\s+([^\s]+)\s+\(ID\s+(\d+)\):`)

// ParseMonitor parses the block of text into a Monitor struct
func ParseMonitor(input string) (*MonitorResult, error) {
	monitor := &MonitorResult{}

	if m := monitorHeaderRegex.FindStringSubmatch(input); m != nil {
		monitor.Name = m[1]
		monitor.Id = atoiSafe(m[2])
	}

	lines := strings.Split(input, "\n")

	for _, line := range lines {
		line = strings.TrimSpace(line)

		switch {
		case strings.HasPrefix(line, "description:"):
			monitor.Description = strings.TrimPrefix(line, "description: ")
		case strings.HasPrefix(line, "make:"):
			monitor.Make = strings.TrimPrefix(line, "make: ")
		case strings.HasPrefix(line, "model:"):
			monitor.Model = strings.TrimPrefix(line, "model: ")
		case strings.HasPrefix(line, "physical size (mm):"):
			val := strings.TrimPrefix(line, "physical size (mm): ")
			parts := strings.Split(val, "x")
			if len(parts) == 2 {
				w, _ := strconv.Atoi(parts[0])
				h, _ := strconv.Atoi(parts[1])
				monitor.PhysicalSizeMM = SizeInPixel{Width: w, Height: h}
			}
		case strings.HasPrefix(line, "serial:"):
			monitor.Serial = strings.TrimPrefix(line, "serial: ")
		case strings.HasPrefix(line, "active workspace:"):
			monitor.ActiveWorkspace = strings.TrimPrefix(line, "active workspace: ")
		case strings.HasPrefix(line, "special workspace:"):
			monitor.SpecialWorkspace = strings.TrimPrefix(line, "special workspace: ")
		case strings.HasPrefix(line, "reserved:"):
			monitor.Reserved = strings.TrimPrefix(line, "reserved: ")
		case strings.HasPrefix(line, "scale:"):
			monitor.Scale = strings.TrimPrefix(line, "scale: ")
		case strings.HasPrefix(line, "transform:"):
			val := strings.TrimPrefix(line, "transform: ")
			num, _ := strconv.Atoi(val)
			monitor.Transform = num != 0
		case strings.HasPrefix(line, "focused:"):
			val := strings.TrimPrefix(line, "focused: ")
			monitor.Focused = val == "yes"
		case strings.HasPrefix(line, "dpmsStatus:"):
			monitor.DpmsStatus = strings.TrimPrefix(line, "dpmsStatus: ")
		case strings.HasPrefix(line, "vrr:"):
			val := strings.TrimPrefix(line, "vrr: ")
			monitor.Vrr = val == "true"
		case strings.HasPrefix(line, "solitary:"):
			monitor.Solitary = strings.TrimPrefix(line, "solitary: ")
		case strings.HasPrefix(line, "solitaryBlockedBy:"):
			monitor.SolitaryBlocked = strings.TrimPrefix(line, "solitaryBlockedBy: ")
		case strings.HasPrefix(line, "activelyTearing:"):
			val := strings.TrimPrefix(line, "activelyTearing: ")
			monitor.ActivelyTearing = val == "true"
		case strings.HasPrefix(line, "tearingBlockedBy:"):
			monitor.TearingBlocked = strings.TrimPrefix(line, "tearingBlockedBy: ")
		case strings.HasPrefix(line, "directScanoutTo:"):
			monitor.DirectScanoutTo = strings.TrimPrefix(line, "directScanoutTo: ")
		case strings.HasPrefix(line, "directScanoutBlockedBy:"):
			monitor.DirectScanoutBlk = strings.TrimPrefix(line, "directScanoutBlockedBy: ")
		case strings.HasPrefix(line, "disabled:"):
			monitor.Disabled = strings.TrimPrefix(line, "disabled: ")
		case strings.HasPrefix(line, "currentFormat:"):
			monitor.CurrentFormat = strings.TrimPrefix(line, "currentFormat: ")
		case strings.HasPrefix(line, "mirrorOf:"):
			monitor.MirrorOf = strings.TrimPrefix(line, "mirrorOf: ")
		case strings.HasPrefix(line, "availableModes:"):
			modes := strings.TrimPrefix(line, "availableModes: ")
			monitor.AvailableModes = strings.Fields(modes)
		default:
			// Resolution line
			reRes := regexp.MustCompile(`^(\d+x\d+)@([\d.]+)\s+at\s+(\S+)$`)
			if m := reRes.FindStringSubmatch(line); m != nil {
				monitor.Resolution = m[1]
				monitor.RefreshRate = m[2]
				monitor.Position = m[3]
			}
		}
	}

	return monitor, nil
}

func atoiSafe(s string) int {
	var v int
	fmt.Sscanf(s, "%d", &v)
	return v
}

// /////////////////////////////////////////////////////////////////////////////
// Workspace Result
// /////////////////////////////////////////////////////////////////////////////

type WorkspacesResult struct {
	Monitors []WorkspaceResult
}

type WorkspaceResult struct {
	Name            string
	MonitorName     string
	MonitorId       string
	Windows         string
	HasFullScreen   string
	LastWindow      string
	LastWindowTitle string
	IsPersistent    string
}

// /////////////////////////////////////////////////////////////////////////////
// Monitor Result
// /////////////////////////////////////////////////////////////////////////////

type WindowsResult struct {
	Monitors []WindowResult
}

type WindowResult struct {
	Address          string
	Title            string
	Mapped           string
	Hidden           string
	At               string
	Size             string
	Workspace        string
	Floating         string
	Pseudo           string
	Monitor          string
	Class            string
	InitialClass     string
	InitialTitle     string
	Pid              string
	Xwayland         string
	Pinned           string
	Fullscreen       string
	FullscreenClient string
	Grouped          string
	Tags             string
	Swallowing       string
	FocusHistoryId   string
	InhibitingIdle   string
	XdgTag           string
	XdgDescription   string
}

// /////////////////////////////////////////////////////////////////////////////
// Devices Result
// /////////////////////////////////////////////////////////////////////////////

type Mouse struct {
	Address      string
	Name         string
	DefaultSpeed float64
	ScrollFactor float64
}

type Keyboard struct {
	Address         string
	Name            string
	Rules           string
	Model           string
	Layout          string
	Variant         string
	Options         string
	ActiveLayoutIdx int
	ActiveKeymap    string
	CapsLock        bool
	NumLock         bool
	Main            bool
}

type Devices2 struct {
	Mice      []Mouse
	Keyboards []Keyboard
	Tablets   int
	Touch     int
	Switches  int
}

// ---- Mouse Parser ----
func parseMice(scanner *bufio.Scanner) []Mouse {
	var mice []Mouse
	var current *Mouse

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		// stop when we hit Keyboards section
		if strings.HasPrefix(line, "Keyboards:") {
			// push last mouse before returning
			if current != nil {
				mice = append(mice, *current)
			}
			// push line back for higher-level parser
			scanner = pushBack(scanner, line)
			break
		}

		switch {
		case strings.HasPrefix(line, "Mouse at "):
			if current != nil {
				mice = append(mice, *current)
			}
			addr := strings.TrimPrefix(line, "Mouse at ")
			current = &Mouse{Address: strings.TrimSuffix(addr, ":")}
		case current != nil && !strings.Contains(line, ":"):
			current.Name = line
		case current != nil && strings.HasPrefix(line, "default speed:"):
			val, _ := strconv.ParseFloat(strings.TrimPrefix(line, "default speed:"), 64)
			current.DefaultSpeed = val
		case current != nil && strings.HasPrefix(line, "scroll factor:"):
			val, _ := strconv.ParseFloat(strings.TrimPrefix(line, "scroll factor:"), 64)
			current.ScrollFactor = val
		}
	}

	if current != nil {
		mice = append(mice, *current)
	}
	return mice
}

// ---- Keyboard Parser ----
func parseKeyboards(scanner *bufio.Scanner) []Keyboard {
	var keyboards []Keyboard
	var current *Keyboard

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		// stop when we hit Tablets section
		if strings.HasPrefix(line, "Tablets:") {
			if current != nil {
				keyboards = append(keyboards, *current)
			}
			scanner = pushBack(scanner, line)
			break
		}

		switch {
		case strings.HasPrefix(line, "Keyboard at "):
			if current != nil {
				keyboards = append(keyboards, *current)
			}
			addr := strings.TrimPrefix(line, "Keyboard at ")
			current = &Keyboard{Address: strings.TrimSuffix(addr, ":")}
		case current != nil && !strings.Contains(line, ":"):
			current.Name = line
		case current != nil && strings.HasPrefix(line, "rules:"):
			rules := strings.TrimPrefix(line, "rules:")
			parts := strings.Split(rules, ",")
			for _, part := range parts {
				part = strings.TrimSpace(part)
				if strings.HasPrefix(part, "r ") {
					current.Rules = strings.Trim(part[2:], `" `)
				} else if strings.HasPrefix(part, "m ") {
					current.Model = strings.Trim(part[2:], `" `)
				} else if strings.HasPrefix(part, "l ") {
					current.Layout = strings.Trim(part[2:], `" `)
				} else if strings.HasPrefix(part, "v ") {
					current.Variant = strings.Trim(part[2:], `" `)
				} else if strings.HasPrefix(part, "o ") {
					current.Options = strings.Trim(part[2:], `" `)
				}
			}
		case current != nil && strings.HasPrefix(line, "active layout index:"):
			val, _ := strconv.Atoi(strings.TrimPrefix(line, "active layout index:"))
			current.ActiveLayoutIdx = val
		case current != nil && strings.HasPrefix(line, "active keymap:"):
			current.ActiveKeymap = strings.TrimPrefix(line, "active keymap:")
		case current != nil && strings.HasPrefix(line, "capsLock:"):
			current.CapsLock = parseBool(strings.TrimPrefix(line, "capsLock:"))
		case current != nil && strings.HasPrefix(line, "numLock:"):
			current.NumLock = parseBool(strings.TrimPrefix(line, "numLock:"))
		case current != nil && strings.HasPrefix(line, "main:"):
			current.Main = parseBool(strings.TrimPrefix(line, "main:"))
		}
	}

	if current != nil {
		keyboards = append(keyboards, *current)
	}
	return keyboards
}

// ---- High-level Devices Parser ----
func ParseDevices(input string) Devices2 {
	scanner := bufio.NewScanner(strings.NewReader(input))
	var devices Devices2

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		switch {
		case strings.HasPrefix(line, "mice:"):
			devices.Mice = parseMice(scanner)
		case strings.HasPrefix(line, "Keyboards:"):
			devices.Keyboards = parseKeyboards(scanner)
		case strings.HasPrefix(line, "Tablets:"):
			val, _ := strconv.Atoi(strings.TrimPrefix(line, "Tablets:"))
			devices.Tablets = val
		case strings.HasPrefix(line, "Touch:"):
			val, _ := strconv.Atoi(strings.TrimPrefix(line, "Touch:"))
			devices.Touch = val
		case strings.HasPrefix(line, "Switches:"):
			val, _ := strconv.Atoi(strings.TrimPrefix(line, "Switches:"))
			devices.Switches = val
		}
	}

	return devices
}

// --- Helper: Push scanner line back ---
func pushBack(scanner *bufio.Scanner, line string) *bufio.Scanner {
	// In practice you'd need a custom scanner wrapper to really push back,
	// here we simply fake it by prepending to the input in ParseDevices
	// or restructuring parsing to avoid needing pushback.
	// For now, it's a placeholder.
	return scanner
}

func parseBool(s string) bool {
	var ok bool
	switch strings.TrimSpace(s) {
	case "yes", "y", "1", "true", "t":
		ok = true
	}
	return ok
}

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

type MonitorResult2 struct {
	Name string
	Id   int

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

package commands

import "strings"

type HyperCommandResult interface {
	String() string
	hyperCommandResult() // discriminator
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

var _ HyperCommandResult = &MonitorsResult{}

type MonitorsResult []*MonitorResult

func (this MonitorsResult) String() string {
	var result []string

	for _, monitor := range this {
		result = append(result, monitor.Name)
	}

	return strings.Join(result, ", ")
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

func parseMonitors(input string) (MonitorsResult, error) {
	var result MonitorsResult
	var err error
	var monitor *MonitorResult

	token := func(v string, sep string) (string, string) {
		i := strings.Index(v, sep)
		if i >= 0 {
			a := strings.TrimSpace(v[0:i])
			b := strings.TrimSpace(v[i+len(sep):])
			return a, b
		}
		return "", v
	}

	for _, line := range strings.Split(input, "\n") {
		switch {
		case line == "":
			continue
		case strings.HasPrefix(line, "Monitor"):
			monitor = &MonitorResult{}
			result = append(result, monitor)
			_, line := token(line, "Monitor")
			monitor.Name, line = token(line, "(ID")
			monitor.Id, _ = token(line, ")")
		case strings.Contains(line, "x") && strings.Contains(line, " at "):
			monitor.Resolution.Width, line = token(line, "x")
			monitor.Resolution.Height, line = token(line, "@")
			monitor.Resolution.RefreshRate, line = token(line, "at")
			monitor.Position.X, monitor.Position.Y = token(line, "x")
		case strings.HasPrefix(line, "description:"):
			_, monitor.Description = token(line, "description:")
		case strings.HasPrefix(line, "make:"):
			_, monitor.Make = token(line, "make:")
		case strings.HasPrefix(line, "model:"):
			_, monitor.Model = token(line, "model:")
		case strings.HasPrefix(line, "serial:"):
			_, monitor.Serial = token(line, "serial:")
		case strings.HasPrefix(line, "physical size (mm): "):
			_, line = token(line, "physical size (mm): ")
			monitor.Size.Height, monitor.Size.Width = token(line, "x")
		case strings.HasPrefix(line, "active workspace: "):
			_, line = token(line, "active workspace: ")
			monitor.ActiveWorkspace, _ = token(line, "(")
		case strings.HasPrefix(line, "special workspace: "):
			_, line = token(line, "special workspace: ")
			monitor.SpecialWorkspace, _ = token(line, "(")
		}
	}

	return result, err

}

// /////////////////////////////////////////////////////////////////////////////

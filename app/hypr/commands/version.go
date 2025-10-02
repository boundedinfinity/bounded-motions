package commands

import (
	"fmt"
	"regexp"
)

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
	return fmt.Sprintf("version [%s]", this.Version)
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

package commands

import "fmt"

// /////////////////////////////////////////////////////////////////////////////
// https://wiki.hypr.land/Configuring/Dispatchers/#list-of-dispatchers

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

var _ HyperCommandResult = &ExecResult{}

type ExecResult struct {
}

func (_ ExecResult) String() string {
	return "exec []"
}

func (_ ExecResult) hyperCommandResult() {}

func parseExec(input string, i int) (*ExecResult, error) {
	var result ExecResult
	var err error

	return &result, err
}

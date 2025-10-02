package commands

import "fmt"

// /////////////////////////////////////////////////////////////////////////////
// https://wiki.hypr.land/Configuring/Dispatchers/#list-of-dispatchers

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

var _ HyperCommandResult = &ExecrResult{}

type ExecrResult struct {
}

func (_ ExecrResult) String() string {
	return "execr []"
}

func (_ ExecrResult) hyperCommandResult() {}

func parserExec(input string, i int) (*ExecrResult, error) {
	var result ExecrResult
	var err error

	return &result, err
}

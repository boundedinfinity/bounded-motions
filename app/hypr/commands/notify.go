package commands

import (
	"fmt"
	"go-motions/hypr/model"
	"time"
)

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

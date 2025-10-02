package commands

import (
	"fmt"
	"strconv"
	"strings"
)

type HyperCommandResult interface {
	String() string
	hyperCommandResult() // discriminator
}

func getSubmatch(subs []string, index int) string {
	var found string

	if len(subs) > index {
		found = subs[index]
	}

	return strings.TrimSpace(found)
}

func getSubmatchInt(subs []string, index int) int {
	found, _ := strconv.Atoi(getSubmatch(subs, index))
	return found
}

func getSubmatchBool(subs []string, index int) bool {
	var found bool
	text := getSubmatch(subs, index)

	switch text {
	case "1", "t", "true", "y", "yes":
		found = true
	case "0", "f", "false", "n", "no":
		found = false
	default:
		panic(fmt.Errorf("invalid bool value: %s", text))
	}

	return found
}

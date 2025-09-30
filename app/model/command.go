package model

type MotionCommand struct {
	Name string
	Keys []string
}

func (this MotionCommand) String() string {
	return this.Name
}

// /////////////////////////////////////////////////////////////////////////////

var MotionCommands = motionCommands{}

type motionCommands struct {
	Commands []MotionCommand
}

func (this motionCommands) Names() []string {
	var list []string

	for _, c := range this.Commands {
		list = append(list, c.String())
	}

	return list
}

func newMotionCommand(name string, keys ...string) {
	MotionCommands.Commands = append(MotionCommands.Commands, MotionCommand{
		Name: name,
		Keys: keys,
	})
}

func init() {
	// Window
	newMotionCommand("window.focus.up", "w", "f", "i")
	newMotionCommand("window.focus.down", "w", "f", "k")
	newMotionCommand("window.focus.left", "w", "f", "j")
	newMotionCommand("window.focus.right", "w", "f", "l")
	newMotionCommand("window.focus.find", "w", "f", "f")

	newMotionCommand("window.move.up", "w", "m", "i")
	newMotionCommand("window.move.down", "w", "m", "k")
	newMotionCommand("window.move.left", "w", "m", "j")
	newMotionCommand("window.move.right", "w", "m", "l")
	newMotionCommand("window.move.cancel", "w", "m", "c")
	newMotionCommand("window.move.done", "w", "m", "<enter>")

	newMotionCommand("window.close", "w", "m", "i")

	// Browser

	newMotionCommand("browser.tab.left", "b", "t", "j")
	newMotionCommand("browser.tab.right", "b", "t", "l")
	newMotionCommand("browser.tab.close", "b", "t", "c")

	newMotionCommand("browser.tab.search", "b", "t", "f")

	// Editor

	newMotionCommand("editor.tab.left", "e", "t", "j")
	newMotionCommand("editor.tab.right", "e", "t", "l")

	newMotionCommand("editor.tab.close", "e", "t", "c")
	newMotionCommand("editor.tab.search", "e", "t", "f")

	newMotionCommand("editor.word.up", "e", "w", "i")
	newMotionCommand("editor.word.down", "e", "w", "k")
	newMotionCommand("editor.word.left", "e", "w", "j")
	newMotionCommand("editor.word.right", "e", "w", "l")

	newMotionCommand("editor.navigate.back", "e", "n", "j")
	newMotionCommand("editor.navigate.forward", "e", "n", "l")
	newMotionCommand("editor.navigate.end", "e", "n", "e")
}

package motions

import (
	"fmt"
	"strings"

	"github.com/gdamore/tcell/v2"
)

type MotionEvent struct {
	Key *tcell.EventKey
}

type MotionResult struct {
	Path    []string
	Options []string
	Message string
	Error   error
}

// //////////////////////////////////////////////////////////////////////
// Motions
// //////////////////////////////////////////////////////////////////////

func Create() *Motions {
	root := new(nil, "Motions", "Motions", nil)

	motions := &Motions{
		root:      root,
		current:   root,
		keyColor:  "[green]",
		textColor: "[grey]",
		nameColor: "[white]",
	}

	line := new(motions.root, "Line", "l", nil)
	new(line, "↑", "i", nil)
	new(line, "Down", "k", nil)
	new(line, "Left", "j", nil)
	new(line, "Right", "r", nil)
	back(line)
	motions.toRoot(line)

	word := new(motions.root, "Word", "w", nil)
	new(word, "Up", "i", nil)
	new(word, "Down", "k", nil)
	new(word, "Left", "j", nil)
	new(word, "Right", "r", nil)
	back(word)
	motions.toRoot(word)

	return motions
}

type Motions struct {
	root      *Motion
	current   *Motion
	keyColor  string
	textColor string
	nameColor string
}

func (this Motions) toRoot(parent *Motion) *Motion {
	apply := func(me MotionEvent) (*Motion, error) { return this.root, nil }
	return new(parent, "Motions", "esc", apply)
}

func (this Motions) path() []string {
	names := []string{}

	for motion := this.current; motion.Parent != nil; motion = motion.Parent {
		name := fmt.Sprintf(`%s%s%s`, this.nameColor, motion.Name, this.textColor)
		names = append([]string{name}, names...)
	}

	names = append([]string{"[grey]Motions"}, names...)

	return names
}

func (this Motions) options() []string {
	options := []string{}

	for _, motion := range this.current.Children {
		text := fmt.Sprintf(`%s%s%s : %s`, this.keyColor, motion.Key, this.textColor, motion.Name)
		options = append(options, text)
	}

	return options
}

func (this *Motions) Initial() MotionResult {
	result := MotionResult{
		Path:    this.path(),
		Options: this.options(),
	}
	return result
}

func (this *Motions) HandleEvent(event MotionEvent) MotionResult {
	var found *Motion
	result := MotionResult{}
	key := event.Key.Name()

	if strings.HasPrefix(key, "Rune") {
		key = strings.Replace(key, "Rune[", "", 1)
		key = strings.Replace(key, "]", "", 1)
		key = strings.ToLower(key)
	}

	for _, motion := range this.current.Children {
		if motion.Key == key {
			found = motion
			break
		}
	}

	if found == nil {
		result.Message = fmt.Sprintf("invalid key: %s", key)
		return result
	}

	if found.Apply != nil {
		new, err := found.Apply(event)

		if err != nil {
			result.Message = err.Error()
		}

		if new != nil {
			this.current = new
		} else {
			this.current = found
		}
	} else {
		this.current = found
	}

	result.Path = this.path()
	result.Options = this.options()

	return result
}

// //////////////////////////////////////////////////////////////////////
// Motion
// //////////////////////////////////////////////////////////////////////

func new(parent *Motion, name string, key string, apply func(MotionEvent) (*Motion, error)) *Motion {
	m := &Motion{
		Name:     name,
		Key:      key,
		Apply:    apply,
		Parent:   parent,
		Children: []*Motion{},
	}

	if parent != nil {
		parent.Children = append(parent.Children, m)
	}

	return m
}

func back(parent *Motion) *Motion {
	apply := func(me MotionEvent) (*Motion, error) {
		if parent != nil && parent.Parent != nil {
			return parent.Parent, nil
		}

		return nil, nil
	}

	return new(parent, "Back", "q", apply)
}

type Motion struct {
	Name     string
	Key      string
	Desc     []string
	Apply    func(MotionEvent) (*Motion, error)
	Parent   *Motion
	Children []*Motion
}

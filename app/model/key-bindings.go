package model

import (
	"slices"
	"strings"
)

//////////////////////////////////////////////////////////////////////////////////////
// Config
//////////////////////////////////////////////////////////////////////////////////////

type KeyBindingJson struct {
	Name     string           `json:"name"`
	Key      string           `json:"key"`
	Expanded bool             `json:"expanded"`
	Children []KeyBindingJson `json:"children"`
}

//////////////////////////////////////////////////////////////////////////////////////
// Binding
//////////////////////////////////////////////////////////////////////////////////////

type KeyBinding struct {
	Name     string
	Key      string
	Expanded bool
	Parent   *KeyBinding
	Children []*KeyBinding
}

func (this *KeyBinding) ToggleExpanded() {
	this.Expanded = !this.Expanded
}

func (this *KeyBinding) AddChild(binding *KeyBinding) {
	this.Children = append(this.Children, binding)
	binding.Parent = this
}

func (this *KeyBinding) NewChild(key, name string) *KeyBinding {
	new := &KeyBinding{
		Key:      key,
		Name:     name,
		Children: []*KeyBinding{},
	}

	this.AddChild(new)
	return new
}

func (this *KeyBinding) WalkDown(fn func(*KeyBinding) bool) bool {
	if fn == nil {
		return false
	}

	if !fn(this) {
		return false
	}

	for _, child := range this.Children {
		if !child.WalkDown(fn) {
			return false
		}
	}

	return true
}

func (this *KeyBinding) WalkUp(fn func(*KeyBinding) bool) bool {
	if fn == nil {
		return false
	}

	if !fn(this) {
		return false
	}

	if this.Parent != nil {
		if !this.Parent.WalkUp(fn) {
			return false
		}
	}

	return true
}

func (this *KeyBinding) Path() []*KeyBinding {
	found := []*KeyBinding{}

	this.WalkUp(func(kb *KeyBinding) bool {
		found = append(found, kb)
		return true
	})

	slices.Reverse(found)
	return found
}

func (this *KeyBinding) MatchChild(key string) (*KeyBinding, bool) {
	var found *KeyBinding
	var ok bool

	for _, child := range this.Children {
		if key == child.Key {
			found = child
			ok = true
			break
		}
	}

	return found, ok
}

func (this *KeyBinding) Find(keys ...string) (*KeyBinding, bool) {
	var found *KeyBinding

	matcher := func(key string) func(*KeyBinding) bool {
		lcKey := strings.ToLower(key)
		return func(binding *KeyBinding) bool {
			if lcKey == strings.ToLower(binding.Key) {
				found = binding
				return false
			}

			return true
		}
	}

	for _, key := range keys {
		if !this.WalkDown(matcher(key)) {
			break
		}
	}

	return found, found != nil
}

func (this *KeyBinding) Matches(keys ...string) bool {
	_, ok := this.Find(keys...)
	return ok
}

//////////////////////////////////////////////////////////////////////////////////////
// Utilities
//////////////////////////////////////////////////////////////////////////////////////

func LoadKeyBindings(config ConfigJson, binding *KeyBinding) error {
	result := walkConfig(config.Keybindings, nil)
	*binding = *result

	return nil
}

func walkConfig(c KeyBindingJson, p *KeyBinding) *KeyBinding {
	b := &KeyBinding{
		Name:     c.Name,
		Key:      c.Key,
		Expanded: c.Expanded,
	}

	if p != nil {
		b.Parent = p
		p.Children = append(p.Children, b)
	}

	if b.Expanded {
		b.WalkUp(func(kb *KeyBinding) bool {
			kb.Expanded = true
			return true
		})
	}

	for _, child := range c.Children {
		walkConfig(child, b)
	}

	return b
}

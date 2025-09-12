// https://github.com/charmbracelet/bubbletea/blob/master/examples/file-picker/main.go
package main

import (
	"errors"
	"go-motions/filepicker"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

func newFilePicker(config *config) filepickerModel {
	model := filepickerModel{
		config:     config,
		filepicker: filepicker.New(),
	}

	model.filepicker.AllowedTypes = config.AllowedExts
	model.filepicker.CurrentDirectory = config.SourceDir
	model.filepicker.ShowSize = true
	model.filepicker.ShowHidden = false
	model.filepicker.AutoHeight = false
	model.filepicker.Height = 50

	return model
}

type filepickerModel struct {
	config       *config
	filepicker   filepicker.Model
	SelectedFile string
	quitting     bool
	err          error
}

func (this filepickerModel) Init() tea.Cmd {
	return this.filepicker.Init()
}

func (this filepickerModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		}
	case clearErrorMsg:
		this.err = nil
	}

	this.filepicker, cmd = this.filepicker.Update(msg)
	cmds = append(cmds, cmd)

	// Did the user select a file?
	if didSelect, path := this.filepicker.DidSelectFile(msg); didSelect {
		// Get the path of the selected file.
		this.SelectedFile = path
	}

	// Did the user select a disabled file?
	// This is only necessary to display an error to the user.
	if didSelect, path := this.filepicker.DidSelectDisabledFile(msg); didSelect {
		// Let's clear the selectedFile and display an error.
		this.err = errors.New(path + " is not valid.")
		this.SelectedFile = ""
		return this, tea.Batch(cmd, clearErrorAfter(2*time.Second))
	}

	return this, tea.Batch(cmds...)
}

func (this filepickerModel) View() string {
	if this.quitting {
		return ""
	}

	return this.filepicker.View()
}

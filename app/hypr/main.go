package hypr

import (
	"context"
	"errors"
	"fmt"
	"go-motions/hypr/commands"
	"go-motions/hypr/events"
	"io/fs"
	"os"
	"path/filepath"
	"sync"
	"time"
)

func New() *Hypr {
	h := &Hypr{
		readDeadLine: 200 * time.Millisecond,
		errInCh:      make(chan error),
		ErrOutCh:     make(chan error),
		eventInCh:    make(chan string),
		EventOutCh:   make(chan events.HyprEvent),
		commandInCh:  make(chan commands.HyprCommand),
		commandRawCh: make(chan commandRaw),
		CommandOutCh: make(chan commands.HyperCommandResult),
		wg:           &sync.WaitGroup{},
	}

	h.cleanup = func() error {
		var errs []error

		if h.errInCh != nil {
			close(h.errInCh)
		}

		if h.ErrOutCh != nil {
			close(h.ErrOutCh)
		}

		if h.eventInCh != nil {
			close(h.eventInCh)
		}

		if h.EventOutCh != nil {
			close(h.EventOutCh)
		}

		if h.commandRawCh != nil {
			close(h.commandRawCh)
		}

		if h.commandInCh != nil {
			close(h.commandInCh)
		}

		if h.CommandOutCh != nil {
			close(h.CommandOutCh)
		}

		return errors.Join(errs...)
	}

	return h
}

type Hypr struct {
	readDeadLine time.Duration
	errInCh      chan error
	ErrOutCh     chan error
	eventInCh    chan string
	EventOutCh   chan events.HyprEvent
	commandInCh  chan commands.HyprCommand
	commandRawCh chan commandRaw
	CommandOutCh chan commands.HyperCommandResult
	wg           *sync.WaitGroup
	cleanup      func() error
}

func (this *Hypr) Start(ctx context.Context) error {
	err := this.initReadSocket(ctx)

	if err != nil {
		return err
	}
Loop:
	for {
		select {
		case <-ctx.Done():
			break Loop
		case err := <-this.errInCh:
			go func() { this.ErrOutCh <- err }()
		case raw := <-this.eventInCh:
			if event, err := events.Parse(raw); err != nil {
				go func() { this.errInCh <- err }()
			} else {
				go func() { this.EventOutCh <- event }()
			}
		case raw := <-this.commandRawCh:
			fmt.Print(raw.Text)
			if result, err := commands.ParseResult(raw.Command, raw.Text); err != nil {
				go func() { this.errInCh <- err }()
			} else {
				go func() { this.CommandOutCh <- result }()
			}
		case cmd := <-this.commandInCh:
			if raw, err := this.writeCommand(cmd); err != nil {
				go func() { this.errInCh <- err }()
			} else {
				//TODO: not sure why this is needed
				go func() { this.commandRawCh <- raw }()
			}
		}
	}

	this.wg.Wait()
	return nil
}

func (this *Hypr) Write(command commands.HyprCommand) {
	if this.commandInCh != nil {
		go func() { this.commandInCh <- command }()
	}
}

func (this *Hypr) Stop() error {
	if this.cleanup != nil {
		return this.cleanup()
	}

	return nil
}

// /////////////////////////////////////////////////////////////////////////////
// Utility functions
// /////////////////////////////////////////////////////////////////////////////

func (_ Hypr) getSocketPath(filename string) (string, error) {
	xdgRuntimeDir := os.Getenv("XDG_RUNTIME_DIR")

	if xdgRuntimeDir == "" {
		return "", errors.New("XDG_RUNTIME_DIR not found")
	}

	hyprlandInstanceSignature := os.Getenv("HYPRLAND_INSTANCE_SIGNATURE")

	if hyprlandInstanceSignature == "" {
		return "", errors.New("HYPRLAND_INSTANCE_SIGNATURE not found")
	}

	socketPath := filepath.Join(xdgRuntimeDir, "hypr", hyprlandInstanceSignature, filename)
	fileInfo, err := os.Stat(socketPath)

	if err != nil {
		return "", fmt.Errorf("%s: %v", socketPath, err)
	}

	if fileInfo.Mode().Type() != fs.ModeSocket {
		return "", fmt.Errorf("%s is not a socket", socketPath)
	}

	return socketPath, nil
}

package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"go-motions/hypr"
	"go-motions/hypr/commands"
)

func main() {
	wg := &sync.WaitGroup{}
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt, syscall.SIGTERM)
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		<-signalCh
		fmt.Println("Received shutdown signal")
		cancel()
	}()

	hserver := hypr.New()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				return
			case err := <-hserver.ErrOutCh:
				fmt.Printf("    Err: %+v\n", err)
			case event := <-hserver.EventOutCh:
				fmt.Printf("  Event: %+v\n", event)
			case result := <-hserver.CommandOutCh:
				fmt.Printf("Command: %+v\n", result)
			}
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		hserver.Write(&commands.Workspaces{})
		hserver.Write(&commands.Monitors{})
		hserver.Write(&commands.Version{})
		hserver.Write(&commands.ActiveWorkspace{})
		hserver.Write(&commands.ActiveWindow{})
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := hserver.Start(ctx); err != nil {
			panic(err)
		}
	}()

	wg.Wait()
}

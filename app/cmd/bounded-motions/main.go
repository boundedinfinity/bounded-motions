package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

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
		time.Sleep(2 * time.Second)
		hserver.Write(&commands.Monitors{})
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

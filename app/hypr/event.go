package hypr

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"io"
	"net"
	"os"
	"time"
)

func (this *Hypr) initReadSocket(ctx context.Context) error {
	path, err := this.getSocketPath(".socket2.sock")

	if err != nil {
		return err
	}

	socket, err := net.Dial("unix", path)

	if err != nil {
		return err
	}

	fmt.Printf("Hypr.initReadSocket(%s) listening \n", path)
	this.wg.Add(1)

	go func() {
		defer this.wg.Done()
		defer socket.Close()
		defer fmt.Printf("Hypr.initReadSocket(%s) shutdown complete\n", path)

		reader := bufio.NewReader(socket)

		for {
			select {
			case <-ctx.Done():
				fmt.Printf("Hypr.initReadSocket(%s) shutdown received\n", path)
				return
			default:
				deadline := time.Now().Add(this.readDeadLine)

				if err := socket.SetReadDeadline(deadline); err != nil {
					this.errInCh <- err
				}

				if line, err := reader.ReadString('\n'); err == nil {
					this.eventInCh <- line
				} else {
					if errors.Is(err, os.ErrDeadlineExceeded) {
						continue
					}

					this.errInCh <- err

					if errors.Is(err, io.EOF) {
						return
					}
				}
			}
		}
	}()

	return nil
}

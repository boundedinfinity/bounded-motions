package hypr

// https://wiki.hypr.land/Configuring/Using-hyprctl/

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"go-motions/hypr/commands"

	"io"
	"net"
	"time"
)

type commandRaw struct {
	Command commands.HyprCommand
	Text    string
}

func (this commandRaw) String() string {
	return fmt.Sprintf(`=======================================================================================
	%s
	=======================================================================================
	%s
	=======================================================================================
	`, this.Command, this.Text)
}

func (this *Hypr) writeCommand(cmd commands.HyprCommand) (commandRaw, error) {
	result := commandRaw{Command: cmd}

	path, err := this.getSocketPath(".socket.sock")

	if err != nil {
		return result, err
	}

	socket, err := net.Dial("unix", path)

	if err != nil {
		return result, err
	}

	defer socket.Close()

	reader := bufio.NewReader(socket)
	deadline := time.Now().Add(this.readDeadLine)

	var buf bytes.Buffer

	if err := socket.SetReadDeadline(deadline); err != nil {
		return result, err
	}

	if _, err := socket.Write([]byte(result.Command.String())); err != nil {
		return result, err
	}

	for {
		if line, err := reader.ReadString('\n'); err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return result, err
		} else {
			buf.WriteString(line)
		}
	}

	result.Text = buf.String()
	return result, nil
}

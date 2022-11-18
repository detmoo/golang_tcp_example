// Desc: TCP echo server and client
package pkg

import (
    "context"
    "fmt"
    "log"
    "net"
	"os"
	"os/signal"
	"syscall"
	"time"
)


type ServerClosedError struct{
    Reason string
    Msg string
}


func (m *ServerClosedError) Error() string {
	return fmt.Sprintf("reason %d: err %v", m.Reason, m.Msg)
}


func DeferCloseListener(listener net.Listener, timeout time.Duration, closureChannel chan<- error, parent context.Context) error {

    closer := func(reason string) ServerClosedError {
        defer listener.Close()
        err := &ServerClosedError{
            Reason: reason,
            Msg: "the listener was forcibly closed",
        }
        closureChannel <- err
        return err
    }

    ctx, stop := signal.NotifyContext(parent, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
    defer stop()

	select {
	case <-time.After(timeout):
        log.Println("listener time expired")
		return closer("timeout")
	case <-ctx.Done():
	    log.Println("cancel, interrupt or termination signal received")
		stop()
		return closer("interrupted")
	}
}

// Desc: TCP echo server and client
package pkg

import (
    "context"
    "errors"
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
    Err error
}


func (m *ServerClosedError) Error() string {
	return fmt.Sprintf("reason %d: err %v", m.Reason, m.Err)
}


func DeferCloseListener(listener net.Listener, timeout time.Duration, closureChannel chan<- error, parent context.Context) error {

    closer := func(reason string) error {
        defer listener.Close()
        err := &ServerClosedError{
            Reason: reason,
            Err: errors.New("the listener was forcibly closed"),
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

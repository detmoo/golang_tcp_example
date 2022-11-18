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
	return fmt.Sprintf("reason %d: err %v", r.Reason, r.Err)
}


func closeListenerE(reason string) error {
    defer listener.Close()
	err := ServerClosedError{
	    Reason: reason,
	    Err: errors.New("the listener was forcibly closed"),
	}
	closureChannel <- err
	return err
}


func DeferCloseListener(listener net.Listener, timeout time.Duration, closureChannel chan<- error, parent context.Context) error {
    ctx, stop := signal.NotifyContext(parent, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
    defer stop()

	select {
	case <-time.After(timeout):
        log.Println("listener time expired")
		return closeListenerE("timeout")
	case <-ctx.Done():
	    log.Println("cancel, interrupt or termination signal received")
		stop()
		return closeListenerE("interrupted")
	}
}

// Desc: TCP echo server and client
package pkg

import (
    "context"
    "fmt"
    "net"
	"os"
	"signal"
	"syscall"
	"time"
)


func DeferCloseListener(listener *net.Listener, timeout time.Duration, parent context.Context) string {
    ctx, stop := signal.NotifyContext(parent, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
    defer stop()

	select {
	case <-time.After(timeout):
	    defer listener.Close()
		fmt.Println("listener time expired")
		return "banana"
	case <-ctx.Done():
	    defer listener.Close()
		stop()
		fmt.Println("cancel, interrupt or termination signal received")
		return "grape"
	}
}

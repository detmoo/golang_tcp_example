package cmd

import (
    "context"
    "fmt"
	"log"
	"net"
    "time"

	"github.com/spf13/cobra"

	"bowdata.test.go_tcp_echo/pkg"
)

var duration string


func runServerCmd() *cobra.Command {
	cmd:= &cobra.Command{
		Use: "server",
		Short: "serves a TCP listener that sends a response to client connections",
		RunE: func(cmd *cobra.Command, args []string) error {
			listener, err := net.Listen("tcp", host+":"+port)
            if err != nil {
                log.Fatal(err)
                return err
            }
            // defer close listener
            timeout, _ := time.ParseDuration(duration)
            closureChannel := make(chan error, 1)
            ctx := context.Background()
            go pkg.DeferCloseListener(listener, timeout, closureChannel, ctx)

            // await connections
            for {
                conn, err := listener.Accept()
                if err != nil {
                    select {
                    case <-closureChannel:
                        close(closureChannel)
                        return err
                    default:
                        close(closureChannel)
                        log.Fatal(err)
                        fmt.Println(err)
                        return err
                    }
                }
                go pkg.HandleIncomingRequest(conn)
            }
		},
	}
	cmd.Flags().StringVar(&duration, "duration", "10s", "time.ParseDuration compatible string")
	return cmd
}

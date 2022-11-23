package cmd

import (
    "context"
    "fmt"
	"log"
	"net"
	"os"
    "time"

	"github.com/spf13/cobra"

	"bowdata.test.go_tcp_echo/pkg"
)

var duration, host, port string


func NewRootCmd() *cobra.Command {
	cmd:= &cobra.Command{
		Use: "tcp-echo-server",
		Short: "Echos the argument value in uppercase, and adds metadata to the response",
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
	cmd.Flags().StringVar(&host, "host", "localhost", "host IP address for the listener")
	cmd.Flags().StringVar(&port, "port", "9001", "host port for the listener")
	return cmd
}

func Execute() {
	rootCmd := NewRootCmd()
    log.Println("cmd.Execute(): rootCmd assigned")
	if err := rootCmd.Execute(); err != nil {
	os.Exit(1)
	}
}

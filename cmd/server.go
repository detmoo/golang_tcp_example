package cmd

import (
	"log"
	"net"
	"os"
    "timeout"

	"github.com/spf13/cobra"

	"bowdata.test.go_tcp_echo/pkg"
)

var host, port string
var timeout time.Duration = (10 * time.Second)

func NewRootCmd() *cobra.Command {
	cmd:= &cobra.Command{
		Use: "tcp-echo-server",
		Short: "Echos the argument value in uppercase, and adds metadata to the response",
		RunE: func(cmd *cobra.Command, args []string) error {
			listen, err := net.Listen("tcp", host+":"+port)
            if err != nil {
                log.Fatal(err)
                return err
            }
            // defer close listener
            defer listen.Close()

            // initialise user interrupt
            go pkg.DeferUserInterrupt(timeout)

            // an infinite loop of incoming connections
            for {
                conn, err := listen.Accept()
                if err != nil {
                    log.Fatal(err)
                    return err
                }
                go pkg.HandleIncomingRequest(conn)
            }
		},
	}
	cmd.Flags().StringVar(&host, "host", "localhost", "host IP address for the listener")
	cmd.Flags().StringVar(&port, "port", "9001", "host port for the listener")
	return cmd
}

func Execute() {
	rootCmd := NewRootCmd()
	if err := rootCmd.Execute(); err != nil {
	os.Exit(1)
	}
}

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


func newRootCmd() *cobra.Command {
	cmd:= &cobra.Command{
		Use: "tcp-server-client",
		Short: "Simple TCP Server/Client",
		Long: "Run TCP Server/Client via CLI e.g. for testing TCP connections",
	}
	cmd.AddCommand(
	    runServerCmd(),
	    runClientCmd(),
	)
	return cmd
}


func Execute() {
	rootCmd := newRootCmd()
	if err := rootCmd.Execute(); err != nil {
	os.Exit(1)
	}
}

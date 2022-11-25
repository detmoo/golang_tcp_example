package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var host, port string


func newRootCmd() *cobra.Command {
	cmd:= &cobra.Command{
		Use: "tcp-server-client",
		Short: "Simple TCP Server/Client",
		Long: "Run TCP Server/Client via CLI e.g. for testing TCP connections",
	}

	cmd.PersistentFlags().StringVar(&host, "host", "localhost", "attempts TCP connection via this IP address")
	cmd.PersistentFlags().StringVar(&port, "port", "9001", "attempts TCP connection via this host port")

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

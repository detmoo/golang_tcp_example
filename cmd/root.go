package cmd

import (
    "io"
	"os"

	"github.com/spf13/cobra"
)

var host, port string


func newRootCmd(out io.Writer) *cobra.Command {
	cmd:= &cobra.Command{
		Use: "tcp-server-client",
		Short: "Simple TCP Server/Client",
		Long: "Run TCP Server/Client via CLI e.g. for testing TCP connections",
	}

	cmd.PersistentFlags()
	cmd.Flags().StringVar(&host, "host", "localhost", "attempts TCP connection via this IP address")
	cmd.Flags().StringVar(&port, "port", "9001", "attempts TCP connection via this host port")

	cmd.AddCommand(
	    runServerCmd(),
	    runClientCmd(out),
	)
	return cmd
}


func Execute() {
	rootCmd := newRootCmd(os.Stdout)
	if err := rootCmd.Execute(); err != nil {
	os.Exit(1)
	}
}

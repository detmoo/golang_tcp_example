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

var message string


func runClientCmd() *cobra.Command {
	cmd:= &cobra.Command{
		Use: "server",
		Short: "serves a TCP listener that parses requests via the pkg.Message interface which is implemented herein",
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
		    requestContent := args[0]
            conn, err := net.Dial("tcp", host+":"+port)
            if err != nil {
                return err
            }
            result, err := pkg.MakeRequest(requestContent, conn)
            if err != nil {
                return err
            }
			return fmt.Fprintf(out, result)
		},
	}
	return cmd
}


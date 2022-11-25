package cmd

import (
    "encoding/json"
    "fmt"
    "io"
	"net"

	"github.com/spf13/cobra"

	"bowdata.test.go_tcp_echo/pkg"
)

var message string


func runClientCmd(out io.Writer) *cobra.Command {
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
            jsonStr, _ := json.Marshal(result)
			fmt.Fprintf(out, string(jsonStr))
			return nil
		},
	}
	return cmd
}
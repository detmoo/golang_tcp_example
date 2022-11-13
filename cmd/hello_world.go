package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"bowdata.test.go_tcp_echo/pkg"
)

var outputCase string

func NewRootCmd() *cobra.Command {
	cmd:= &cobra.Command{
		Use: "hello",
		Short: "Hello will echo the argument value, in caps or lowercase according to the flag",
		RunE: func(cmd *cobra.Command, args []string) error {
			res, err := pkg.HelloInput(args[0], outputCase)
			if err != nil {
				return err
			}
			fmt.Fprintf(cmd.OutOrStdout(), "%s", res)
			return nil
		},
	}
	cmd.Flags().StringVar(&outputCase, "output-case", "", "pass 'upper' or 'lower'")
	return cmd
}

func Execute() {
	rootCmd := NewRootCmd()
	if err := rootCmd.Execute(); err != nil {
	os.Exit(1)
	}
}

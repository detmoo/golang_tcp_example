package cmd

import (
	"bytes"
	"io"
	"net"
	"testing"

	"bowdata.test.go_tcp_echo/pkg"
)


HOST := "localhost"
PORT := "9001"

func TestEchoServer(t *testing.T) {
	rootCmd := NewRootCmd()
	b := bytes.NewBufferString("")
	rootCmd.SetOut(b)
	rootCmd.SetArgs([]string{"--host", HOST, "--port", PORT})
	rootCmd.Execute()
	out, err := io.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}
	connClient, _ = net.Dial("TCP", HOST+":"+PORT)
	fmt.Fprint(connClient, text+"\n")
}

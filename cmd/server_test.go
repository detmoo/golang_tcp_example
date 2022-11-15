package cmd

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"testing"

	"bowdata.test.go_tcp_echo/pkg"
)


var HOST string = "localhost"
var PORT string = "9001"
var REQUEST_CONTENT = "nature is great!"

func TestEchoServer(t *testing.T) {
	rootCmd := NewRootCmd()
	b := bytes.NewBufferString("")
	rootCmd.SetOut(b)
	rootCmd.SetArgs([]string{"--host", HOST, "--port", PORT})
	rootCmd.Execute()
	out, err := io.ReadAll(b)
	if err != nil {
	    fmt.Println("Read Buffer Error:", err)
		t.Fatal(err)
	}
	conn, err := net.Dial("tcp", HOST+":"+PORT)
	if err != nil {
	    fmt.Println("Dial Error:", err)
		t.Fatal(err)
	}
	fmt.Fprint(conn, REQUEST_CONTENT+"\n")
	fmt.Println("This is the out string:", out)
}

package cmd

import (
	"bytes"
	"io"
	"net"
	"os"
	"strings"
	"testing"
	"time"

	"bowdata.test.go_tcp_echo/pkg"
)


var clientTests = map[string]serverTestCase{
    "sends message and expects return": serverTestCase{
        host: "localhost",
        port: "9004",
        send: "mambo is great!",
        expected: "TCP listener received: mambo is great!",
    },
}


func TestEchoClient(t *testing.T) {
    for _, test := range clientTests {
        // setup a test listener for the client to connect to
        listener, err := net.Listen("tcp", test.host+":"+test.port)
        if err != nil {
            t.Error("TestEchoClient connection error:", err)
        }

        // background accepts first client connection
        go func() {
            conn, err := listener.Accept()
            if err != nil {
                t.Error("TestEchoClient listener error:", err)
            }
            pkg.HandleIncomingRequest(conn)
        }()

        // run the client command
    	rootCmd := newRootCmd(os.Stdout)
        b := bytes.NewBufferString("")
        rootCmd.SetOut(b)
        rootCmd.SetArgs([]string{"client", test.send,"--host", test.host, "--port", test.port})
        time.Sleep(2 * time.Second)  // to ensure the listener to ready to receive client connections
        rootCmd.Execute()

        // read the stdout
        out, err := io.ReadAll(b)
        if err != nil {
            t.Fatal(err)
        }
        if !strings.Contains(string(out), test.expected) {
            t.Fatalf("expected \"%s\" got \"%s\"", test.expected, string(out))
        }
    }
}

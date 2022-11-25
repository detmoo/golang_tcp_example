package cmd

import (
	"bytes"
	"log"
	"net"
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
    for testName, test := range clientTests {
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
    	rootCmd := newRootCmd()
        b := bytes.NewBufferString("")
        rootCmd.SetOut(b)
        rootCmd.SetArgs([]string{"client", test.send,"--host", test.host, "--port", test.port})
        time.Sleep(2 * time.Second)  // to ensure the listener to ready to receive client connections
        rootCmd.Execute()

		t.Errorf("xxxxxxxxx Failing Test String xxxxxxxxxx")

    }
}

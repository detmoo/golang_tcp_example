package cmd

import (
	"bytes"
	"log"
	"net"
	"os"
	"testing"
	"time"

	"bowdata.test.go_tcp_echo/pkg"
)


type serverTestCase struct {
        host string
        port string
        send string
        expected string
}

var serverTests = map[string]serverTestCase{
    "sends message and expects return": serverTestCase{
        host: "localhost",
        port: "9001",
        send: "mambo is great!",
        expected: "TCP listener received: mambo is great!",
    },
}

func TestEchoServer(t *testing.T) {
    for testName, test := range serverTests {
    	rootCmd := newRootCmd(os.Stdout)
        b := bytes.NewBufferString("")
        rootCmd.SetOut(b)
        rootCmd.SetArgs([]string{"server", "--host", test.host, "--port", test.port})
        go rootCmd.Execute()

        time.Sleep(2 * time.Second)  // to ensure the listener to ready to receive client connections
        log.Println("listener goroutine started. client dialling...")
        conn, err := net.Dial("tcp", test.host+":"+test.port)
        if err != nil {
            t.Error("TestEchoServer could not dial server:", err)
        }

        result, err := pkg.MakeRequest(test.send, conn)
        if err != nil {
            t.Error("TestEchoServer could not make request:", err)
        }

        if result.Content != test.expected{
			t.Errorf("TestEchoServer expected result: %s, but got: %s", test.expected, result.Content)
		}
    }
}

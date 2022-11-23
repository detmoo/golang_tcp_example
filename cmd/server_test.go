package cmd

import (
	"bytes"
	"log"
	"net"
	"testing"
	"time"

	"bowdata.test.go_tcp_echo/pkg"
)


var HOST string = "localhost"
var PORT string = "9001"
var REQUEST_CONTENT = "nature is great!"


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
        expected: "on2",
    },
}

func TestEchoServer(t *testing.T) {
    for testName, test := range serverTests {
    	rootCmd := NewRootCmd()
        b := bytes.NewBufferString("")
        rootCmd.SetOut(b)
        rootCmd.SetArgs([]string{"--host", HOST, "--port", PORT})
        go rootCmd.Execute()

        time.Sleep(3 * time.Second)  // to ensure the listener to ready to receive client connections
        log.Println("listener goroutine started. client dialling...")
        conn, err := net.Dial("tcp", HOST+":"+PORT)
        if err != nil {
            t.Error("client could not dial server:", err)
        }

        request := new(pkg.Message)
        request.Content = test.send
        request.Metadata = pkg.MetadataSchema{
            Timestamp: time.Now().Format("Monday, 02-Jan-06 15:04:05.123 MST"),
            Tag: testName,
            }
        log.Println("this is the request.Content:", request.Content)
        log.Println("this is the request.Metadata:", request.Metadata)
        result, err := pkg.MakeRequest(*request, conn)
        if err != nil {
            t.Error("test client could not make request:", err)
        }

        if result.Content != test.expected{
			t.Errorf("Expected result: %s, but got: %s", test.expected, result.Content)
		}
    }
}

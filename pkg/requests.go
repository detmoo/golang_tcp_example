// Desc: TCP echo server and client
package pkg

import (
    "encoding/json"
    "log"
    "fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)


func HandleIncomingRequest(conn net.Conn) error {
    // store incoming data
    buffer := make([]byte, 1024)
    _, err := conn.Read(buffer)
    if err != nil {
        log.Fatal(err)
        return err
    }
    receivedMsg := new(Message)
    receivedMsg.parse(buffer)
    // respond
    response := getResponse(receivedMsg)
    response.write(conn)

    // close conn
    conn.Close()
    return nil
}


type MetadataSchema struct {
    Timestamp string `json:"timestamp"`
    Tag string `json:"tag"`
}


type Message struct {
    Content string `json:"content"`
    Metadata MetadataSchema `json:"metadata"`
}


func (t Message) write(conn net.Conn) error {
    jsonStr, _ := json.Marshal(t)
    _, err := conn.Write([]byte(jsonStr))
        if err != nil {
            log.Fatal(err)
            return err
        }
    return nil
}


func (t *Message) parse(data []byte) {
    json.Unmarshal(data, t)
}


func getResponse(input *Message) Message {
    msg := new(Message)
    msg.Content = input.Content
    msg.Metadata = MetadataSchema{
        Timestamp: time.Now().Format("Monday, 02-Jan-06 15:04:05 MST"),
        Tag: "mambo",
        }
    return *msg
}


func DeferUserInterrupt(timeout int) {
    ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
    defer stop()

    select {
    case <-time.After(timeout * time.Second):
        fmt.Println("timeout signal received")
        stop()
        os.Exit(0)
    case <-ctx.Done():
		stop()
		fmt.Println("unix signal received")
		os.Exit(0)
    }
}

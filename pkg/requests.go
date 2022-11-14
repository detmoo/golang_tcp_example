// Desc: TCP echo server and client
package pkg

import (
    "encoding/json"
	"fmt"
	"net"
	"strings"
	"time"
)


func handleIncomingRequest(conn net.Conn) {
    // store incoming data
    buffer := make([]byte, 1024)
    _, err := conn.Read(buffer)
    if err != nil {
        log.Fatal(err)
    }
    request := parse(buffer)
    // respond
    response := getResponse(request)
    write(response)

    // close conn
    conn.Close()
}


type Metadata struct {
    time string `json:"time"`
    tag string `json:"tag"`
}


type Message struct {
    content string `json:"content"`
    metadata Metadata `json:"metadata"`
}


func (t Message) write() string {
    jsonStr, _ := json.Marshal(t)
    conn.Write([]byte(jsonStr))
}


func (t Message) parse(data []byte) {
    msg := new(Message)
    json.Unmarshal(data, msg)
    return msg
}


func getResponse(input Message) Message {
    msg := Message{
        content: input.content,
        metadata: Metadata{
            time: time.Now().Format("Monday, 02-Jan-06 15:04:05 MST"),
            tag: "mambo"
        }
    }
    return msg
}

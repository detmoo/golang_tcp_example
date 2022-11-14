// Desc: TCP echo server and client
package pkg

import (
    "encoding/json"
    "log"
	"net"
	"time"
)


func handleIncomingRequest(conn net.Conn) error {
    // store incoming data
    buffer := make([]byte, 1024)
    _, err := conn.Read(buffer)
    if err != nil {
        log.Fatal(err)
        return err
    }
    receivedMsg := new(Message)
    request := receivedMsg.parse(buffer)
    // respond
    response := getResponse(request)
    response.write(conn)

    // close conn
    conn.Close()
    return nil
}


type Metadata struct {
    time string `json:"time"`
    tag string `json:"tag"`
}


type Message struct {
    content string `json:"content"`
    metadata Metadata `json:"metadata"`
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


func (t Message) parse(data []byte) Message {
    json.Unmarshal(data, t)
    return t
}


func getResponse(input Message) Message {
    msg := Message{
        content: input.content,
        metadata: Metadata{
            time: time.Now().Format("Monday, 02-Jan-06 15:04:05 MST"),
            tag: "mambo",
        },
    }
    return msg
}

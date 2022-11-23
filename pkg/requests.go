// Desc: TCP echo server and client
package pkg

import (
    "encoding/json"
    "log"
	"net"
	"time"
)


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
    if err := json.Unmarshal(data, t); err != nil {
        log.Fatal("error on json unmarshall of byte slice:", err)
    }
}


func HandleIncomingRequest(conn net.Conn) error {
    // receive
    defer conn.Close()
    buffer := make([]byte, 1024)
    size, err := conn.Read(buffer)
    if err != nil {
        log.Fatal(err)
        return err
    }
    receivedMsg := new(Message)
    receivedMsg.parse(buffer[:size])

    // respond
    response := getResponse(receivedMsg)
    err = response.write(conn)
    if err != nil {
        log.Fatal(err)
        return err
    }

    return nil
}


func MakeRequest(msg Message, conn net.Conn) (answer Message, err error) {
    // send
    if err = msg.write(conn); err != nil {
        log.Fatal(err)
    }

    // receive
    output := make([]byte, 1024)
    size, err := conn.Read(output)
    if err != nil {
        log.Fatal(err)
    }

    answer.parse(output[:size])
    return
}


func getResponse(input *Message) Message {
    msg := new(Message)
    msg.Content = "TCP listener received Message.Content: "+input.Content
    msg.Metadata = MetadataSchema{
        Timestamp: time.Now().Format("Monday, 02-Jan-06 15:04:05.0000 MST"),
        Tag: "TCPServer",
        }
    return *msg
}

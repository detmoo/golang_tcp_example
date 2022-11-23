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
    json.Unmarshal(data, t)
}


func HandleIncomingRequest(conn net.Conn) error {
    // receive
    // defer conn.Close()
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
    err = response.write(conn)
    if err != nil {
        log.Fatal(err)
        return err
    }

    conn.Close()
    return nil
}


func MakeRequest(msg Message, conn net.Conn) (answer Message, err error) {
    output := make([]byte, 1024)
    if _, err = conn.Read(output); err != nil {
        log.Fatal(err)
    }
    answer.parse(output)
    return
}


func getResponse(input *Message) Message {
    msg := new(Message)
    msg.Content = input.Content
    msg.Metadata = MetadataSchema{
        Timestamp: time.Now().Format("Monday, 02-Jan-06 15:04:05 MST"),
        Tag: "server boo!",
        }
    return *msg
}

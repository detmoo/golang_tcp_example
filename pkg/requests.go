// Desc: TCP echo server and client
package pkg

import (
    "encoding/json"
    "log"
	"net"
	"os"
	"time"
)


var SERVER_MSG_TAG_ENV_VAR string = "BD_SERVER_TAG"
var CLIENT_MSG_TAG_ENV_VAR string = "BD_CLIENT_TAG"


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


func (t *Message) compose(content, tag string) Message {
    t.Content = content
    t.Metadata = MetadataSchema{
        Timestamp: time.Now().Format("Monday, 02-Jan-06 15:04:05.0000 MST"),
        Tag: tag,
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
    response := getResponse(receivedMsg.Content)
    err = response.write(conn)
    if err != nil {
        log.Fatal(err)
        return err
    }

    return nil
}


func MakeRequest(content string, conn net.Conn) (answer Message, err error) {
    // prepare message
    msg := composeMessage(content, getMsgTagFromEnv(CLIENT_MSG_TAG_ENV_VAR))

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

    // load response
    answer.parse(output[:size])
    return
}


func getResponse(requestContent string) Message {
    msg := new(Message)
    msg.compose("TCP listener received: "+requestContent, getMsgTagFromEnv(SERVER_MSG_TAG_ENV_VAR))
    return *msg
}


func getMsgTagFromEnv(key string) string {
    msgTag := os.Getenv(key)
    if len(msgTag) == 0 {
        msgTag = "untagged-tcp-endpoint"
    }
    return msgTag
}

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"golang.org/x/net/websocket"
)

const (
	TYPE_MSG int = iota
	TYPE_GETID
	TYPE_SETID
	TYPE_NEWUSER
)

type DataGram struct {
	Type    int         `json:"type"`
	Version int         `json:"version"`
	Data    interface{} `json:"data"`
}
type Message struct {
	SenderID   string    `json:"sender_id"`
	ReceiverID string    `json:"receiver_id"`
	Content    string    `json:"content"`
	TimeStamp  time.Time `json:"time_stamp"`
	MsgVersion int       `json:"msg_version"`
}

type Client struct {
	id string
	ch chan *DataGram
}

var msgToMux chan *DataGram

var clientChan chan *Client
var clients map[string]*Client

func main() {
	clientChan = make(chan *Client)
	msgToMux = make(chan *DataGram)
	clients = make(map[string]*Client)
	go mux()
	http.Handle("/", websocket.Handler(HandleMessage))

	if err := http.ListenAndServe(":60012", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
func mux() {
	var newClient *Client
	var msg *DataGram
	for {
		select {
		case newClient = <-clientChan:
			clients[newClient.id] = newClient
		case msg = <-msgToMux:
			log.Printf("muxing %v", msg)

			for _, c := range clients {
				log.Printf("copying %v", msg)
				c.ch <- msg
			}

		}

	}

}

// 注意是持久连接
func HandleMessage(ws *websocket.Conn) {
	var err error
	clientch := make(chan *DataGram)
	client := &Client{
		ws.Request().RemoteAddr,
		clientch,
	}
	// 握手
	dg := &DataGram{
		Type:    TYPE_SETID,
		Version: 0,
		Data:    client.id,
	}
	msg := dg.Marshal()
	if err = websocket.Message.Send(ws, msg); err != nil {
		fmt.Println("Can't send")
		return
	}

	clientChan <- client
	go func() {

		for {
			var reply string

			if err = websocket.Message.Receive(ws, &reply); err != nil {
				fmt.Println("Can't receive")
				break
			}
			dg := &DataGram{}
			err := dg.Unmarshal(reply)

			if err != nil {
				log.Printf("Received failed : %v", dg)
				break

			}
			log.Printf("Received : %v", dg)
			msgToMux <- dg

		}
	}()

	for {
		dataGram := <-client.ch
		msg := dataGram.Marshal()
		log.Printf("Sending : %s", msg)
		if err = websocket.Message.Send(ws, msg); err != nil {
			fmt.Println("Can't send")
			break
		}

	}
}
func (m *DataGram) Unmarshal(s string) error {
	err := json.Unmarshal([]byte(s), m)
	if err != nil {
		log.Printf("Fail to unmarshal %s ,%v", s, err)

	}

	return err
}
func (m *DataGram) Marshal() string {
	msg, e := json.Marshal(m)
	if e != nil {
		log.Printf("Fail to marshal %v ,%v", m, e)
		return e.Error()

	}
	return string(msg)

}
func Echo(ws *websocket.Conn) {
	var err error

	for {
		var reply string

		if err = websocket.Message.Receive(ws, &reply); err != nil {
			fmt.Println("Can't receive")
			break
		}
		fmt.Println("Received back from client: " + reply)

		msg := "Received:  " + reply
		fmt.Println("Sending to client: " + msg)

		if err = websocket.Message.Send(ws, msg); err != nil {
			fmt.Println("Can't send")
			break
		}
	}
}

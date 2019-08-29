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
	TYPE_MSG     = 0
	TYPE_GETID   = 1
	TYPE_SETID   = 2
	TYPE_NEWUSER = 3
)

type DataGram struct {
	Type    int         `json:"type"`
	Version int         `json:"version"`
	Data    interface{} `json:"data"`
	Message *Message    `json:"message"`
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

var dgToMux chan *DataGram

var clientChan chan *Client
var clients map[string]*Client

func main() {
	clientChan = make(chan *Client)
	dgToMux = make(chan *DataGram)
	clients = make(map[string]*Client)
	go mux()
	http.Handle("/", websocket.Handler(HandleMessage))
	log.Print("Server string")

	if err := http.ListenAndServe(":60012", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
func dealWithNewDatagram(datagram *DataGram) {

	log.Printf("muxing %v", datagram)
	switch datagram.Type {

	case TYPE_SETID:
		// 新用户加入消息,
		if datagram.Message != nil {
			m := datagram.Message
			c := clients[m.ReceiverID]
			c.ch <- datagram
			newMsg := m.Copy()
			newMsg.ReceiverID = "all"
			datagramNew := NewDataGram(TYPE_NEWUSER, newMsg, nil)
			datagram = datagramNew

			// 发送一个全体消息通知新人到来
			dgToMux <- datagramNew
		}

	case TYPE_NEWUSER:
		log.Printf("sending to all %v", datagram)
		for _, c := range clients {
			c.ch <- datagram
		}
		return

	case TYPE_MSG:
		// 直接给全体
		for _, c := range clients {
			c.ch <- datagram
		}

		// c := clients[datagram.Message.ReceiverID]
		// c.ch <- datagram

	}
}
func mux() {
	var newClient *Client
	var datagram *DataGram
	for {
		select {
		case newClient = <-clientChan:
			clients[newClient.id] = newClient
		case datagram = <-dgToMux:
			dealWithNewDatagram(datagram)

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
		Message: NewMessage("system", client.id, "Welcome", time.Now()),
	}
	datagram := dg.Marshal()
	if err = websocket.Message.Send(ws, datagram); err != nil {
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
			dgToMux <- dg

		}
	}()

	for {
		dataGram := <-client.ch
		datagram := dataGram.Marshal()
		log.Printf("Sending : %s", datagram)
		if err = websocket.Message.Send(ws, datagram); err != nil {
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
	datagram, e := json.Marshal(m)
	if e != nil {
		log.Printf("Fail to marshal %v ,%v", m, e)
		return e.Error()

	}
	return string(datagram)

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
func NewDataGram(t int, msg *Message, data interface{}) *DataGram {
	return &DataGram{
		Type:    t,
		Version: 0,

		Data:    data,
		Message: msg,
	}

}
func NewMessage(senderid string, receiverid string, content string, timeStamp time.Time) *Message {

	return &Message{
		SenderID:   senderid,
		ReceiverID: receiverid,
		Content:    content,
		TimeStamp:  timeStamp,
		MsgVersion: 1,
	}
}
func (m *Message) Copy() *Message {
	return NewMessage(
		m.SenderID,
		m.ReceiverID,
		m.Content,
		m.TimeStamp,
	)

}

// gopl 8.10
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

// 用一个传送字符串的通道表示client
type client chan string

var (
	// 主goroutine 通过这个通道传递新客户端登录的信息, 然后通知广播goroutine将消息传给每个client
	entering = make(chan client)
	// 所有客户端将退出消息通过这个通道传给广播goroutine
	leaving = make(chan client)
	// 所有客户端将文本消息通过这个通道传给广播
	message = make(chan string)
)

// main 接受请求 启动其他goroutine
func main() {
	host := "0.0.0.0:60066"
	listener, err := net.Listen("tcp", host)
	if err != nil {
		log.Fatalf("Fail to listen, %v", err)
	}
	log.Printf("Listening on %s", host)
	go broadcast()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf(err.Error())
		}
		go handler(conn)
	}

}

// 处理新的客户端连接，
func handler(conn net.Conn) {
	cli := make(chan string)
	// 启动后台goroutine 接收其他client的消息,写入对应的连接，消息由broadcast 转发
	go func(conn net.Conn, ch client) {
		for msg := range ch {
			fmt.Fprintln(conn, msg)
		}
	}(conn, cli)
	// 为了简单将登录登出消息放在这里
	host := conn.RemoteAddr().String()
	log.Printf("new client attached: %s", host)
	cli <- "Welcome " + host
	cli <- "Tell me you name:"
	whoami := "Noname"

	input := bufio.NewScanner(conn)
	for input.Scan() {
		s := input.Text()
		if strings.Trim(s, "\n\t \r") != "" {
			whoami = s
			log.Printf("%s: name set[%s]", host, whoami)
			break
		}
	}
	entering <- cli
	message <- fmt.Sprintf("%s(%s) has successfully landed!", whoami, host)
	for input.Scan() {
		msg := input.Text()
		message <- fmt.Sprintf("[%s]: %s", whoami, msg)
	}
	leaving <- cli
	message <- whoami + "has left."
	log.Printf("detached: %s", host)
	conn.Close()

}

// 广播goroutine
// 消息分发中心
func broadcast() {
	clients := make(map[client]bool)
	for {
		select {
		// 将消息分发给所有的client
		case msg := <-message:
			for cli := range clients {
				cli <- msg
			}
		case cli := <-entering:
			clients[cli] = true
		case cli := <-leaving:
			delete(clients, cli)
			close(cli)
		}
	}

}

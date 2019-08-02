package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"
)

func main() {

	port := "8000"
	if len(os.Args) < 3 || os.Args[1] != "-p" {

		log.Printf("Using default port %s.\nUsage: %s -p port", port, os.Args[0])

	} else {

		port = os.Args[2]
	}
	listenHost := fmt.Sprintf("localhost:%s", port)
	listener, err := net.Listen("tcp", listenHost)
	log.Printf("Listening %s", listenHost)
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)

	}

}
func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:01:05\n"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)

	}

}

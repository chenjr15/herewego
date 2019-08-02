package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

// TimeItem contains time and zone info
type TimeItem struct {
	Zone string
	Host string
	Time string
}

func main() {
	c := make(chan *TimeItem)
	itemMap := parser(os.Args[1:])
	for _, v := range itemMap {
		go getTime(v, c)

	}
	if len(itemMap) == 0 {
		log.Print("Empty input.")
		return
	}

	for t := range c {
		fmt.Printf("----------")
		if t == nil {
			return
		}
		for _, v := range itemMap {
			fmt.Println(v.Zone, v.Time)
		}
	}
}
func parser(items []string) map[string]*TimeItem {
	itemMap := make(map[string]*TimeItem)

	for _, item := range items {
		fileds := strings.Split(item, "=")
		zone, hostport := fileds[0], fileds[1]
		itemMap[zone] = &TimeItem{
			Zone: zone,
			Host: hostport,
			Time: "",
		}
	}
	return itemMap
}
func getTime(ti *TimeItem, c chan *TimeItem) {
	conn, err := net.Dial("tcp", ti.Host)
	if err != nil {
		log.Fatalf("Fail to connect to %s , %s", ti.Host, err)
	}
	defer conn.Close()
	defer close(c)
	input := bufio.NewScanner(conn)
	for input.Scan() {
		ti.Time = input.Text()
		c <- ti
	}

}

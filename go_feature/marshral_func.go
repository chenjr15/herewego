package main

import (
	"encoding/json"
	"log"
)

type S struct {
	A int
	B *int
	C float64
	D func() string
	E chan struct{}
}

func main() {
	s := S{
		A: 1,
		B: nil,
		C: 12.15,
		D: func() string {
			return "NowCoder"
		},
		E: make(chan struct{}),
	}

	_, err := json.Marshal(s)
	// 不支持序列化函数 json: unsupported type: func() string

	if err != nil {
		log.Printf("err occurred..")
		log.Println(err)
		return
	}
	log.Printf("everything is ok.")
	return

}

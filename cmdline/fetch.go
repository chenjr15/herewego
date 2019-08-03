package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var url = flag.String("url", "", "The url to get.")

func main() {
	flag.Parse()
	resp, err := http.Get(*url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(string(bs))

}

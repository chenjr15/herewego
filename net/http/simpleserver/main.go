package main

import (
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/pelletier/go-toml"
)

var configFilename = flag.String("c", "config.toml", "toml config file")
var config *Config

func main() {

	log.Printf("Listening: %s\n", config.HTTPConfig.Host)
	http.HandleFunc("/", GetGeneralHandler(config))
	http.HandleFunc(config.RouteConfig.Echo, Echo)
	http.HandleFunc(config.RouteConfig.Time, ShowTime)
	http.HandleFunc(config.RouteConfig.ListDir, ListDir)

	log.Fatal(http.ListenAndServe(config.HTTPConfig.Host, nil))
}
func init() {
	flag.Parse()
	log.Printf("Welcome to simple http server.")

	config = GetDefaultConfig()
	configContent, err := ioutil.ReadFile(*configFilename)

	if os.IsNotExist(err) {
		log.Printf("Writing default configuration to %s", *configFilename)
		ioutil.WriteFile(*configFilename, GetDefaultContents(), 0644)
		return
	} else if err != nil {
		log.Fatal(err)
	}
	err = toml.Unmarshal(configContent, config)
	if err != nil {
		log.Fatalf("Failed to unmarshal configfile ,%v ", err)
	}
}

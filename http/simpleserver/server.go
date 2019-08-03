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

type HTTPConfig struct {
	Host string `toml:"host"`
}
type Config struct {
	HTTPConfig HTTPConfig
}

func main() {

	log.Printf("Listening: %s\n", config.HTTPConfig.Host)
	http.HandleFunc("/", helloworld)

	log.Fatal(http.ListenAndServe(config.HTTPConfig.Host, nil))
}
func helloWorld(w http.ResponseWriter, r *http.Request) {
	n, err := w.Write([]byte("Hello World!"))
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("Write %d bytes", n)
	return
}
func init() {
	flag.Parse()
	log.Printf("Welcome to simple http server.")

	config = getDefaultConfig()
	configContent, err := ioutil.ReadFile(*configFilename)

	if os.IsNotExist(err) {
		log.Printf("Writing default configuration to %s", *configFilename)
		ioutil.WriteFile(*configFilename, getDefaultContents(), 0644)
		return
	} else if err != nil {
		log.Fatal(err)
	}
	err = toml.Unmarshal(configContent, config)
	if err != nil {
		log.Fatalf("Failed to unmarshal configfile ,%v ", err)
	}
}
func getDefaultConfig() (defaultConfig *Config) {

	return &Config{
		HTTPConfig: HTTPConfig{
			Host: "localhost:60066",
		},
	}
}

func getDefaultContents() []byte {
	bytes, err := toml.Marshal(getDefaultConfig())
	if err != nil {
		return []byte{}
	}
	return bytes

}

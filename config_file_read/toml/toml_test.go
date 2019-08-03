package toml

import (
	"log"
	"os"
	"testing"

	"github.com/pelletier/go-toml"
)

func TestGetKey(t *testing.T) {
	filename := "./example.toml"
	tree, err := toml.LoadFile(filename)
	if err != nil {
		t.Errorf("Failed to load %s , %s", filename, err)
	}
	title := tree.Get("title").(string)
	dbServer := tree.Get("database.server").(string)
	t.Logf("Title : '%s' database.server: '%s'", title, dbServer)

	clients := tree.Get("clients.data").([]interface{})
	t.Logf("slice %[1]T '%[1]v' ", clients[0])

}

type Database struct {
	Server        string `toml:"server"`
	Ports         []int  `toml:"ports"`
	ConnectionMax int    `toml:"connection_max"`
	Enabled       bool   `toml:"enabled"`
}
type Config struct {
	Database Database `toml:"database"`
}

func TestMarshal(t *testing.T) {
	database := Database{
		Server:        "localhost",
		Ports:         []int{2, 3},
		ConnectionMax: 233,
		Enabled:       true,
	}
	config := Config{
		Database: database,
	}
	s, err := toml.Marshal(config)
	if err != nil {
		t.Fatalf("Failed to marshal %v,%v\n", database, err)
	}
	t.Log(string(s))
	testfile := "test.toml"
	file, err := os.OpenFile(testfile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		t.Fatalf("Fail to open file %s , %v\n", testfile, err)
	}
	defer file.Close()
	n, err := file.Write(s)
	if err != nil {

		t.Fatalf("Fail to write file %s , %v \n", testfile, err)
	}
	t.Logf("write %d bytes to %s\n", n, testfile)

}

func TestUnmarshal(t *testing.T) {

	testfile := "test.toml"
	file, err := os.Open(testfile)
	if err != nil {
		t.Fatalf("Fail to open file %s , %v\n", testfile, err)
	}
	defer file.Close()
	s := make([]byte, 8192)
	n, err := file.Read(s)
	if err != nil {
		t.Fatalf("Fail to read file %s , %v \n", testfile, err)
	}
	s = s[:n]
	t.Logf("read %d bytes from  %s\n", n, testfile)
	config := Config{}

	toml.Unmarshal(s, &config)
	if err != nil {
		t.Fatalf("Fail to unmarshal %s , %v \n", string(s), err)
	}
	t.Log(config)

}

func init() {

	testfile := "test.toml"
	err := os.Remove(testfile)
	if err != nil {
		log.Fatal(err)

	}
}

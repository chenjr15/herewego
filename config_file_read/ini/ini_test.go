package main

import (
	"log"
	"testing"

	"gopkg.in/ini.v1"
)

func TestIniRead(t *testing.T) {

	cfg, err := ini.Load("example.ini")
	if err != nil {
		t.Fatalf("Fail to load ini file : %s,%s", "example.ini", err)
	}
	t.Log("App Mode:", cfg.Section("").Key("app_mode").String())
	t.Log("Data Path:", cfg.Section("paths").Key("data").String())
	notExists := cfg.Section("paths").Key("not_exists").String()
	if notExists != "" {
		log.Fatal("Error while reading not_exists")
	}
	t.Log("Not Exists :", notExists)
	httpProt := cfg.Section("server").Key("protocol").In("http", []string{"http", "https"})
	if httpProt != "http" {
		t.Fatal("Wrong key protocol:", httpProt)
	}

	t.Log("Server Protocol:", httpProt)
	t.Log("Undefined :", cfg.Section("server").Key("undifined").In("default", []string{"http", "https"}))
	// 1111 是默认值
	t.Logf("Port number: %[1]T %[1]d\n", cfg.Section("server").Key("http_port").MustInt(1111))

}
func TestWrite(t *testing.T) {
	value := "set by go-ini"
	cfg, err := ini.Load("example.ini")
	if err != nil {
		t.Fatalf("Fail to load ini file : %s,%s", "example.ini", err)
	}
	cfg.Section("Section_A").Key("key_b").SetValue(value)
	cfg.SaveTo("example_modified.ini")
	cfg2, err := ini.Load("example_modified.ini")
	if err != nil {
		t.Fatalf("Fail to load ini file : %s,%s", "example_modified.ini", err)
	}
	if value != cfg2.Section("Section_A").Key("key_b").String() {
		t.Fatal("Write failed! need ", value)

	}

	t.Log("Write success!")
}

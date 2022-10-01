package main

import (
	"fmt"
	"os"
	"testing"
)

type args struct {
	n int
	k int
}

var tests []struct {
	name string
	args args
}

func init() {
	tests = []struct {
		name string
		args args
	}{
		{"100协程 65535长度", args{100, 65535}},
	}

}
func Test_printNChan(t *testing.T) {
	f, _ := os.OpenFile("printNChan.txt", os.O_WRONLY, 0)
	setupPrint(func(a ...interface{}) (n int, err error) {
		return fmt.Fprintln(f, a...)
	})
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			printNChan(tt.args.n, tt.args.k)
		})
	}
	_ = f.Close()
}

func Test_printNMutex(t *testing.T) {

	f, _ := os.OpenFile("printNMutex.txt", os.O_WRONLY, 0)
	setupPrint(func(a ...interface{}) (n int, err error) {
		return fmt.Fprintln(f, a...)
	})
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			printNMutex(tt.args.n, tt.args.k)
		})
	}
	_ = f.Close()

}

func Test_printNWaitGroup(t *testing.T) {
	f, _ := os.OpenFile("printNWaitGroup.txt", os.O_WRONLY, 0)
	setupPrint(func(a ...interface{}) (n int, err error) {
		//fmt.Println("1223")
		return fmt.Fprintln(f, a...)
	})
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			printNWaitGroup(tt.args.n, tt.args.k)

		})
	}

	_ = f.Close()

}

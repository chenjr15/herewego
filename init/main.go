package main

import (
	"fmt"
	"init/pkg"
)
import _ "init/pkg"

func main() {

	pkg.Hello("Hello!")
}
func init() {
	fmt.Println("init1 in main ")
}

func init() {
	fmt.Println("init2 in main ")
}

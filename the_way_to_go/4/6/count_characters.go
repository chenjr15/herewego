package main

import "fmt"
import "unicode/utf8"

func count(s string) (bytecount, runecount int) {
	bytecount = len(s)

	runecount = utf8.RuneCountInString(s)
	// fmt.Println(r)
	return bytecount, runecount
}

func main() {
	var strs = []string{"asSASA ddd dsjkdsjs dk", "asSASA ddd dsjkdsjsこん dk"}
	for i := 0; i < len(strs); i++ {
		s := strs[i]
		b, r := count(s)
		fmt.Printf("%v %v %v", s, b, r)
	}

}
 
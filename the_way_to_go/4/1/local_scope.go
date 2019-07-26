package main

var a = "G"

func main() {
	n()
	m()
	n()
	// no new line
	// GOG
}

func n() {
	print(a)
}
func m() {
	a := "O"
	print(a)
}

package main

var a = "G"

func main() {
	n()
	m()
	n()
	// GOO
}
func n() {
	print(a)
}

func m() {
	a = "O"
	print(a)
}

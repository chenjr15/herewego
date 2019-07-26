package main

var a string

func main() {
	a = "G"
	print(a)
	f1()
	// GOG
}
func f1() {
	a := "O"
	print(a)
	f2()

}

func f2() {
	// 这里会调用全局变量
	print(a)
}

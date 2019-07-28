package multi

import "fmt"

func multiTable() (result string) {
	result = "\n"
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			result += fmt.Sprintf("%d x %d = %d \t", j, i, i*j)
		}
		result += "\n"
	}
	return

}

package main

import "fmt"

func main() {
	arrF := [...]float32{1.1, 1.2, 1.3, 1.4, 1.5}
	slF := arrF[2:]

	fmt.Println("sum of array", arrF, "=", SumArray(arrF))
	fmt.Println("sum of slice", slF, "=", SumSlice(slF))
	avg, sum := SumAndAverage(slF)
	fmt.Println("average, sum of slice", slF, "=", avg, sum)
}
func SumAndAverage(sl []float32) (avg, sum float32) {
	for _, num := range sl {
		sum += num
	}
	avg = sum / float32(len(sl))
	return
}

func SumSlice(sl []float32) (sum float32) {
	for _, num := range sl {
		sum += num
	}
	return
}
func SumArray(arr [5]float32) (sum float32) {
	for _, num := range arr {
		sum += num
	}
	return
}

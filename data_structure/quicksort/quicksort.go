package main

import "fmt"

func main() {
	nums := []int{47, 85, 51, 81, 95, 73, 72, 96, 11, 40, 61}
	test(nums)
	test([]int{3, 3, 3, 3, 3, 3})
}

func test(nums []int) {

	fmt.Println("Before sort", nums)
	QuickSort(nums)
	fmt.Println("After sort", nums)
	for i := 1; i < len(nums); i++ {
		if nums[i] < nums[i-1] {
			fmt.Println("Wrong ", nums[i], nums[i-1])
		}
	}
}

// QuickSort 快速排序
func QuickSort(nums []int) {
	p := 0

	for len(nums) > 1 {

		p = Partition(nums)
		QuickSort(nums[0:p])
		nums = nums[p+1 : len(nums)]

	}

}

// Partition  快排划分函数
func Partition(nums []int) int {
	n := len(nums)
	if n <= 3 {
		if n == 2 && nums[0] > nums[1] {
			nums[0], nums[1] = nums[1], nums[0]
		}
		return 0
	}
	pivot := len(nums) - 1
	i, j := 0, pivot-1
	for i < j {
		for i < pivot && nums[i] < nums[pivot] {
			i++
		}
		for j > -1 && nums[j] > nums[pivot] {
			j--
		}
		// 还没有越界， 尝试交换
		if i < j {
			fmt.Println("swap ", i, j)
			nums[i], nums[j] = nums[j], nums[i]
			i++
			j--
		} else {
			// 指针相交了，分完了
			break
		}
	}
	nums[i], nums[pivot] = nums[pivot], nums[i]
	return i

}

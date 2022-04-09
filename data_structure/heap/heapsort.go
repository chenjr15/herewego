package main

import "fmt"

func main() {
	nums := []int{47, 85, 51, 81, 95, 73, 72, 96, 11, 40, 61}
	fmt.Println("Heap sort ", nums)
	HeapSort(nums)
	fmt.Println("After sort", nums)
	for i := 1; i < len(nums); i++ {
		if nums[i] < nums[i-1] {
			fmt.Println("Wrong ", nums[i], nums[i-1])
		}
	}
}

// HeapSort 堆排序
func HeapSort(nums []int) {
	// 堆排序， 就是用数组搞一个完全二叉树的情况，
	// 然后每次把最大的浮到顶上, 最后得到的是 最大的在最顶上
	// 这样就是建堆。建完堆之后，把最大的换到最后面，让堆的规模-1
	// 然后重复建堆，直到只剩下一个元素
	n := len(nums)

	// 找根节点 root = (p-1)/2
	// 找子节点 left = 2 * root+1, right= 2 * root +2

	// 建堆, 从最后一个叶子节点开始，逐个和他的根节点比较
	// 如果根节点比较小，就把当前节点换上去
	for n > 0 {
		p := n - 1
		for p > 0 {
			if nums[p] > nums[(p-1)/2] {
				// 交换当前节点和根节点
				nums[p], nums[(p-1)/2] = nums[(p-1)/2], nums[p]
			}
			p--
		}
		n--
		nums[n], nums[0] = nums[0], nums[n]
	}

}

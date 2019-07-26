package backpack

// Backpack problem
func Backpack(backSize int, nums []int) (total int) {
	dp := make([][]int, len(nums)+1)
	// 零值初始化
	for i := 0; i < len(nums)+1; i++ {
		dp[i] = make([]int, backSize+1)
	}
	cali := func(size, i int) int {
		if i < 1 {
			return 0
		}
		if size < nums[i-1] {
			return dp[i-1][size]
		}
		dontadd := dp[i-1][size]
		add := (nums[i-1] + dp[i-1][size-nums[i-1]])
		if add > dontadd {
			return add

		} else {
			return dontadd
		}

	}
	for i := 1; i <= len(nums); i++ {
		for j := 1; j <= backSize; j++ {
			dp[i][j] = cali(j, i)
		}
	}
	return dp[len(nums)][backSize]

}

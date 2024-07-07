package solutions

import "math"

// 动态规划
// dp[i][0]表示到i为止的乘积最小值
// dp[i][1]表示到i为止的乘积最大值
func maxProduct(nums []int) int {
	n := len(nums)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, 2)
	}
	ans := math.MinInt
	dp[0][0] = nums[0]
	dp[0][1] = nums[0]
	for i := 1; i < n; i++ {
		if nums[i] < 0 {
			dp[i][1] = max(dp[i-1][0]*nums[i], nums[i])
			dp[i][0] = min(dp[i-1][1]*nums[i], nums[i])
		} else {
			dp[i][1] = max(dp[i-1][1]*nums[i], nums[i])
			dp[i][0] = min(dp[i-1][0]*nums[i], nums[i])
		}
	}
	for i := 0; i < n; i++ {
		ans = max(ans, dp[i][1])
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

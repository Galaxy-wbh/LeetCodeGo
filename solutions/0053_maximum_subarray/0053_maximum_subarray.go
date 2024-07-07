package solutions

import "math"

// 方法一：
// 前缀和
func maxSubArray(nums []int) int {
	preSum := 0
	minPreSum := 0
	ans := math.MinInt
	for _, num := range nums {
		preSum += num                      // 当前前缀和
		ans = max(preSum-minPreSum, ans)   //减去前缀和的最小值
		minPreSum = min(preSum, minPreSum) // 更新前缀和的最小值
	}
	return ans
}

// 方法二：
// 动态规划
// dp[i]表示下标为i时的最大子序列和
// dp[i] = nums[i] 如果i单独组成一个子数组
//       = dp[i-1] + nums[i]    nums[i]和前面的子数组拼起来
// dp[i] = nums[i] // i = 0
// 		 = max(f[i-1], 0) + nums[i] // i >=1

func maxSubArray2(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	ans := math.MinInt
	dp[0] = nums[0]
	for i := 1; i < n; i++ {
		dp[i] = max(dp[i-1], 0) + nums[i]
		ans = max(ans, dp[i])
	}
	return max(dp[0], ans)
}

// 优化dp数组，只需要2个变量
func maxSubArray3(nums []int) int {
	f := 0
	ans := math.MinInt
	for _, x := range nums {
		f = max(f, 0) + x
		ans = max(ans, f)
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

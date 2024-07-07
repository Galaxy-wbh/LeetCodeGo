package solutions

// 回溯三问
// 当前操作：枚举nums[j], 需要满足j<i && nums[j] < nums[i]
// 子问题：以nums[i]结尾的LIS长度
// 下一个子问题：以nums[j]结尾的LIS长度
// dfs(i) = max{dfs(j)} + 1 , j < i && nums[j] < nums[i]
// ----------- 回溯做法，会超时 -------------
func lengthOfLIS(nums []int) int {
	n := len(nums)
	var dfs func(i int) int
	dfs = func(i int) int {
		res := 0
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				res = max(res, dfs(j))
			}
		}
		return res + 1
	}
	res := 0
	for i := 0; i < n; i++ {
		res = max(res, dfs(i))
	}
	return res
}

// 改成递推，递归->数组
func lengthOfLIS2(nums []int) int {
	n := len(nums)
	f := make([]int, n)
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				f[i] = max(f[i], f[j])
			}
		}
		f[i] += 1
	}
	res := 0
	for i := 0; i < n; i++ {
		res = max(res, f[i])
	}
	return res
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

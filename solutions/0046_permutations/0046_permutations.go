package solutions

// 排列型回溯
func permute(nums []int) [][]int {
	n := len(nums)
	ans := make([][]int, 0)
	path := make([]int, n)
	visited := make([]bool, n) // 标识是否访问过
	var dfs func(int)
	dfs = func(i int) {
		if i == n {
			tmp := make([]int, n)
			copy(tmp, path)
			ans = append(ans, tmp)
			return
		}
		for j := 0; j < n; j++ {
			if !visited[j] {
				visited[j] = true
				path[i] = nums[j]
				dfs(i + 1)
				visited[j] = false
			}
		}
	}
	dfs(0)
	return ans
}

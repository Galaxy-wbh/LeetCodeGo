package solutions

// 从答案角度出发，枚举下个该选的数
func subsets(nums []int) [][]int {
	n := len(nums)
	path := make([]int, 0)
	ans := make([][]int, 0)
	var dfs func(i int)
	dfs = func(i int) {
		tmp := make([]int, len(path))
		copy(tmp, path)
		ans = append(ans, tmp)
		for j := i; j < n; j++ {
			path = append(path, nums[j])
			dfs(j + 1)
			path = path[:len(path)-1]
		}
	}
	dfs(0)
	return ans
}

func subsets2(nums []int) [][]int {
	n := len(nums)
	path := make([]int, 0)
	ans := make([][]int, 0)
	var dfs func(i int)
	dfs = func(i int) {
		if i == n {
			tmp := make([]int, len(path))
			copy(tmp, path)
			ans = append(ans, tmp)
			return
		}
		// 不选
		dfs(i + 1)

		// 选
		path = append(path, nums[i])
		dfs(i + 1)
		path = path[:len(path)-1]
	}
	dfs(0)
	return ans
}

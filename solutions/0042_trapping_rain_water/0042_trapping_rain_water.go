package solutions

// 方法一：
// 计算每一个格子中的前缀最大值和后缀最大值，每个格子中能存储的水的容积为
// min(preMax[i], suffMax[i] ) - height[i]
// 时间复杂度O(n) 空间复杂度O(n)
func trap(height []int) int {
	n := len(height)
	preMax := make([]int, n)
	suffMax := make([]int, n)
	preMax[0] = height[0]
	for i := 1; i < n; i++ {
		preMax[i] = max(preMax[i-1], height[i])
	}
	suffMax[n-1] = height[n-1]
	for i := n - 2; i >= 0; i-- {
		suffMax[i] = max(suffMax[i+1], height[i])
	}
	res := 0
	for i := 0; i < n; i++ {
		res += max(min(preMax[i], suffMax[i])-height[i], 0)
	}
	return res
}

// 方法二：
// 方法一的进阶版，优化空间复杂度。双指针思路
// 当左边前缀最大值小于右边后缀最大值时，这个格子的水体积就可以算出来了，左指针往右走（因为水体积只和前缀最大、后缀最大有关）
// 时间复杂度O(n) 空间复杂度O(1)
func trap2(height []int) int {
	n := len(height)
	left, right := 0, n-1
	preMax, suffMax := 0, 0
	ans := 0
	for left <= right {
		preMax = max(preMax, height[left])
		suffMax = max(suffMax, height[right])
		if preMax < suffMax {
			ans += preMax - height[left]
			left++
		} else {
			ans += suffMax - height[right]
			right--
		}
	}
	return ans
}

// 方法三：
// 单调栈
// 找上一个更大元素，在找的过程中填坑
func trap3(height []int) int {
	ans := 0
	stack := make([]int, 0)
	for i, h := range height {
		for len(stack) > 0 && h >= height[stack[len(stack)-1]] {
			bottomH := height[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				break
			}
			left := stack[len(stack)-1]
			ans += (min(height[left], h) - bottomH) * (i - left - 1)
		}
		stack = append(stack, i)
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

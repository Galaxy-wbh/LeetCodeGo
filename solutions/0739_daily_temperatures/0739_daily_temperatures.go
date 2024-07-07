package solutions

// 从右往左，栈内元素为下一个更大的数存在栈里
func dailyTemperatures(temperatures []int) []int {
	n := len(temperatures)
	ans := make([]int, n)
	stack := make([]int, 0)
	for i := n - 1; i >= 0; i-- {
		t := temperatures[i]
		// 温度大于栈顶元素，pop
		for len(stack) > 0 && t >= temperatures[stack[len(stack)-1]] {
			stack = stack[:len(stack)-1]
		}
		// 栈内还有元素比该值大
		if len(stack) > 0 {
			ans[i] = stack[len(stack)-1] - i
		}
		stack = append(stack, i)
	}
	return ans
}

// 从左往右，栈内元素为还没找到下一个比其温度高的下标
func dailyTemperatures2(temperatures []int) []int {
	n := len(temperatures)
	ans := make([]int, n)
	stack := make([]int, 0)
	for i := 0; i < n; i++ {
		t := temperatures[i]
		for len(stack) > 0 && t > temperatures[stack[len(stack)-1]] {
			ans[stack[len(stack)-1]] = i - stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	return ans
}

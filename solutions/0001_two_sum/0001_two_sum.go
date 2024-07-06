package solutions

// map提前存储的思想，加快答案的查找
// 空间换时间，时间复杂度将至O(n)
func twoSum(nums []int, target int) []int {
	m := map[int]int{}
	for i, num := range nums {
		if index, ok := m[target-num]; ok {
			return []int{i, index}
		}
		m[num] = i
	}
	return []int{}
}

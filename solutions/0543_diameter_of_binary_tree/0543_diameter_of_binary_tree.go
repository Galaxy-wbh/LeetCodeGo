package solutions

import (
	"LeetCodeGo/structures"
)

type TreeNode = structures.TreeNode

func diameterOfBinaryTree(root *TreeNode) int {
	ans := 0
	var dfs func(root *TreeNode) int
	dfs = func(root *TreeNode) int {
		if root == nil {
			return -1
		}
		left := dfs(root.Left) + 1
		right := dfs(root.Right) + 1
		ans = max(ans, left+right)
		return max(left, right)
	}
	dfs(root)
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

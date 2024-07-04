package solutions

import (
	"LeetCodeGo/structures"
	"fmt"
	"testing"
)

type q206 struct {
	para206
	ans206
}

type para206 struct {
	one []int
}

type ans206 struct {
	one []int
}

func TestReverseList(t *testing.T) {
	qs := []q206{
		{
			para206{[]int{1, 2, 3, 4, 5}},
			ans206{[]int{5, 4, 3, 2, 1}},
		},
	}
	fmt.Printf("------------------------Leetcode Problem 206------------------------\n")

	for _, q := range qs {
		_, p := q.ans206, q.para206
		fmt.Printf("【input】:%v       【output】:%v\n", p, structures.List2Ints(reverseList(structures.Ints2List(p.one))))
	}
}

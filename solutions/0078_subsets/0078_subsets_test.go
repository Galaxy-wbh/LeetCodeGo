package solutions

import (
	"fmt"
	"testing"
)

type q78 struct {
	para78
	ans78
}

type para78 struct {
	one []int
}

type ans78 struct {
	one [][]int
}

func TestSubsets(t *testing.T) {
	qs := []q78{
		{
			para78: para78{[]int{1, 2, 3}},
			ans78:  ans78{[][]int{{}, {1}, {1, 2}, {1, 2, 3}, {1, 3}, {2}, {2, 3}, {3}}}},
	}
	fmt.Printf("------------------------Leetcode Problem 78------------------------\n")

	for _, q := range qs {
		_, p := q.ans78, q.para78
		fmt.Printf("【input】:%v       【output】:%v\n", p, subsets2(p.one))
	}
	fmt.Printf("\n\n\n")
}

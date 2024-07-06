package solutions

import "sort"

func groupAnagrams(strs []string) [][]string {
	m := map[string][]string{}
	for _, str := range strs {
		s := []byte(str)
		sort.Slice(s, func(i, j int) bool {
			return s[i] < s[j]
		})
		sortedStr := string(s)
		m[sortedStr] = append(m[sortedStr], str)
	}
	ans := make([][]string, 0)
	for _, val := range m {
		ans = append(ans, val)
	}
	return ans
}

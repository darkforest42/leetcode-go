package leetcode03

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func lengthOfLongestSubstring(s string) int {
	m := make(map[int32]int)
	var i, maxLen int
	i, maxLen = 0, 0
	for j, c := range s {
		if p, ok := m[c]; ok {
			i = max(p, i)
		}
		maxLen = max(maxLen, j - i + 1)
		m[c] = j+1
	}
	return maxLen
}


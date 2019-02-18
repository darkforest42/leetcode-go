package leetcode14

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	shortest := strs[0]
	for _, s := range strs {
		if len(s) < len(shortest) {
			shortest = s
		}
	}
	for i, c := range shortest {
		for j := 0; j < len(strs); j++ {
			if strs[j][i] != byte(c) {
				return strs[j][:i]
			}
		}
	}
	return shortest
}

package twoSum01

func TwoSum(nums []int, target int) []int {
	var result []int
	numIndex := make(map[int]int)
	for index, value := range nums {
		if v, ok := numIndex[target-value]; ok {
			result = append(result, v, index)
			return result
		}
		numIndex[value] = index
	}
	return result
}

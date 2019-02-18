package bucketSort347

func TopKFrequent(nums []int, k int) []int {
	frequencyForNum := make(map[int]int)
	var frequency int
	var ok bool
	for _, num := range nums {
		if frequency, ok = frequencyForNum[num]; !ok {
			frequency = 0
		}
		frequencyForNum[num] = frequency + 1

	}

	bucketsForNum := make([][]int, len(nums)+1)
	for num, frequency := range frequencyForNum {
		bucketsForNum[frequency] = append(bucketsForNum[frequency], num)
	}
	topK := []int{}
	for i:= len(bucketsForNum)-1; i>=0 && len(topK) < k; i-- {
		topK = append(topK, bucketsForNum[i]...)
	}
	return topK
}

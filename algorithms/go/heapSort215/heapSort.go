package heapSort215

func swap(nums []int, i, j int){
	tmp := nums[i]
	nums[i] = nums[j]
	nums[j] = tmp
}

func heapify(nums []int, n, i int) {
	maxPos := i
	for {
		if 2*i+1 < n && nums[2*i+1] > nums[i] { maxPos = 2*i + 1 }
		if 2*i+2 < n && nums[maxPos] < nums[2*i+2] {maxPos = 2*i + 2}
		if maxPos == i { break }
		swap(nums, i, maxPos)
		i = maxPos
	}
}

func sort(nums []int){
	k := len(nums) - 1
	for ; k > 1; {
		swap(nums, 1, k)
		k--
		heapify(nums, k, 1)
	}
}
func FindKthLargest(nums []int, k int) int {
	numsLen := len(nums)
	for i := numsLen/2; i > 0; i-- {
		heapify(nums, numsLen, i)
	}
	sort(nums)
	return nums[numsLen - k]
}

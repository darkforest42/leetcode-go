package quickSort215

func swap(nums []int, i, j int){
	tmp := nums[i]
	nums[i] = nums[j]
	nums[j] = tmp
}

func partition(nums []int, p , r int) int {
	priot := nums[r]
	i := p
	for j := p; j < r; j++ {
		if nums[j] < priot {
			swap(nums, i, j)
			i = i+1
		}
	}
	swap(nums, i, r)
	return i
}

func FindKthLargest(nums []int, k int) int {
	k = len(nums) - k
	l , h := 0, len(nums)-1
	var q int
	for l <= h {
		q = partition(nums, l, h)
		if q == k {
			return nums[q]
		}else if q < k {
			l = q+1
		}else {
			h = q-1
		}
	}
	return nums[q]
}

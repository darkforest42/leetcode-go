package singleElement540

func singleNonDuplicate(nums []int) int {
	l, h := 0, len(nums)-1
	mid := 0
	for l < h {
		mid = l + (h - l)/2
		if mid % 2 == 1{
			mid--
		}
		if nums[mid] == nums[mid+1] {
			l = mid + 2
		}else{
			h = mid
		}
	}
	return nums[l]
}

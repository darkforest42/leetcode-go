package searchMatrix74

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) < 1 || len(matrix[0]) < 1 {
		return false
	}
	m := len(matrix)
	n := len(matrix[0])

	l , h, mid := 0, m*n, 0
	for l + 1 < h{
		mid = l + (h-l)/2
		if matrix[mid/n][mid%n] > target {
			h = mid
		}else{
			l = mid
		}
	}
	return matrix[l/n][l%n] == target
}

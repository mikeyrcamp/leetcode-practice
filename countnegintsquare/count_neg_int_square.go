package countnegintsquare

// CountNegativeIntsNaive counts a sorted square of ints
// 0 is not a negative number. Solution below is O(n^2)
func CountNegativeIntsNaive(nums [][]int) int {
	var negCount int
	for _, row := range nums {
		for _, v := range row {
			if v < 0 {
				negCount++
			}
		}
	}
	return negCount
}

// CountNegativeInts counts a sorted square of ints
// 0 is not a negative number. Solution below is O(n^2)
func CountNegativeInts(nums [][]int) int {
	if len(nums) == 0 {
		return 0
	}

	var negCount int
	n := len(nums[0])

	for rowIdx, colIdx := 0, n-1; rowIdx < n && colIdx >= 0; rowIdx++ {
		// Find the first non-negative in the row
		for colIdx >= 0 {
			// If this one is negative check the next one to the right
			// move that way if it is
			if nums[rowIdx][colIdx] < 0 {
				break
			} else {
				colIdx--
			}
		}

		// if j >= 0 then we know we have a negative number
		if colIdx >= 0 {
			negCount += colIdx + 1
		}
	}

	return negCount
}

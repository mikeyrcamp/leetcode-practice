package secondlargest

// Edge Cases: len(nums) < 2
func SecondLargest(nums []int) int {
	if len(nums) < 2 {
		return -1
	}

	largest := -1
	secondLargest := -1
	for _, v := range nums {
		if largest == -1 {
			largest = v
		} else if v > largest {
			secondLargest = largest
			largest = v
		} else if secondLargest == -1 {
			secondLargest = v
		} else if v > secondLargest {
			secondLargest = v
		}
	}
	return secondLargest
}

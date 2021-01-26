package sumpairstoten

func PairSumIsTen(nums []int) []int {
	values := map[int]bool{}

	for _, v := range nums {
		rem := 10 - v
		if _, ok := values[rem]; ok {
			return []int{rem, v}
		}
		values[v] = true
	}

	return nil
}

package search

func linearsearch(array []int, value int) int {
	for i, v := range array {
		if v == value {
			return i
		}
	}

	return -1
}

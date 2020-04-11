func countElements(arr []int) int {
	mapped := make(map[int]int)
	for _, i := range arr {
			mapped[i]++
	}
	sum := 0
	for k, _ := range mapped {
			sum += mapped[k-1]
	}
	return sum
}
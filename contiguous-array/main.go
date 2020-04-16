// Loop each element of the array, and take the count up to that point (0 = -1, 1 = 1)
// Have a hashmap which contains the index at which that count occurred.
// Check against that hashmap each time, and use this for the maxlength calculation

func findMaxLength(nums []int) int {
	var maxlength int
	mapped := make(map[int]int)
	mapped[0] = -1
	var count int
	for i:= 0; i<len(nums); i++ {
			val := getVal(nums[i])
			count += val
			c, ok := mapped[count]
			if !ok {
					mapped[count] = i
			} else {
					if i - c > maxlength {
							maxlength = i - c
					}
			}
	}
	return maxlength
}

func getVal(n int) int {
	if n == 0 {
			return -1
	}
	return 1
}
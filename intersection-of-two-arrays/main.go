func intersect(nums1 []int, nums2 []int) []int {
	//     basic boi version using a frequency map for one of the arrays and a loop checker for the other
			map1 := make(map[int]int)
			for _, v := range nums1 {
					map1[v] += 1
			}
			
			var results []int
			for _, v := range nums2 {
					if map1[v] > 0 {
							results = append(results, v)
							map1[v] -= 1
					}
			}
			
			return results
	}
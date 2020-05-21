func merge(nums1 []int, m int, nums2 []int, n int) {
	// two pointers
	p1, p2 := 0, 0
	for i := 0; i < m+n; i++ {
		// if you have finished with nums2
		if p2 >= n {
			break
		}
		// if you have reached the end of nums1's values, just copy in the rest of nums2
		if p1 >= m {
			nums1[i] = nums2[p2]
			p2++
			continue
		}
		// if no merge needed, continue
		if nums1[i] < nums2[p2] {
			p1++
			continue
		}
		// shift all elements of nums1 past i right by 1
		for j := m + n - 1; j >= i+1; j-- {
			nums1[j] = nums1[j-1]
		}
		// add nums2 element to nums1
		nums1[i] = nums2[p2]
		p2++
	}
}
func merge(nums1 []int, m int, nums2 []int, n int) {
	// assuming nums1 has enough space to hold elements from nums2
	var pointer int
	for i := 0; i < len(nums1) && n > 0; i++ {
		// if we have successfully merged up to this point
		if i >= m {
			nums1[i] = nums2[pointer]
			pointer++
			continue
		}

		if nums1[i] > nums2[pointer] {
			//swap the values around
			temp := nums1[i]
			nums1[i] = nums2[pointer]
			nums2[pointer] = temp
			tempPointer := pointer
			// insertion sort the swapped number
			for tempPointer < n-1 && nums2[tempPointer] > nums2[tempPointer+1] {
				t := nums2[tempPointer]
				nums2[tempPointer] = nums2[tempPointer+1]
				nums2[tempPointer+1] = t
				tempPointer++
			}
		}
	}
}
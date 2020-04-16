func productExceptSelf(nums []int) []int {
	//     without division is hard, since theres an obvious O(2n) solution using division
	//     only in O(n) is tricky as well, since theres the other obvious O(n^2) brute force approach
	//     the other approach is another O(2n) approach, using 3 arrays in total (so not constant space)
			left := make([]int,len(nums))
			left[0] = 1
			right := make([]int,len(nums))
			right[len(nums)-1] = 1
			l := 1
			r := 1
			for i := 1; i < len(nums); i++ {
					l = l * nums[i-1]
					r = r * nums[len(nums)-i]
					left[i] = l
					right[len(nums)-1-i] = r
			}
			for i := 0; i < len(nums); i++ {
					nums[i] = left[i] * right[i]
			}
			return nums
	}
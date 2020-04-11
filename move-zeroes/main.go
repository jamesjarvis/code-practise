func moveZeroes(nums []int) {
		var countzero int
		for i:=0; i<len(nums);i++{
				if nums[i]==0{
						countzero++
				} else if countzero > 0 {
						nums[i-countzero]=nums[i]
						nums[i] = 0
				}
		}
}
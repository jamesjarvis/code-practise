import "sort"

func threeSum(nums []int) [][]int {
    var solutions [][]int
    sort.Ints(nums)
    for i:= 0; i<len(nums);i++{
        if i != 0 && nums[i] == nums[i-1] {
            continue
        }
        seenCount := make(map[int]int)
        for j:= i+1; j<len(nums);j++ {
            if seenCount[-nums[i]-nums[j]] > 0 {
                solutions = append(solutions, []int{nums[i],nums[j],-nums[i]-nums[j]})
//                 Skip duplicates
                for j + 1 < len(nums) && nums[j+1] == nums[j] {
                    j++
                }
            }
            seenCount[nums[j]]++
        }
    }
    return solutions
}

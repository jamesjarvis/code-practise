import "sort"

func threeSum(nums []int) [][]int {
    var solutions [][]int
    sort.Ints(nums)
    for i:= 0; i<len(nums);i++{
        if i != 0 && nums[i] == nums[i-1] {
            continue
        }
        j := i+1
        k := len(nums) -1
        for j < k {
            if nums[i] + nums[j] + nums[k] == 0 {
                solutions = append(solutions, []int{nums[i],nums[j],nums[k]})
                j++
//                 never let j refer to the same value twice in an output to avoid duplicates
                for j<k && nums[j]==nums[j-1] {
                    j++
                }
            } else if nums[i] + nums[j] + nums[k] < 0 {
                j++
            } else {
                k--
            }
        }
    }
    return solutions
}

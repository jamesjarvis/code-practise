import "math"

func maxSubArray(nums []int) int {
    max := nums[0]
    currMax := nums[0]
    for i := 1; i< len(nums); i++ {
        currMax = int(math.Max(float64(nums[i]),float64(nums[i]+currMax)))
        max = int(math.Max(float64(max),float64(currMax)))
    }
    return max
}

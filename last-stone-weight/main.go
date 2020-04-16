import "sort"

func lastStoneWeight(stones []int) int {
    for len(stones) > 1 {
        sort.Ints(stones)
        end := len(stones)-1
        if stones[end] > stones[end-1] {
            stones[end-1] = stones[end]-stones[end-1]
        } else if stones[end] < stones[end-1] {
            stones[end-1] = stones[end-1]-stones[end]
        } else {
            stones[end-1] = 0
        }
        stones = stones[:end]
    }
    return stones[0]
}
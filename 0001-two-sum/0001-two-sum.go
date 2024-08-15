func twoSum(nums []int, target int) []int {
    sp := make(map[int]int)

    for index, val := range nums {
        complement := target - val
        // Check if the complement exists in the map
        if complementIndex, exists := sp[complement]; exists {
            return []int{complementIndex, index}
        }
        sp[val] = index

    }
    return []int{}
}
func subsets(nums []int) [][]int {
    result := make([][]int, 0, 1<<len(nums)) 
    current := make([]int, 0, len(nums)) 
    bacTrack(nums, current, 0, &result)
    return result
}

func bacTrack(nums []int, current []int, start int, result *[][]int){
    temp := make([]int, len(current))
    copy(temp, current)
    *result= append(*result,temp)
    for i := start; i<len(nums); i++{
        current = append(current, nums[i])
        bacTrack(nums, current, i+1, result)
        current = current[:len(current)-1]
    }
}
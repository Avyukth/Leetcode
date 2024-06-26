func distinctAverages(nums []int) int {
    sort.Ints(nums)
    left, right := 0, len(nums)-1
    averages := make(map[float64]struct{})

    for left<right{
        minval := nums[left]
        maxval := nums[right]
        average := float64(minval+maxval)/2
        averages[average]=struct{}{}
        left++
        right--

    }
    return len(averages)
}
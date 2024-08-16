func findTargetSumWays(nums []int, target int) int {
    total := 0
    for _, num := range nums{
        total += num
    }
    if total < abs(target) || (total+target)%2 != 0{
        return 0
    } 
    subsetS := (total+target)/2
    dp := make([]int, subsetS+1)
    dp[0] = 1
    for _, num := range nums{
        for i := subsetS ; i>= num ;i --{
            dp[i] += dp[i - num]
        }
    }
    return dp[subsetS]
}
func abs(i int) int{
    if i<0{
        return - i
    }
    return i
}
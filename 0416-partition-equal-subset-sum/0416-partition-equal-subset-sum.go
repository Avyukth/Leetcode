func canPartition(nums []int) bool {
    totalS := 0
    for _, num := range nums{
        totalS += num
    }
    if totalS %2 !=0 {
        return false
    }
    target := totalS/2
    dp := make([]bool, target+1)
    dp[0]=true
    for _, num := range nums{
        for j:=target;  j >=num; j--{
            dp[j] = dp[j] || dp[j-num]
        }
    }
    return dp[target]
}
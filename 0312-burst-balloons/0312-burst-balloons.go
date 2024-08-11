func maxCoins(nums []int) int {
    n:= len(nums)
    newNums := make([]int, n+2)
    newNums[0], newNums[n+1] =1,1
    copy(newNums[1:], nums)
    dp := make([][]int , n+2);
    for i := range dp {
        dp[i] = make([]int, n+2)
    }
    for lt :=1 ; lt <=n ; lt++ {
        for left :=1; left <=n-lt+1; left++{
            right :=  left +lt-1
            for i :=left; i<= right; i++{
                coins := dp[left][i-1] + newNums[left-1]* newNums[i]* newNums[right+1]+ dp[i+1][right]
                if coins > dp[left][right]{
                     dp[left][right] = coins
                }
            }
        }
    } 
    return dp[1][n]
}
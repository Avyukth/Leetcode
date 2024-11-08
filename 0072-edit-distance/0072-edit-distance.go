func minDistance(word1 string, word2 string) int {
    m,n := len(word1), len(word2)
    dp := make([][]int, m+1)
    for i := range dp{
        dp[i] = make ([]int, n+1)
        dp[i][0] = i
    }
    for i:=0; i <n+1; i++{
        dp[0][i] = i
    }

    for i:=1; i <m+1; i++{
        for j:=1;j <n+1; j++{
            if word1[i-1] == word2[j-1]{
                dp[i][j] = dp[i-1][j-1]
            }else{
                dp[i][j] = 1+ min(min(dp[i-1][j-1], dp[i-1][j]), dp[i][j-1])
            }
        }

    }
    return dp[m][n]
}
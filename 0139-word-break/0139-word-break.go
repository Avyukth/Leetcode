func wordBreak(s string, wordDict []string) bool {
    n:= len(s)
    wordMap := make(map[string]bool)
    for _, word := range wordDict{
        wordMap [word] = true
    }
    dp := make([]bool, n+1)
    dp[0] = true

    for i := 1; i<n+1; i++{
        for j :=0 ; j<i ; j++{
            if dp[j] && wordMap[s[j:i]]{
                dp[i]=true
            } 
        }
    }
    return dp[n]
}
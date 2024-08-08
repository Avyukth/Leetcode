func wordBreak(s string, wordDict []string) bool {
    wordDic := make (map[string]bool)
    for _, word := range wordDict{
        wordDic[word]= true
    }
    dp:=make([]bool, len(s)+1)
    dp[0] =true 
    for i:=1 ; i <len(s)+1 ; i++{
        for j:=0 ; j <i ; j++{
            if dp[j] && wordDic[s[j:i]] {
                dp[i] = true
                break
            }
        }
    }
    return dp[len(s)]
}
func isAnagram(s string, t string) bool {
    if len(s) != len(t){
        return false
    }
    dp:= make(map[rune]int)

    for _, char := range s{
            dp[char]++
    }
    for _, char := range t{
        dp[char] --
        if dp[char] <0{
            return false
        }
    }
    for _, count := range dp{
        if count != 0 {
            return false
        }
    }
    return true
}
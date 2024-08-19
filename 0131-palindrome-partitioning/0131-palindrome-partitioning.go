func partition(s string) [][]string {
            var result [][]string
        var current []string
        backtrack(s, current, &result, 0)
        return result
}


func backtrack(s string, current []string, result *[][]string, start int ){
    if start>= len(s){
        temp := make([]string, len(current))
        copy(temp, current)
        *result = append(*result, temp)
        return
    }
    for end := start; end<len(s); end ++ {
        if isPalindrome(s,start,end){
            current = append(current, s[start:end+1])
            backtrack(s, current, result, end+1)
            current = current[:len(current)-1]
        }
    }
}

func isPalindrome(s string, st, ed int)bool{
    for st <ed{
        if s[st] != s[ed]{
            return false
        }
        st ++ 
        ed --
    }
    return true
}
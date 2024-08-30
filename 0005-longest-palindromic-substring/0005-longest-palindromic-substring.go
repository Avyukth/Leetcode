func longestPalindrome(s string) string {
    n:= len(s)
    if n==0 {
        return ""
    }
    st , end := 0,0
    for i:=0 ;i<n ; i++{
        len1 := expand(s,i,i)
        len2 := expand(s,i,i+1)
        lt := max(len1, len2)
        if lt > end-st{
            st = i - (lt-1)/2
            end = i + (lt)/2 
        }
    }
    return s[st:end+1]
}

func expand(s string, i, j int)int{
    n:= len(s)
    for i>=0 &&  j < n && s[i] == s[j]{
        i --
        j ++
    } 
    return j-i-1
}
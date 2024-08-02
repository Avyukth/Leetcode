func longestPalindrome(s string) string {
    if len(s) ==0{
        return ""
    }
    start, end := 0, 0
    for i:=0; i<len(s); i++{
        len1, left1, right1 := expand(s,i,i)
        len2, left2, right2 := expand(s,i,i+1)
        if len1> end-start{
            start, end = left1, right1
        }
        if len2 > end-start {
            start, end = left2, right2
        }
    }
    return s[start:end+1]
}

func  expand(s string, left, right int)(int, int, int){
    for left >=0 && right < len(s) && s[left] ==s[right]{
        left--
        right ++
    }
    return right - left-1, left+1, right-1
}
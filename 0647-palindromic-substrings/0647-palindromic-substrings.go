func countSubstrings(s string) int {
    n:=len(s)
    count := 0 
    for i :=0 ; i <n ; i++{
        count += expand(s,i, i)
        count += expand(s,i, i+1)

    }
    return count
}

func expand(s string, i, j int)int{
    n:= len(s)
    count := 0
    for i >=0 && j < n && s[i] ==s[j]{
        count ++
        i --
        j++
    }
    return count
}
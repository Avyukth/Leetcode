func countSubstrings(s string) int {
    n:= len(s)
    count :=0

    for i :=0 ; i<n; i++{
        count += expandC(s, i, i)
        count += expandC(s, i, i+1)

    }
    return count

}

func expandC(s string, i, j int) int{
    count := 0
    for i>=0 && j <len(s) && s[i] == s[j]{
        count++
        i--
        j++
    }
    return count
}
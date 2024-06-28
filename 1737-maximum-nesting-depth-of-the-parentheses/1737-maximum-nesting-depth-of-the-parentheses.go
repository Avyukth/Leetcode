func maxDepth(s string) int {
    maxDepth := 0
    currD := 0
    for _, char := range s{
        if char == '('{
            currD ++
            if currD > maxDepth {
                maxDepth = currD
            }
        }else if  char == ')'{
            currD --
        }
    }
    return maxDepth
}
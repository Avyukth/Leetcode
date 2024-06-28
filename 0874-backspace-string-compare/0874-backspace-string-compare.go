func backspaceCompare(s string, t string) bool {
    return processS(s) == processS(t)
}

func processS(s string) string {
    var stack []rune
    for _, char := range s{
        if char == '#' {
            if len(stack)>0 {
                stack =  stack[:len(stack)-1]
            }
        }else{
            stack = append(stack, char)
        }
    }
    return string(stack)
}
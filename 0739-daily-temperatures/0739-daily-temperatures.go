func dailyTemperatures(t []int) []int {
    n:=len(t)
    ans := make([]int, n)
    stack := make([]int, 0, n)

    for  i :=0 ; i<n ; i++{
        for len(stack)>0  && t[i] > t[stack[len(stack)-1]] {
            prevI:= stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            ans[prevI] = i - prevI
        }
        stack = append(stack, i)
    }
    return ans
}
func totalSteps(nums []int) int {
    stack := make([][2]int, 0, len(nums))
    maxS:=0
    for i:=len(nums)-1; i>=0; i--{
        steps := 0
        for  len(stack)>0 && nums[i] > stack[len(stack)-1][0] {
            steps = max(steps+1,stack[len(stack)-1][1] )
            stack = stack[:len(stack)-1]
        }
        maxS = max(maxS, steps)
        stack = append(stack ,[2]int{nums[i], steps})
    }
    return maxS
}

func max(a,b int)int{
    if a>b {
        return a
    }
    return b
}

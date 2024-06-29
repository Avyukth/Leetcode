func nextGreaterElements(nums []int) []int {
    n:= len(nums)
    result :=make([]int, n)
    for i :=range result{
        result[i] =-1
    }

    stack := make([]int, 0, n)
    for i:=0; i<2*n; i++{
        index := i%n
        for len(stack)>0 && nums[stack[len(stack)-1]]<nums[index]{
            result[stack[len(stack)-1]] = nums[index]
            stack = stack[:len(stack)-1]
        }
        if i<n{
            stack = append(stack, index)
        }
    }
    return result
}
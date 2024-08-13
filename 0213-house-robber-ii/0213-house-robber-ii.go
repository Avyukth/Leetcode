func rob(nums []int) int {
    n:= len(nums)
    if n == 0 {
        return 0
    }
    if n==1{
        return nums[0]
    }
    val1 := robLinear(nums[1:n])
    val2 := robLinear(nums[0:n-1])
    return max(val1, val2)

}

func robLinear(nums []int) int {
    prev1 := 0
    prev2 := 0
    for _, num := range nums{
        temp :=  prev1
        prev1 = max(prev1,prev2+num )
        prev2 =  temp
    } 
    return max(prev1, prev2)
}
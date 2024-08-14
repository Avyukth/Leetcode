func containsDuplicate(nums []int) bool {
    n:= len(nums)
    ma := make(map[int]bool, n)
    for _, val := range nums{
        ma[val] = true
    }
    return len(ma) != n
}
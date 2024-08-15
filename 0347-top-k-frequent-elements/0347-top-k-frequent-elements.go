func topKFrequent(nums []int, k int) []int {
    dp := make(map[int]int)
    for _, val:= range nums{
        dp[val]++
    }
    var countArr [][]int
    for i, num := range dp{
        countArr= append(countArr, []int{i, num})
    }
    sort.Slice(countArr, func(i, j int)bool{
        return countArr[i][1]> countArr[j][1]
    })
    var res []int
    for i :=0; i <k; i++{
        res = append(res,countArr[i][0])
    }
    return res

}
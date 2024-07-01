func numFactoredBinaryTrees(arr []int) int {
    const MOD = 1000000007
    sort.Ints(arr)
    dp := make(map[int]int)
    for i, num := range arr{
        dp[num] = 1
        for j :=0 ; j< i; j++{
            if num % arr[j] == 0{
                right:= num/arr[j]
                if count, exist := dp[right];exist{
                    dp[num] = (dp[num]+dp[arr[j]]*count) %MOD
                }
            }
        }
    }
    total := 0
    for _, count := range dp{
        total = (total+count) % MOD
    }
    return total
}
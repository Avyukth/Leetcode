func countSubMultisets(nums []int, l int, r int) int {
    const MOD = 1000000007
    const N = 20000
    dp := make([]int, N+1)
	dp[0] = 1
	m := make(map[int]int)
	m[0] = 1
	for _, x := range nums {
		m[x]++
	}
	mult := m[0]
	delete(m, 0)
	for a, c := range m {
		help := make([]int, N+1)
		for i := 1; i <= N; i++ {
			sum := 0
			if i-a >= 0 {
				sum = (sum + dp[i-a]) % MOD
				sum = (sum + help[i-a]) % MOD
			}
			if i-a*(c+1) >= 0 {
				sum = (sum - dp[i-a*(c+1)] + MOD) % MOD
			}
			help[i] = sum
		}
		for i := 0; i <= N; i++ {
			dp[i] = (dp[i] + help[i]) % MOD
		}
	}
	sum := 0
	for i := l; i <= r; i++ {
		sum = (sum + dp[i]) % MOD
	}
	return (sum * mult) % MOD
}
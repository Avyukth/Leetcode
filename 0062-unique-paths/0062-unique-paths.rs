impl Solution {
    pub fn unique_paths(m: i32, n: i32) -> i32 {
        let (m,n) = (m as usize, n as usize);
        let mut dp  = vec![vec![0; n]; m];
    
        for i in 0..m {
            dp[i][0]=1;
        }
        for i in 0..n{
            dp[0][i]= 1;
        }
        for i in 1..m{
            for j in 1..n{
                dp[i][j]= dp[i-1][j] + dp[i][j-1];
            }
        }
        dp[m-1][n-1] as i32
    }
}
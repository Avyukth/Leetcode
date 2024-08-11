impl Solution {
    pub fn num_distinct(s: String, t: String) -> i32 {
        let  s : Vec<char> = s.chars().collect();
        let  t : Vec<char> = t.chars().collect(); 
        let (m,n) = (s.len(), t.len());
        let mut dp = vec![vec![0;n+1];m+1];
        for i in 0..=m{
            dp[i][0]= 1;
        }
        for i in  1..=m{
            for j in 1..=n{
                if s[i-1]== t[j-1]{
                    dp[i][j] =  dp[i-1][j-1] + dp[i-1][j];  
                } else{
                    dp[i][j] = dp[i-1][j]; 
                }
            } 
        }
        dp[m][n] as i32
    }
}
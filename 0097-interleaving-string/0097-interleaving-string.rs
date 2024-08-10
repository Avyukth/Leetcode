impl Solution {
    pub fn is_interleave(s1: String, s2: String, s3: String) -> bool {
        let (m,n,l) = (s1.len(),s2.len(),s3.len());
        if m+n !=l{
            return false;
        }
        let (s1, s2, s3) = (s1.as_bytes(), s2.as_bytes(), s3.as_bytes());
        let mut dp = vec![false;n+1];
        for i  in 0..=m{
            for j  in 0..=n{
                if i ==0 && j==0 {
                    dp[j]= true;
                } else if i == 0 {
                    dp[j]= dp[j-1] && s2[j-1] == s3[j-1]
                } else if j == 0 {
                    dp[j]= dp[j] && s1[i-1] == s3[i-1]
                }else{
                    dp[j] = (dp[j]  && s1[i-1] == s3[i+j-1]) || (dp[j-1]  && s2[j-1] == s3[i+j-1]) ;
                }

            }

        }
        return dp[n]
    }
}
impl Solution {
    pub fn length_of_lis(nums: Vec<i32>) -> i32 {
        if nums.is_empty(){
            return 0;
        }
        let n = nums.len();
        let mut dp = vec![1;n];
        let mut max_lt =1;
        for i in 1..n{
            for j in 0..i{
                if nums[i]> nums[j]{
                    dp[i] = dp[i].max(1+dp[j]);
                }
            }
            max_lt = max_lt.max(dp[i]);
        }
        max_lt
    }
}
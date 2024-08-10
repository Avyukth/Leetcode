impl Solution {
    pub fn find_target_sum_ways(nums: Vec<i32>, target: i32) -> i32 {
        let total : i32 = nums.iter().sum();
        if total < target.abs() || (total + target)%2 != 0 {
            return 0;
        }
        let subset_sum  = (total+target)/2 ;
        let mut dp = vec![0; (subset_sum+1) as usize];
        dp[0] = 1;
        for num in nums{
            for i in (num as usize..= subset_sum as usize).rev(){
                dp[i]+= dp[i-num as usize];
            }
        }
        dp[subset_sum as usize]
    }
}
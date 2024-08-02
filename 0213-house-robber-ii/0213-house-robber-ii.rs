impl Solution {
    pub fn rob(nums: Vec<i32>) -> i32 {
        if nums.len() ==1{
            return nums[0];
        }
        let n = nums.len();
        let max1 =  Solution::rob_linear(&nums[0..n-1]);
        let max2 =  Solution::rob_linear(&nums[1..n]);
        max1.max(max2)

    }
    pub fn rob_linear(nums: &[i32]) -> i32{
        let mut prev1 =0;
        let mut prev2 =0;
        
        for &val in nums{
            let temp = prev1;
            prev1 = prev1.max(prev2 + val);
            prev2 = temp;
        }
        prev1
    }
}
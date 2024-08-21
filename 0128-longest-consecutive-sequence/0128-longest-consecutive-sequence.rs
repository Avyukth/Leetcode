
use std::collections::HashSet;
impl Solution {
    pub fn longest_consecutive(nums: Vec<i32>) -> i32 {
        let mut curr_num =0;
        let mut curr_st =0;
        let mut long_st =0;
        let num_set : HashSet<i32> = nums.into_iter().collect();
        for &num in &num_set{
            if !num_set.contains(&(num-1)){
                curr_num = num;
                curr_st =1;
                while num_set.contains(&(curr_num+1)){
                    curr_num += 1;
                    curr_st += 1;
                }
                long_st = long_st.max(curr_st);
            }

        }
        long_st

    }
}
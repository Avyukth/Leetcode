impl Solution {
    pub fn missing_number(nums: Vec<i32>) -> i32 {
        let mut result:i32 = nums.len() as i32;
        for i in 0..nums.len(){
            result ^= i as i32^nums[i as usize];
        }
        result
    }
}
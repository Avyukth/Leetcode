impl Solution {
    pub fn subsets_with_dup(nums: Vec<i32>) -> Vec<Vec<i32>> {
        let mut nums = nums;
        nums.sort(); 
        let mut result = Vec::new();
        let mut current = Vec::new(); 
        Self::backtrack(&nums , &mut current, &mut result, 0);
        result
    }
    fn backtrack(nums : &[i32], current: &mut Vec<i32>, result : &mut Vec<Vec<i32>>, start: usize){
        result.push(current.clone());
        for i in start..nums.len(){
            if i >start && nums[i] == nums[i-1]{
                continue;
            }
            current.push(nums[i]);
            Self::backtrack(nums , current, result, i+1);
            current.pop();
        }
    }

}
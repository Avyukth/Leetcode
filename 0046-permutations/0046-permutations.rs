impl Solution {
    pub fn permute(nums: Vec<i32>) -> Vec<Vec<i32>> {
        let mut current = Vec::new();
        let mut result = Vec::new();
        let mut nums = nums;
        nums.sort_unstable();
        let mut used = vec![false; nums.len()];
        Self::backtrack(&nums, &mut current, &mut result, &mut used);
        result
    }
    fn backtrack(nums : &[i32], current : &mut Vec<i32>, result : &mut Vec<Vec<i32>>, used :&mut[bool]){
        if current.len() == nums.len(){
            result.push(current.clone());
            return ;
        }
        for i in 0..nums.len(){
            if used[i] {
                continue;
            }
            used[i] = true;
            current.push(nums[i]);
            Self::backtrack(nums,  current, result, used);
            used[i]= false;
            current.pop();
        }


    }
}
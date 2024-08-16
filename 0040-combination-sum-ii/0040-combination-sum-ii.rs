impl Solution {
    pub fn combination_sum2(candidates: Vec<i32>, target: i32) -> Vec<Vec<i32>> {
        let mut candidates = candidates.clone();
        candidates.sort_unstable();
        let mut current = Vec::new();
        let mut result = Vec::new();
        Self::backtrack(&candidates, target, 0, &mut current, &mut result);
        result
    }
    fn backtrack(candidates: &[i32], target: i32, start: usize, current : &mut Vec<i32>, result : &mut Vec<Vec<i32>>){
        if target == 0 {
            result.push(current.clone());
            return ;
        }
        for i in start..candidates.len(){
            if i> start &&  candidates[i-1] == candidates[i]{
                continue;
            }
            if candidates[i]> target{
                break;
            }
            current.push(candidates[i]);
            Self::backtrack(candidates, target - candidates[i], i+1,  current,  result);
            current.pop();
        }
    }

}
impl Solution {
    pub fn search_matrix(matrix: Vec<Vec<i32>>, target: i32) -> bool {
        if matrix.is_empty() || matrix[0].is_empty(){
            return false;
        }
        let m = matrix.len() as i32;
        let n = matrix[0].len() as i32;
        let mut left = 0;
        let mut right = m*n-1;
        while left <= right{
            let mid = left + (right-left)/2;
            let mut row = (mid/n) as usize;
            let mut col = (mid%n) as usize;
            if matrix[row][col] ==target{
                return true;
            }else if  matrix[row][col]< target{
                left= mid+1;
            }else{
                right = mid-1;
            }
        } 
        false

    }
    
}
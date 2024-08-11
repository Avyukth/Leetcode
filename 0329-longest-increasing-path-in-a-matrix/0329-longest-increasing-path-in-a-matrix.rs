impl Solution {
    pub fn longest_increasing_path(matrix: Vec<Vec<i32>>) -> i32 {
        if matrix.is_empty() || matrix[0].is_empty(){
            return 0;
        } 
        let (m, n) = (matrix.len(), matrix[0].len());
        let mut memo = vec![vec![0;n];m];
        let mut max_path = 0 ;
        for i in 0..m{
            for j in 0..n{
                max_path = max_path.max(Self::dfs(&matrix, &mut memo, i, j, m,n));
            }
        }
        max_path
    }

     fn dfs(matrix :&Vec<Vec<i32>>,  memo : &mut Vec<Vec<i32>>,  i: usize, j :usize, m: usize, n :usize)->i32{
        if memo[i][j]!= 0{
            return memo[i][j];
        }
        let directions = [(0,1), (1,0), (0,-1),(-1,0)];
        let mut maxLt =1;
        for (di, dj) in directions.iter(){
            let ni = i as i32 + di;
            let nj = j as i32 + dj;
            if ni>=0  && ni <m as i32 && nj >=0  && nj <n as i32{
                let (ni, nj ) = (ni as usize,nj as usize  );
                if matrix[ni][nj] >matrix[i][j]{
                    maxLt =  maxLt.max(1+ Self::dfs(matrix, memo, ni, nj, m, n));
                }
            }

        }
        memo[i][j] = maxLt;
        maxLt
    }
}
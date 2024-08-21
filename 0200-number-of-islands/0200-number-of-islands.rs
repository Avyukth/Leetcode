impl Solution {
    pub fn num_islands(grid: Vec<Vec<char>>) -> i32 {
        let mut grid = grid;
        let mut count =0;
        let row = grid.len();
        let col = grid[0].len();
        for i in 0..row{
            for j in 0..col{
                if grid[i][j] =='1' {
                    Self::dfs(&mut grid, i,j, row, col);
                    count +=1;
                }
            }
        } 
        count
    }

    fn dfs(grid :&mut Vec<Vec<char>>, i: usize, j: usize, row: usize, col: usize){
        if grid[i][j] !='1'{
            return ;
        }
        grid[i][j] = '0';
        let directions: [(i32, i32);4] = [(-1, 0), (0, 1), (1, 0), (0, -1)];
        for (di, dj) in directions.iter(){
            let ni = i as i32 + di;
            let nj = j as i32 + dj;
            if ni >=0  && ni<row as i32 && nj >=0  && nj<col as i32{
                Self::dfs( grid, ni as usize,nj as usize, row, col)
            }

        }

    }
}
impl Solution {
    pub fn exist(board: Vec<Vec<char>>, word: String) -> bool {
        let m = board.len();
        let n = board[0].len();
        let word  = word.chars().collect::<Vec<char>>();
        let mut visited  =vec![vec![false;n];m];
        for i in 0..m{
            for j in 0..n{
                if Self::dfs(&board, &word, &mut visited,i,j,0){
                    return true
                }
            }

        }
        false

    }
    fn dfs(board :&Vec<Vec<char>>, word :&Vec<char> , visited: &mut Vec<Vec<bool>> ,i :usize,j :usize,k :usize)->bool{
        if k==word.len(){
            return true;
        }
        if i>= board.len() || i<0 || j>= board[0].len() || j<0 || visited [i][j] ||  board[i][j] != word[k]{
            return false;
        }
        visited [i][j] = true;
        let directions = [(0, 1), (1, 0), (0, -1), (-1, 0)];
        for (di, dj) in directions.iter(){
            let ni = di + i as i32;
            let nj = dj + j as i32;
            if ni>=0 && nj >=0{
                let ni = ni as usize;
                let nj = nj as usize;

                if Self::dfs( board, word, visited,ni,nj,k+1){
                    return true;
                }
            }

        }
        visited [i][j] = false;
        false
    }
}
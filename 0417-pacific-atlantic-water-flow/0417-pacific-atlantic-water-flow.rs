use std::collections::VecDeque;
impl Solution {
    pub fn pacific_atlantic(heights: Vec<Vec<i32>>) -> Vec<Vec<i32>> {
        let (m,n) = match (heights.len(), heights.first().map(|row| row.len())){
            (0, _)|(_, Some(0))| (_, None) => return Vec::new(),(m, Some(n))=>(m,n),
        };
        let mut pacific = vec![vec![false;n];m];
        let mut atlantic = vec![vec![false;n];m];
        let mut queue = VecDeque::with_capacity(m.max(n));

        for i in 0..m{
            queue.push_back((i,0));
        }
        

        for j in 0..n{
            queue.push_back((0,j));
        }

        Self::bfs(&heights, &mut pacific, &mut queue);
        
        for i in 0..m{
            queue.push_back((i,n-1));
        }
        

        for j in 0..n{
            queue.push_back((m-1,j));
        }
        Self::bfs(&heights, &mut atlantic, &mut queue);

        (0..m)
            .flat_map(|i| (0..n).map(move |j| (i, j)))
            .filter(|&(i, j)| pacific[i][j] && atlantic[i][j])
            .map(|(i, j)| vec![i as i32, j as i32])
            .collect()
    }

    fn bfs (heights: &[Vec<i32>], reachable: &mut [Vec<bool>], queue: &mut VecDeque<(usize, usize)>){
        let (m, n) = (heights.len(), heights[0].len());
        let directions = [(0, 1), (1, 0), (0, -1), (-1, 0)];

        while let Some((i, j)) = queue.pop_front() {
            reachable[i][j] = true;

            for (di, dj) in directions.iter() {
                let ni = i as i32 + di;
                let nj = j as i32 + dj;

                if ni >= 0 && ni < m as i32 && nj >= 0 && nj < n as i32 {
                    let (ni, nj) = (ni as usize, nj as usize);
                    if !reachable[ni][nj] && heights[ni][nj] >= heights[i][j] {
                        queue.push_back((ni, nj));
                    }
                }
            }
        }
    }
}
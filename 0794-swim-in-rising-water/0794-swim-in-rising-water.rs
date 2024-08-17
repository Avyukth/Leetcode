use std::collections::VecDeque;
impl Solution {
    pub fn swim_in_water(grid: Vec<Vec<i32>>) -> i32 {
    let n = grid.len();
    let  (mut left, mut right) = (grid[0][0],(n * n - 1) as i32);
    while left<right{
        let mid =left + (right-left)/2;
        if Self::can_reach_bottom_right(&grid, mid){
            right = mid;
        }else{
            left = mid+1 ;
        }
        }
        left
    }
    fn can_reach_bottom_right(grid: &Vec<Vec<i32>>, time: i32) -> bool {
        let n = grid.len();
        let mut visited = vec![vec![false; n]; n];
        let mut queue = VecDeque::new();
        let directions = [(0, 1), (1, 0), (0, -1), (-1, 0)];

        if grid[0][0] > time {
            return false;
        }

        queue.push_back((0, 0));
        visited[0][0] = true;

        while let Some((x, y)) = queue.pop_front() {
            if x == n - 1 && y == n - 1 {
                return true;
            }

            for (dx, dy) in &directions {
                let nx = x as i32 + dx;
                let ny = y as i32 + dy;

                if nx >= 0 && nx < n as i32 && ny >= 0 && ny < n as i32 {
                    let nx = nx as usize;
                    let ny = ny as usize;
                    if !visited[nx][ny] && grid[nx][ny] <= time {
                        visited[nx][ny] = true;
                        queue.push_back((nx, ny));
                    }
                }
            }
        }

        false
    }
}
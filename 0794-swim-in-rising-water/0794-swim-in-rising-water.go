func swimInWater(grid [][]int) int {
    n:= len(grid)
    left, right := grid[0][0] , n*n-1
    for left <right {
        mid := left + (right-left)/2
        if canReach(grid,n, mid){
            right = mid
        }else{
            left = mid+1
        }
    }
    return left
}

func canReach(grid [][]int ,n , t int)bool{
    visited := make([][]bool, n)
    for visit := range visited{
        visited[visit] =make([]bool, n)
    }
    if grid[0][0] > t{
        return false
    }
    que := list.New()
    que.PushBack([2]int{0, 0})
    dirs := [4][2]int {{0, -1}, {-1, 0}, {0, 1},{1, 0}}
    for que.Len()>0{
        curr := que.Remove(que.Front()).([2]int)
        if curr[0] == n-1 && curr[1] == n-1{
            return true
        }
        for _, dir:= range dirs{
            nx, ny :=  curr[0] +dir[0], curr[1] +dir[1]
            if nx>=0 &&  nx<n  && ny>=0 && ny<n && !visited[nx][ny] && grid[nx][ny]<=t{
                que.PushBack([2]int{nx, ny})
                visited[nx][ny] =true
            }
        }
    }
    return false
}
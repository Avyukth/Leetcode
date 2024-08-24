func maxAreaOfIsland(grid [][]int) int {
    m:= len(grid)
    n:= len(grid[0])
    if m==0 || n==0{
        return 0
    }
    maxC := 0
    for i:= 0; i<m ; i++{
        for j:= 0; j<n ; j++{
            if grid[i][j] ==1{
                maxC = max(maxC, dfs(grid , m,n , i, j))

            }
            

        }
    }
    return maxC


}


func dfs(grid [][]int, m,n , i, j int) int  {
    if i < 0 || i >= m || j < 0 || j >= n || grid[i][j] == 0 {
		return 0
	}
    grid[i][j] =0
    area :=1
    directions := [4][2]int{{0,1},{0,-1},{1,0},{-1,0}}
    for _, dir := range directions{
        ni, nj := i+ dir[0], j+ dir[1]
        area += dfs(grid , m,n , ni, nj)
    }
    return area
}
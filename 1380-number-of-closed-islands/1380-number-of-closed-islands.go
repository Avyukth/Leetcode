func closedIsland(grid [][]int) int {
    if len(grid) == 0 || len(grid[0]) == 0 {
        return 0
    }
    
    rows, cols := len(grid), len(grid[0])
    count := 0
    
    // First, mark all land cells connected to the border as visited
    for i := 0; i < rows; i++ {
        dfs(grid, i, 0, rows, cols)
        dfs(grid, i, cols-1, rows, cols)
    }
    for j := 0; j < cols; j++ {
        dfs(grid, 0, j, rows, cols)
        dfs(grid, rows-1, j, rows, cols)
    }
    
    // Now count the closed islands
    for i := 1; i < rows-1; i++ {
        for j := 1; j < cols-1; j++ {
            if grid[i][j] == 0 {
                dfs(grid, i, j, rows, cols)
                count++
            }
        }
    }
    
    return count
}

func dfs(grid [][]int, i, j, rows, cols int) {
    if i < 0 || i >= rows || j < 0 || j >= cols || grid[i][j] == 1 {
        return
    }
    
    // Mark as visited
    grid[i][j] = 1
    
    // Directions: up, down, left, right
    directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
    
    for _, dir := range directions {
        newI, newJ := i + dir[0], j + dir[1]
        dfs(grid, newI, newJ, rows, cols)
    }
}
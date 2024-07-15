func minPathCost(grid [][]int, moveCost [][]int) int {
    m:= len(grid)
    n:= len(grid[0])

    dp := make([][]int, m)
    for i:=0; i<m; i++{
        dp[i]= make([]int, n)
        for j:= range dp[i]{
            dp[i][j] = math.MaxInt32

        }
    }
    for i:=0; i<n; i++{
        dp[0][i] = grid[0][i]
    }

     for i:=0; i<m-1; i++{
         for j:=0; j<n; j++{
             for k:=0; k<n; k++{
                dp[i+1][k] = min(dp[i+1][k], dp[i][j]+grid[i+1][k]+moveCost[grid[i][j]][k])
             }
         }
     }
     result:= math.MaxInt32
      for i:=0; i<n; i++{
        result = min(result, dp[m-1][i])
      }
      return result

}

type Matrix struct {
    data [][]int
    memo [][]int
    m, n int
}

func longestIncreasingPath(matrix [][]int) int {
    if len(matrix) ==0 || len(matrix[0]) ==0 {
        return 0
    }
    m, n := len(matrix), len(matrix[0])
    memo := make([][]int, m)
    for i := 0 ; i<m; i++{
        memo[i] =make([]int,n)
    }
    matrixS := &Matrix{
        data : matrix,
        memo:memo,
        m:m,
        n:n,
    }
    maxLt := 0
    for i := 0 ; i<m; i++{
        for j := 0 ; j<n; j++{
            path := dfs(i,j, matrixS)
            if path > maxLt{
                maxLt = path
            }
        }
    }
    return maxLt

}

func dfs (i,j int, matrix *Matrix) int{
    if matrix.memo[i][j] !=0{
        return matrix.memo[i][j]
    }
    directions := [][]int{{0,1},{1,0},{0,-1},{-1,0}}
    maxLt := 1
    for _, dir:= range directions{
        ni, nj := i+dir[0], j+dir[1]
        if ni>=0 && ni <matrix.m && nj>=0 && nj < matrix.n && matrix.data[ni][nj] > matrix.data[i][j]{
            lt := 1+ dfs( ni, nj, matrix)
            if lt>maxLt{
                maxLt = lt
            }
        }
    }
    matrix.memo[i][j] =  maxLt 
    return maxLt
}
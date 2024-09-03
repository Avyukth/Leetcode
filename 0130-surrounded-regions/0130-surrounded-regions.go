var directions = [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

func solve(board [][]byte)  {
    if len(board)==0 ||len(board[0])==0{
        return
    }
    m, n := len(board), len(board[0])
    for i := 0; i < m; i++ {
        dfs(board, i, 0)
        dfs(board, i, n-1)
    }
    for j := 0; j < n; j++ {
        dfs(board, 0, j)
        dfs(board, m-1, j)
    }
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if board[i][j] == 'O' {
                board[i][j] = 'X'
            } else if board[i][j] == 'S' {
                board[i][j] = 'O'
            }
        }
    }

}

func dfs(board [][]byte, i, j int){
    m, n := len(board), len(board[0])
    if i <0 || i>=m || j<0 || j>= n || board[i][j] !='O'{
        return
    }
    board[i][j] = 'S'
    for _ , dir:= range directions{
        dfs(board,i+dir[0], j+dir[1])
    }
}
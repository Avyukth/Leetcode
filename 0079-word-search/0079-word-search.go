var directions = [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
func exist(board [][]byte, word string) bool {
    m, n := len(board), len(board[0])
    for i := 0 ; i<m; i++{
        for j := 0 ; j<n; j++{
            if dfs(board, word,  i, j,0 ){
                return true
            }
        }
    }
    return false
}

func dfs(board [][]byte, word string,i,j,k int)bool{
    if k == len(word){
        return true
    }
    if !isValid(board, i, j ) || board[i][j] != word[k]{
        return false
    }
    temp := board[i][j]
    board[i][j] ='#'

    for _, dir:= range directions{
        newI, newJ := i + dir[0], j + dir[1]
        if dfs(board, word, newI, newJ,k+1 ){
            board[i][j] = temp
            return true
        }
    }
    board[i][j] = temp
    return false  
}
func isValid(board [][]byte,i,j int)bool{
     m, n := len(board), len(board[0])
     return i>= 0 && i <m && j >=0 && j<n
}
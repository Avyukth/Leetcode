func findSmallestSetOfVertices(n int, edges [][]int) []int {
    inC := make([]int, n)
    for _, edge :=range edges{
        to:= edge[1]
        inC[to]++
    }

    var result []int
    for i:=0;i<n;i++{
        if inC[i]==0{
            result= append(result, i)
        }
    }
    return result
}
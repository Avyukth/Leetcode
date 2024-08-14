func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
    distances := make([]int, n)
    for i := range  distances{
        distances[i] = math.MaxInt32
    }
    distances[src] = 0
    for i := 0 ; i <=k; i++ {
        temp := make([]int, n)
        copy(temp,distances )
        for _,  fl := range flights{
            f, t , p :=  fl[0], fl[1] , fl[2]
            if distances[f] != math.MaxInt32{
                temp[t] = min(temp[t], distances[f]+p)
            }
        }
        distances =temp
    } 
    if distances[dst] == math.MaxInt32{
        return -1
    }  
    return distances[dst]
}
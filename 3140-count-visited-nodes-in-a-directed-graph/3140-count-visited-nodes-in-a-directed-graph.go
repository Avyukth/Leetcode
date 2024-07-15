func countVisitedNodes(edges []int) []int {
   sz := len(edges)
    answer := make([]int, sz)
    
    for indx := 0; indx < sz; indx++ {
        visited := make(map[int]bool)
        var nodeStack []int
        currSrc := indx
        
        for !visited[currSrc] && answer[currSrc] == 0 {
            visited[currSrc] = true
            nodeStack = append(nodeStack, currSrc)
            currSrc = edges[currSrc]
        }
        
        if visited[currSrc] {
            cycleLength := len(nodeStack) - indexOf(nodeStack, currSrc)
            for i := 0; i < cycleLength; i++ {
                answer[nodeStack[len(nodeStack)-1]] = cycleLength
                nodeStack = nodeStack[:len(nodeStack)-1]
            }
        }
        
        for len(nodeStack) > 0 {
            currSrc = nodeStack[len(nodeStack)-1]
            nodeStack = nodeStack[:len(nodeStack)-1]
            answer[currSrc] = answer[edges[currSrc]] + 1
        }
    }
    
    return answer
}

func indexOf(slice []int, item int) int {
    for i, v := range slice {
        if v == item {
            return i
        }
    }
    return -1
}
func canFinish(numCourses int, prerequisites [][]int) bool {
    graph := make([][]int, numCourses)
    for _, preR := range(prerequisites){
        course , preRqt := preR[0], preR[1]
        graph[course]=append(graph[course], preRqt)
    }
    fmt.Println("graph", graph)
    fmt.Println("prerequisites",prerequisites)

    visited := make([]int, numCourses)
    for course :=0; course<numCourses ; course ++{
        if !dfs(course, graph, visited){
            return false
        }
    }
    return true

}


func dfs(node int, graph[][]int, color []int)bool{
    if color[node]!=0{
        return color[node]==2
    }
    color[node]=1
    for _, nei := range graph[node]{
        if !dfs(nei, graph, color){
            return false
        }
    }
    color[node] =2
    return true

}
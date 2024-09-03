func findOrder(numCourses int, prerequisites [][]int) []int {
    graph := make([][]int,  numCourses)
    inDegree := make([]int,  numCourses)
    for _, preR := range prerequisites{
        course, pre := preR[0], preR[1]
        graph[pre] = append(graph[pre], course)
        inDegree[course]++
    }
    var queue []int
    for course := 0; course<numCourses; course++{
        if inDegree[course]==0{
            queue = append(queue, course)
        }
    }
    order := make([]int, 0, numCourses)

    for len(queue)>0{
        course := queue[0]
        queue = queue[1:]
        order = append(order, course)
        for _, nextC := range graph[course]{
            inDegree[nextC]--
            if inDegree[nextC] == 0 {
                queue = append(queue, nextC)
            }
        }
    }
    if len(order) == numCourses {
        return order
    }
    
    return []int{} 

}
func canVisitAllRooms(rooms [][]int) bool {
    n:= len(rooms)
    visited := make([]bool, n)
    visited[0]= true
    queue :=[]int{0}
    for len(queue)>0{
        currR  := queue[0]
        queue = queue[1:]

        for _, key:= range rooms[currR] {
            if !visited[key]{
                visited[key]= true
                queue = append(queue, key)
            }
        }
    }


    for _, v :=range visited{
        if !v{
            return false
        }
    }
    return true
}
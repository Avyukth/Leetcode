type Graph map[string][]string
func findItinerary(tickets [][]string) []string {
    graph := buildGraph(tickets)
    result := make([]string, 0, len(tickets)+1)
    dfs(graph, "JFK", &result)
    reverse(result)
    return result
}
func buildGraph(tickets [][]string)Graph{
    graph := make (Graph)
    for _, ticket := range  tickets{
        from, to := ticket[0], ticket[1]
        graph[from] =  append(graph[from], to)
    }
    for from := range graph{
        sort.Strings(graph[from])
    }
    return graph
}

func dfs(graph Graph, airport string, result *[]string){
    for len(graph[airport])>0{
        next := graph[airport][0]
        graph[airport] = graph[airport][1:]
        dfs(graph, next, result)
    }
    *result = append(*result, airport)
}
func reverse(s [] string){
    for i, j := 0, len(s)-1 ; i<j ; i, j = i+1,j-1{
        s[i], s[j] = s[j], s[i]
    } 
}
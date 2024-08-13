type UnionFind struct {
    parent  []int
    rank    []int
}
type Edge struct {
    weight, u, v int
    }

func NewUnionFind(n int)*UnionFind{
    parent := make([]int, n)
    rank := make([]int, n)
    for i := range n{
        parent[i]=i
    }
    return &UnionFind{
        parent: parent,
        rank : rank,
    }
}

func (uf *UnionFind) find(x int)int{
    if uf.parent[x] != x{
        uf.parent[x] = uf.find(uf.parent[x])
    }
    return uf.parent[x]
}

func (uf *UnionFind) union(x, y int)bool{
    px, py := uf.find(x), uf.find(y)
    if px == py {
        return false
    }
    if uf.rank[px] <uf.rank[py] {
        uf.parent[px] = py
    }else  if uf.rank[px] >uf.rank[py]{
        uf.parent[py] = px
    }else{
        uf.parent[py] = px
       uf.rank[px] ++
    }
    return true
}

func minCostConnectPoints(points [][]int) int {
    n:= len(points)
    edges := make([]Edge, 0, n*(n-1)/2)
    for i := 0 ; i<n ;i++{
        for j := 0 ; j<n ;j++{
            weight := int(math.Abs(float64(points[i][0]-points[j][0])) + math.Abs(float64(points[i][1]-points[j][1])))
            edges = append(edges, Edge{weight, i, j})
        }

    }
    sort.Slice(edges, func(i,j int)bool{
        return  edges[i].weight < edges[j].weight
    })
    uf := NewUnionFind(n)
    totalCost := 0
    edgeUsed := 0
    for _, edge := range edges{
        if uf.union(edge.u, edge.v ){
            totalCost += edge.weight
            edgeUsed++
            if edgeUsed ==n-1{
                break
            }
        }
    }
    return totalCost

}
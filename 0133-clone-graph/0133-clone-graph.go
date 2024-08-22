/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

func cloneGraph(node *Node) *Node {
    if node == nil{
        return nil
    }
    cloned:= make(map[int]*Node)
    return dfs(node, cloned)
}

func dfs(node *Node, cloned map[int]*Node) *Node{
    if clone, exists := cloned[node.Val];exists{
        return clone
    }

    clone := &Node{
        Val: node.Val,
        Neighbors : make([]*Node, 0, len(node.Neighbors)),
    }
    cloned[node.Val] = clone
    for _, nei := range node.Neighbors{
        clone.Neighbors = append(clone.Neighbors, dfs(nei, cloned))
    }
    return clone
}
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func zigzagLevelOrder(root *TreeNode) [][]int {
    if root == nil{
        return [][]int{}
    }
    result := [][]int{}
    q := []*TreeNode{root}
    l2R := true
    for len(q)>0{
        lSize:= len(q)
        lVal := make([]int, lSize)
        for i := 0; i <lSize; i++{
            node := q[0]
            q = q[1:]
            if l2R {
                lVal [i] = node.Val
            }else{
                lVal [lSize-i-1] = node.Val
            }
            if node.Left != nil{
                q = append(q, node.Left)
            }
            if node.Right != nil{
                q = append(q, node.Right)
            }
        }
        result = append(result, lVal)
        l2R = !l2R
    }
    return result
}
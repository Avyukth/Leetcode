/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func deepestLeavesSum(root *TreeNode) int {
    if root ==nil {
        return 0
    }
    q:= []*TreeNode{root}
    levelS:=0
    for len(q)>0{
        lSize :=len(q)
        levelS =0
        for i := 0; i<lSize; i++{
            node := q[0]
            q=q[1:]
            levelS += node.Val
            if node.Left != nil{
                q= append(q, node.Left)
            }
            if node.Right != nil{
                q= append(q, node.Right)
            }
        }
        if len(q)==0{
            return levelS
        }
    }
    return levelS



}
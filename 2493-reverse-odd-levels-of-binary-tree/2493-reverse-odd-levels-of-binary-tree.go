/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func reverseOddLevels(root *TreeNode) *TreeNode {
    if root ==nil{
        return nil
    }
    q:= []*TreeNode{root}
    level:=0

    for len(q)>0{
        levelS := len(q)
        levelN:=make([]*TreeNode, levelS)
        for i:=0 ; i < levelS; i++{
            node := q[0]
            q=q[1:]
            levelN[i] = node
            if node.Left !=  nil{
                q = append(q, node.Left)
            }
            if node.Right !=  nil{
                q = append(q, node.Right)
            }

        }
        if level%2==1{
            for i, j :=0, levelS-1 ; i<j; i,j = i+1 , j-1{
                levelN[i].Val , levelN[j].Val =levelN[j].Val ,levelN[i].Val
            } 
        }
        level++
    }
    return root
}
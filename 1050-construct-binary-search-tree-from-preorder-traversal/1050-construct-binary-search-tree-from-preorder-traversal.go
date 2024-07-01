/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func bstFromPreorder(preorder []int) *TreeNode {
    if len(preorder) == 0 {
        return nil
    }
    root := &TreeNode{Val:preorder[0]}
    for _, val := range preorder[1:]{
        insertBst(root,val)
    }
    return root
}

func insertBst(root *TreeNode, val int)*TreeNode {
    if root == nil{
        return &TreeNode{Val: val}
    }
    if root.Val<val{
        root.Right = insertBst(root.Right, val)
    }else{
        root.Left = insertBst(root.Left, val)
    }
    return root
}
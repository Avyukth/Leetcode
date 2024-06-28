/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func constructMaximumBinaryTree(nums []int) *TreeNode {
    if len(nums)==0{
        return nil
    }
    maxI := finMaxI(nums)
    root := &TreeNode{Val: nums[maxI]}
    root.Left = constructMaximumBinaryTree(nums[:maxI])
    root.Right = constructMaximumBinaryTree(nums[maxI+1:])
    return root

}

func finMaxI(nums []int) int{
    maxI := 0 
    for i, num := range nums {
        if num >nums[maxI]{
            maxI = i
        }
    }
    return maxI
}
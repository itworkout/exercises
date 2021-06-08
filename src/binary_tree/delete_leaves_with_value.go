//Given a binary tree root and an integer target, delete all the leaf nodes with value target.
//
//Note that once you delete a leaf node with value target,
//if it's parent node becomes a leaf node and has the value target,
//it should also be deleted (you need to continue doing that until you can't).
package binary_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func removeLeafNodes(root *TreeNode, target int) *TreeNode {
	if root.Left != nil {
		root.Left = removeLeafNodes(root.Left, target)
	}
	if root.Right != nil {
		root.Right = removeLeafNodes(root.Right, target)
	}

	if root.Left == nil && root.Right == nil && root.Val == target {
		return nil
	}
	return root
}

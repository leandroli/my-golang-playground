package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := new(TreeNode)
	root.Val = preorder[0]
	var pos int
	for inorder[pos] != root.Val {
		pos++
	}
	root.Left = buildTree(preorder[1:pos+1], inorder[0:pos])
	root.Right = buildTree(preorder[pos+1:], inorder[pos+1:])
	return root
}

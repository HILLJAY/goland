package main

import "work-place/src/leedCode"

func main() {

	//var root *leedCode.TreeNode
	//root = &leedCode.TreeNode{Val: 0}
	////node2 := &leedCode.TreeNode{Val: 1}
	////node3 := &leedCode.TreeNode{Val: 2}
	////node4 := &leedCode.TreeNode{Val: 3}
	//
	////root.Left = node2
	////root.Right = node3
	////node2.Left = node4
	////
	////arr := leedCode.InorderTraversal(root)
	////fmt.Println(arr)
	//
	//boo := leedCode.IsValidBST(root)
	//fmt.Println(boo)

	inorder := []int{9, 3, 15, 20, 7}
	postorder := []int{9, 15, 7, 20, 3}
	leedCode.BuildTree(inorder, postorder)
}

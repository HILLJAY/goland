package leedCode

import (
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func IsValidBST(root *TreeNode) bool {

	if root == nil {
		return true
	}

	stack := []*TreeNode{}
	inorder := math.MinInt64
	for len(stack) > 0 || root != nil {

		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if root.Val <= inorder {
			return false
		}

		inorder = root.Val
		root = root.Right
	}

	return true
}

func inorderReign(root *TreeNode, lower int, upper int) bool {

	if root == nil {

		return true
	}

	if root.Val < lower || root.Val > upper {

		return false
	}

	left := inorderReign(root.Left, lower, root.Val)
	right := inorderReign(root.Left, root.Val, upper)

	return left && right
}

//func inorderFix(root *TreeNode){
//
//	if root==nil {
//		return
//	}
//
//	inorderFix(root.Left)
//	arr = append(arr,root.Val)
//	inorderFix(root.Right)
//
//}

var treeMap map[int]int

func BuildTree(inorder []int, postorder []int) *TreeNode {

	treeMap = make(map[int]int)

	for i := 0; i < len(inorder); i++ {

		treeMap[inorder[i]] = i
	}

	return buildTreeWithIndex(inorder, postorder, 0, len(inorder)-1, 0, len(inorder)-1)
}

func buildTreeWithIndex(inorder []int, postorder []int, inorder_left int, inorder_right int, postorder_left int, postorder_right int) *TreeNode {

	if inorder_left > inorder_right || postorder_left > postorder_right {

		return nil
	}

	rootVal := postorder[postorder_right]
	inorderRootIndex := treeMap[rootVal]
	leftTreeLength := inorderRootIndex - inorder_left

	root := &TreeNode{

		Val: rootVal,
	}

	root.Left = buildTreeWithIndex(inorder, postorder, inorder_left, inorderRootIndex-1, postorder_left, postorder_left+leftTreeLength-1)

	root.Right = buildTreeWithIndex(inorder, postorder, inorderRootIndex+1, inorder_right, postorder_left+leftTreeLength, postorder_right-1)

	return root
}

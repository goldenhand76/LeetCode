package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func delNodes(root *TreeNode, to_delete []int) []*TreeNode {
	if root == nil {
		return []*TreeNode{}
	}

	delSet := make(map[int]struct{})
	for _, v := range to_delete {
		delSet[v] = struct{}{}
	}

	res := make([]*TreeNode, 0)
	var postorder func(root *TreeNode) *TreeNode
	postorder = func(root *TreeNode) *TreeNode {
		if root == nil {
			return nil
		}
		root.Left = postorder(root.Left)
		root.Right = postorder(root.Right)
		if _, ok := delSet[root.Val]; ok {
			if root.Left != nil {
				res = append(res, root.Left)
			}
			if root.Right != nil {
				res = append(res, root.Right)
			}
			return nil
		}
		return root
	}

	_ = postorder(root)
	if _, ok := delSet[root.Val]; !ok {
		res = append(res, root)
	}
	return res
}
func main() {
	// root := &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 5}}, Right: &TreeNode{Val: 3, Left: &TreeNode{Val: 6}, Right: &TreeNode{Val: 7}}}
	// nodes := []int{3, 5}
	// root := &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}, Right: &TreeNode{Val: 4}}
	// nodes := []int{3}
	// root := &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 3}}}
	// nodes := []int{2, 3}
	root := &TreeNode{Val: 1, Left: &TreeNode{Val: 5}, Right: &TreeNode{Val: 2, Right: &TreeNode{Val: 3, Left: &TreeNode{Val: 4}}}}
	nodes := []int{1, 5}
	for i, v := range delNodes(root, nodes) {
		println("Tree ", i, ":")
		printTree(v)
	}

}

func printTree(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Println(root.Val)
	printTree(root.Left)
	printTree(root.Right)
}

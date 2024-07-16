package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func createBinaryTree(descriptions [][]int) *TreeNode {
	descriptions = preprocess(descriptions)
	root := &TreeNode{Val: descriptions[0][0]}
	root = addChildren(root, descriptions[0])

	for i := 1; i < len(descriptions); i++ {
		if descriptions[i][0] == root.Val {
			root = addChildren(root, descriptions[i])
		} else if descriptions[i][1] == root.Val {
			oldRoot := root
			root = &TreeNode{Val: descriptions[i][0]}
			if descriptions[i][2] == 1 {
				root.Left = oldRoot
			} else {
				root.Right = oldRoot
			}
		} else {
			root, _ = find(root, descriptions[i])
		}
	}
	return root
}

func preprocess(descriptions [][]int) [][]int {
	unique := false
	for i := 0; i < len(descriptions); i++ {
		unique = true
		for j := 0; j < len(descriptions); j++ {
			if descriptions[i][0] == descriptions[j][1] {
				unique = false
				break
			}
		}
		if unique {
			descriptions[0], descriptions[i] = descriptions[i], descriptions[0]
			for o := i; i < len(descriptions); o++ {
				if descriptions[0][0] == descriptions[o][0] {
					descriptions[1], descriptions[o] = descriptions[o], descriptions[1]
					break
				}
			}
			break
		}
	}

	for i := 0; i < len(descriptions)-1; i++ {
		for j := i; j < len(descriptions); j++ {
			if descriptions[i][1] == descriptions[j][0] {
				descriptions[i+1], descriptions[j] = descriptions[j], descriptions[i+1]
				break
			}
		}
	}

	return descriptions
}

func find(root *TreeNode, node []int) (*TreeNode, error) {
	nodes := []*TreeNode{}
	nodes = append(nodes, root.Left, root.Right)
	for i := 0; i <= len(node); i++ {
		if nodes[i] != nil && node[0] == nodes[i].Val {
			nodes[i] = addChildren(nodes[i], node)
			return root, nil
		} else if nodes[i] != nil {
			nodes = append(nodes, nodes[i].Left, nodes[i].Right)
		}
	}
	return root, fmt.Errorf("cant find the node")
}

func addChildren(root *TreeNode, node []int) *TreeNode {
	if node[2] == 1 {
		root.Left = &TreeNode{Val: node[1]}
	} else {
		root.Right = &TreeNode{Val: node[1]}
	}
	return root
}

func main() {
	result := &TreeNode{}
	// result = createBinaryTree([][]int{{20, 15, 1}, {20, 17, 0}, {50, 20, 1}, {50, 80, 0}, {80, 19, 1}})
	// fmt.Print(result)
	// result = createBinaryTree([][]int{{1, 2, 1}, {2, 3, 0}, {3, 4, 1}})
	// fmt.Print(result)
	result = createBinaryTree([][]int{{85, 82, 1}, {74, 85, 1}, {39, 70, 0}, {82, 38, 1}, {74, 39, 0}, {39, 13, 1}})
	printTree(result)
}

func printTree(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Println(root.Val)
	printTree(root.Left)
	printTree(root.Right)
}

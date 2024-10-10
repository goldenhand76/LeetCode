package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 55}
	root.Right.Left = &TreeNode{Val: 12}
	root.Right.Right = &TreeNode{Val: 56}
	root.Right.Left.Left = &TreeNode{Val: 9}
	root.Left = &TreeNode{Val: 66}
	root.Left.Left = &TreeNode{Val: 97}
	root.Left.Left.Left = &TreeNode{Val: 54}
	root.Left.Right = &TreeNode{Val: 60}
	root.Left.Right.Left = &TreeNode{Val: 49}
	root.Left.Right.Left.Left = &TreeNode{Val: 90}
	countPairs(root, 5)
}

func countPairs(root *TreeNode, distance int) int {
	var dfs func(*TreeNode) []int
	count := 0

	dfs = func(node *TreeNode) []int {
		if node == nil {
			return nil
		}
		if node.Left == nil && node.Right == nil {
			return []int{0}
		}

		leftDistances := dfs(node.Left)
		rightDistances := dfs(node.Right)

		for _, ld := range leftDistances {
			for _, rd := range rightDistances {
				if ld+rd+2 <= distance {
					count++
				}
			}
		}

		newDistances := []int{}
		for _, ld := range leftDistances {
			if ld+1 < distance {
				newDistances = append(newDistances, ld+1)
			}
		}
		for _, rd := range rightDistances {
			if rd+1 < distance {
				newDistances = append(newDistances, rd+1)
			}
		}
		return newDistances
	}

	dfs(root)
	fmt.Println(count)
	return count
}

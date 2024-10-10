package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getDirections(root *TreeNode, startValue int, destValue int) string {
	result := ""
	var path1 string = ""
	var path2 string = ""
	var ancestors1 []int = []int{root.Val}
	var ancestors2 []int = []int{root.Val}

	findAncestors(root, destValue, &ancestors1, &path1)
	fmt.Println(ancestors1)
	findAncestors(root, startValue, &ancestors2, &path2)
	fmt.Println(ancestors2)

	LCA := findLCA(ancestors1, ancestors2)
	fmt.Println(LCA)

	if LCA == ancestors1[len(ancestors1)-1] {
		for i := len(ancestors2) - 2; i >= 0; i-- {
			if ancestors2[i] == LCA {
				result += "U"
				break
			}
			result += "U"
		}
		return result
	} else if LCA == ancestors2[len(ancestors2)-1] {
		for i := len(ancestors1) - 2; i >= 0; i-- {
			if ancestors1[i] == LCA {
				result += (path1)[:len(path1)-i]
			}
		}
		return result
	}

	if len(ancestors1) == 1 {
		for i := len(ancestors2) - 2; i >= 0; i-- {
			result += "U"
		}
	} else if len(ancestors2) == 1 {
		result = path1
	} else {
		for i := len(ancestors2) - 2; i >= 0; i-- {
			if ancestors2[i] == LCA {
				result += "U"
				break
			}
			result += "U"
		}

		for i := len(ancestors1) - 2; i >= 0; i-- {
			if ancestors1[i] == LCA {
				result += (path1)[:len(path1)-i]
				break
			}
		}
	}

	return result
}

func findLCA(a1 []int, a2 []int) int {
	for i := len(a1) - 1; i >= 0; i-- {
		for j := len(a2) - 1; j >= 0; j-- {
			if a1[i] == a2[j] {
				return a1[i]
			}
		}
	}
	return -1
}

func findAncestors(root *TreeNode, val int, ancestors *[]int, path *string) bool {
	if root.Val == val {
		return true
	}

	if root.Left != nil {
		*path += "L"
		*ancestors = append(*ancestors, root.Left.Val)
		if findAncestors(root.Left, val, ancestors, path) {
			return true
		}
		*ancestors = (*ancestors)[:len(*ancestors)-1]
		*path = (*path)[:len(*path)-1] // backtrack
	}

	if root.Right != nil {
		*path += "R"
		*ancestors = append(*ancestors, root.Right.Val)
		if findAncestors(root.Right, val, ancestors, path) {
			return true
		}
		*ancestors = (*ancestors)[:len(*ancestors)-1]
		*path = (*path)[:len(*path)-1] // backtrack
	}

	return false
}

func main() {
	r := getDirections(&TreeNode{Val: 5, Left: &TreeNode{Val: 1, Left: &TreeNode{Val: 3}}, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 6}, Right: &TreeNode{Val: 4}}}, 3, 6)
	// r := getDirections(&TreeNode{Val: 1, Left: &TreeNode{Val: 2}}, 2, 1)

	fmt.Print(r)
}

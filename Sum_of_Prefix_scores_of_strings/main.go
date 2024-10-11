package main

import "fmt"

// Define the Trie Node structure
type Node struct {
	Children    map[rune]*Node
	PrefixCount int
}

func insert(root *Node, word string) {
	node := root
	for _, ch := range word {
		if node.Children == nil {
			node.Children = make(map[rune]*Node)
		}
		if _, exists := node.Children[ch]; !exists {
			node.Children[ch] = &Node{}
		}
		node = node.Children[ch]
		node.PrefixCount++ // Increment the prefix count at each node
	}
}

func getPrefixScore(root *Node, word string) int {
	node := root
	score := 0
	for _, ch := range word {
		if node.Children == nil {
			break
		}
		node = node.Children[ch]
		score += node.PrefixCount // Add the count for this prefix
	}
	return score
}

func sumPrefixScores(words []string) []int {
	root := &Node{}

	// Insert each word into the Trie
	for _, word := range words {
		insert(root, word)
	}

	// Calculate the prefix score for each word
	result := make([]int, len(words))
	for i, word := range words {
		result[i] = getPrefixScore(root, word)
	}
	return result
}

func main() {
	words := []string{"a", "ab", "abc", "cab"}
	result := sumPrefixScores(words)
	fmt.Println(result) // Output: [4, 3, 2, 1]
}

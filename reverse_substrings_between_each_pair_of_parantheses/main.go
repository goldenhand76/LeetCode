package main

import (
	"fmt"
	"strings"
)

// Function to reverse a string
func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// Function to reverse contents inside parentheses and remove the brackets
func reverseParentheses(s string) string {
	stack := []string{}
	for _, char := range s {
		if char == ')' {
			// Pop characters until '('
			tmp := ""
			for len(stack) > 0 && stack[len(stack)-1] != "(" {
				tmp = stack[len(stack)-1] + tmp
				stack = stack[:len(stack)-1]
			}
			// Pop the '('
			if len(stack) > 0 && stack[len(stack)-1] == "(" {
				stack = stack[:len(stack)-1]
			}
			// Reverse the temporary string and push back to stack
			stack = append(stack, reverse(tmp))
		} else {
			// Push current character to stack
			stack = append(stack, string(char))
		}
	}
	// Join the stack to form the final string
	return strings.Join(stack, "")
}

func main() {
	input := "(ed(et(oc))el)"
	result := reverseParentheses(input)
	fmt.Println(result) // Output: leetcode
}

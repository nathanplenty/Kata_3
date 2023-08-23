package main

import (
	"bytes"
	"fmt"
)

func isValid(s string) bool {
	// Initialize stack as placeholder for brackets
	type Stack struct {
		Chars []byte
	}
	stack := Stack{}
	// Initialize current character for current bracket
	var current byte = '0'
	// Get each character as i
	for i := 0; i < len(s); i++ {
		current = byte(s[i])
		// Check if character can be used
		switch {
		// If opener bracket, append it
		case true == bytes.ContainsAny([]byte(string(current)), "([{"):
			stack.Chars = append(stack.Chars, current)
		// If correct closer bracket, remove it and its opener
		case current-1 == stack.Chars[(len(stack.Chars))-1]:
			stack.Chars = stack.Chars[:(len(stack.Chars))-1]
		case current-2 == stack.Chars[(len(stack.Chars))-1]:
			stack.Chars = stack.Chars[:(len(stack.Chars))-1]
		default:
			return false
		}
	}
	// Check if stack is empty, no brackets left
	if len(stack.Chars) == 0 {
		return true
	}
	return false
}

func main() {
	fmt.Println("Is Valid?:", isValid("()"))
	fmt.Println("Is Valid?:", isValid("()[]{}"))
	fmt.Println("Is Valid?:", isValid("(])"))
	fmt.Println("Is Valid?:", isValid("([)]"))
	fmt.Println("Is Valid?:", isValid("{[]}"))
}

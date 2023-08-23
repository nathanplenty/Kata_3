package main

import (
	"bytes"
	"fmt"
)

func isValid(stringIn string) bool {
	// Initialize stack as placeholder for brackets
	type Stack struct {
		Chars []byte
	}
	stack := Stack{}
	// Initialize current character for current bracket
	var currentCharacter byte = '0'
	// Get each character as char
	for char := 0; char < len(stringIn); char++ {
		currentCharacter = byte(stringIn[char])
		// Check if character can be used
		switch {
		// If opener bracket, append it
		case true == bytes.ContainsAny([]byte(string(currentCharacter)), "([{"):
			stack.Chars = append(stack.Chars, currentCharacter)
		// If correct closer bracket, remove it and its opener
		case currentCharacter-1 == stack.Chars[(len(stack.Chars))-1]:
			stack.Chars = stack.Chars[:(len(stack.Chars))-1]
		case currentCharacter-2 == stack.Chars[(len(stack.Chars))-1]:
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

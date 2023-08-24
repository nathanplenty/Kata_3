package main

import (
	"fmt"
	"strings"
)

func isValid(stringIn string) bool {
	lengthOfStringIn := strings.Count(stringIn, "") - 1
	if lengthOfStringIn < 2 {
		return false
	}
	type Stack struct {
		Chars []rune
	}
	stack := Stack{}
	var currentCharacter rune
	for char := 0; char < lengthOfStringIn; char++ {
		currentCharacter = rune(stringIn[char])
		switch {
		case true == strings.ContainsAny((string(currentCharacter)), "([{"):
			stack.Chars = append(stack.Chars, currentCharacter)
		case currentCharacter-1 == stack.Chars[(len(stack.Chars))-1]:
			stack.Chars = stack.Chars[:(len(stack.Chars))-1]
		case currentCharacter-2 == stack.Chars[(len(stack.Chars))-1]:
			stack.Chars = stack.Chars[:(len(stack.Chars))-1]
		default:
			return false
		}
	}
	return len(stack.Chars) == 0
}

func main() {
	fmt.Println("Is Valid?:", isValid("()"))
	fmt.Println("Is Valid?:", isValid("()[]{}"))
	fmt.Println("Is Valid?:", isValid("(])"))
	fmt.Println("Is Valid?:", isValid("([)]"))
	fmt.Println("Is Valid?:", isValid("{[]}"))
}

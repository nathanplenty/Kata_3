package main

import (
	"bytes"
	"fmt"
)

func isValid(s string) bool {
	// 1. Get string -> (([]){}[])
	str := s
	// 1.1. Get len of string
	l := len(str)
	// 1.2. Check if string is len 2 or more
	// 1.F. return false -> kill function
	if l < 2 {
		return false
	}
	type Stack struct {
		Chars []interface{}
	}
	stack := Stack{}
	// 2.1. init current char -> current
	var current byte = '0'
	// 3. Get string[n] (n+1) -> current = string[n]
	for i := 0; i < l; i++ {
		current = byte(str[i])
		// 3.1. Check if opener bracket
		switch {
		case true == bytes.ContainsAny([]byte(string(current)), "([{"):
			// 3.T. Put string[n] on stack -> stack[n] = current -> GoTo 3
			stack.Chars = append(stack.Chars, current)
		default:
			// 3.F. Check if index 0
			if len(stack.Chars) == 0 {
				// 3.T.1. return false -> kill function
				return false
			}
			// 3.F.1. Check if current is its closer
			switch {
			// 3.T.2. Del stack[n] -> GoTo 3
			case current-1 == stack.Chars[(len(stack.Chars))-1]:
				stack.Chars = stack.Chars[:(len(stack.Chars))-1]
			case current-2 == stack.Chars[(len(stack.Chars))-1]:
				stack.Chars = stack.Chars[:(len(stack.Chars))-1]
			default:
				// 3.F.2. return false -> kill function
				return false
			}
		}
	}
	// 4. Check if stack is empty
	if len(stack.Chars) == 0 {
		// 4.T. return true -> kill function
		return true
	}
	// 4.F. return false -> kill function
	return false
}

func main() {
	fmt.Println("Is Valid?:", isValid("()"))
	fmt.Println("Is Valid?:", isValid("()[]{}"))
	fmt.Println("Is Valid?:", isValid("(])"))
	fmt.Println("Is Valid?:", isValid("([)]"))
	fmt.Println("Is Valid?:", isValid("{[]}"))
}

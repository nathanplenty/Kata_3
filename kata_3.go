package main

import (
	"fmt"
)

func isValid(s string) bool {
	b := false
	// Pseudo Code
	// Get string -> "test(number)index{}end"
	fmt.Println("Input:", s)
	// Search for open bracket -> "test'('number)index{}end"
	// Save bracket -> br = "("
	bl := false
	br1 := '0'
	i := 0
out1:
	for i = 0; i < len(s); i++ {
		switch {
		case s[i] == '(':
			bl = true
			br1 = '('
			break out1
		case s[i] == '{':
			bl = true
			br1 = '{'
			break out1
		case s[i] == '[':
			bl = true
			br1 = '['
			break out1
		default:
			continue
		}
	}
	// If no bracket found -> False
	if bl == false {
		return b
	}
	// Delete front of open bracket -> "(number)index{}end"
	m := i + 1
	fmt.Println("Bracket:", string(br1))
	for i = 0; i < m; i++ {
		s = s[1:]
	}
	fmt.Println("Trimed:", s)
	// Search for next bracket -> (number')'index{}end"
	bl = false
	br2 := '0'
	i = 0
out2:
	for i = 0; i < len(s); i++ {
		switch {
		case s[i] == ')':
			bl = true
			br2 = ')'
			break out2
		case s[i] == '}':
			bl = true
			br2 = '}'
			break out2
		case s[i] == ']':
			bl = true
			br2 = ']'
			break out2
		default:
			continue
		}
	}
	m = i + 1
	fmt.Println("Bracket:", string(br2))
	// Is bracket valid to saved bracket? -> True / False
	bl = false
	fmt.Println("Bracket-Brothers:", string(br1), string(br2))
	switch {
	case br1 == '(':
		if br1 == br2+1 {
			bl = true
			fmt.Println("()", bl)
		}
	case br1 == '{':
		if br1 == br2+2 {
			bl = true
			fmt.Println("{}", bl)
		}
	case br1 == '[':
		if br1 == br2+2 {
			bl = true
			fmt.Println("[]", bl)
		}
	default:
		//should never happen
		fmt.Println("WHY IS THIS?!")
	}
	b = bl
	// Delete front of closer bracket including bracket-> "index{}end"
	for i = 0; i < m; i++ {
		s = s[1:]
	}
	fmt.Println("Trimed:", s)
	// Repeat
	return b
}

func main() {
	b := false
	b = isValid("test(number)index{}end")
	fmt.Println("Is Valid:", b)
}

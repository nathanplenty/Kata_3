package main

import (
	"fmt"
)

func isValid(s string) bool {
	b := false
	bl := false
	br1 := '0'
	br2 := '0'
	i := 0
	// Pseudo Code
	// Get string -> "test(number)index{}end"
	fmt.Println("Input:", s)
	// Search for opener bracket -> "test'('number)index{}end"
	// Save bracket -> br = "("
	g1 := func() {
		bl = false
	out:
		for i = 0; i < len(s); i++ {
			switch {
			case s[i] == '(':
				bl = true
				br1 = '('
				break out
			case s[i] == '{':
				bl = true
				br1 = '{'
				break out
			case s[i] == '[':
				bl = true
				br1 = '['
				break out
			default:
				continue
			}
		}
	}
	g1()
	// If no opener bracket found -> False
	if bl == false {
		return false
	}

	// Delete front of open bracket -> "(number)index{}end"
	m := i + 1
	// fmt.Println("Bracket:", string(br1))
	for i = 0; i < m; i++ {
		s = s[1:]
	}
	// fmt.Println("Trimed:", s)
	// Search for next bracket -> (number')'index{}end"
	i = 0
	g2 := func() {
		bl = false
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
	}
	g2()
	// if no bracket found -> False
	if bl == false {
		return false
	}
	m = i + 1
	// fmt.Println("Bracket:", string(br2))
	// Is bracket valid to saved bracket? -> True / False
	bl = false
	// fmt.Println("Bracket-Brothers:", string(br1), string(br2))
	bb := func() {
		b = true
		switch {
		case br1 == '(':
			if br1 == br2-1 {
				bl = true
				// fmt.Println("()", bl)
			}
		case br1 == '{':
			if br1 == br2-2 {
				bl = true
				// fmt.Println("{}", bl)
			}
		case br1 == '[':
			if br1 == br2-2 {
				bl = true
				// fmt.Println("[]", bl)
			}
		default:
			//should never happen
			// fmt.Println("WHY IS THIS?!")
		}
	}
	bb()
	// Delete front of closer bracket including bracket-> "index{}end"
	for i = 0; i < m; i++ {
		s = s[1:]
	}
	// fmt.Println("Trimed:", s)
	// Repeat
	b = isValid(s)
	return b
}

func main() {
	fmt.Println(isValid("()"))
	fmt.Println(isValid("()[]{}"))
	fmt.Println(isValid("(]"))
	fmt.Println(isValid("([)]"))
	fmt.Println(isValid("{[]}"))
}

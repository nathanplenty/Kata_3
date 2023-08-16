package main

import "fmt"

func isValid(s string) bool {
	// stuff
	b := false
	if s == "Test" {
		b = true
		fmt.Println("Is Valid?", b)
	}
	return b
}

func main() {
	// stuff
	isValid("Test")
}

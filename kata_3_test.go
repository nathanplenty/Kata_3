package main

import (
	"fmt"
	"testing"
)

func TestIsValidYes(t *testing.T) {
	fmt.Print("True Test: ")
	s := "()"
	b := isValid(s)
	if b != true {
		t.Errorf("b = %v; want true", b)
	}
	fmt.Println(b)
}

func Test3IsValidNo(t *testing.T) {
	fmt.Print("False Test: ")
	s := "([)]"
	b := isValid(s)
	if b != false {
		t.Errorf("b = %v; want false", b)
	}
	fmt.Println(b)
}

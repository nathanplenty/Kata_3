package main

import (
	"testing"
)

func TestIsValidYes(t *testing.T) {
	s := "()"
	b := isValid(s)
	if b != true {
		t.Errorf("b = %v; want true", b)
	}
}

func Test3IsValidNo(t *testing.T) {
	s := "([)]"
	b := isValid(s)
	if b != false {
		t.Errorf("b = %v; want false", b)
	}
}

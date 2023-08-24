package main

import (
	"testing"
)

func TestIsValidYes(t *testing.T) {
	stringIn := "()"
	boolean := isValid(stringIn)
	if boolean != true {
		t.Errorf("b = %v; want true", boolean)
	}
}

func TestIsValidNo(t *testing.T) {
	stringIn := "([)]"
	boolean := isValid(stringIn)
	if boolean != false {
		t.Errorf("boolean = %v; want false", boolean)
	}
}

func TestIsValidTooSmall(t *testing.T) {
	stringIn := ""
	boolean := isValid(stringIn)
	if boolean != false {
		t.Errorf("boolean = %v; want false", boolean)
	}
}

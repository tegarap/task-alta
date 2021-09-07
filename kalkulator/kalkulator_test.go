package kalkulator

import (
	"testing"
)

func TestAdditon(t *testing.T) {
	if Additon(1, 2) != 3 {
		t.Error("Expected 1 (+) 2 to equal 3")
	}
	if Additon(-1, -2) != -3 {
		t.Error("Expected -1 (+) -2 to equal -3")
	}
}
func TestSubstract(t *testing.T) {
	if Substract(3, 2) != 1 {
		t.Error("Expected 3 (-) 2 to equal 1")
	}
	if Substract(-3, -2) != -1 {
		t.Error("Expected -3 (-) -2 to equal -1")
	}
}
func TestMult(t *testing.T) {
	if Mult(1, 2) != 3 {
		t.Error("Expected 1 (*) 2 to equal 2")
	}
	if Mult(-1, -2) != -3 {
		t.Error("Expected -1 (*) -2 to equal 2")
	}
}
func TestDiv(t *testing.T) {
	if Div(3, 1) != 3 {
		t.Error("Expected 3 (/) 1 to equal 3")
	}
	if Div(-3, -1) != 3 {
		t.Error("Expected -3 (/) -1 to equal 3")
	}
}
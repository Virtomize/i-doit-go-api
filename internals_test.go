package goidoit

import (
	"testing"
)

type inttest struct {
	n        int
	expected int
}

type constructortest struct {
	input    []string
	expected bool
}

func TestGetID(t *testing.T) {
	var tests = []inttest{
		{1, 1},
		{5, 5},
		{100, 100},
	}
	for _, pairs := range tests {
		var test = 0
		TResetID()
		for i := 1; i <= pairs.n; i++ {
			test = TgetID()
		}
		if test != pairs.expected {
			t.Errorf("Expected %v got %v\n", pairs.expected, test)
		}
	}
}

func TestDebugPrint(t *testing.T) {
}

func TestDebug(t *testing.T) {
	Debug(false)
	if TgetDebug() {
		t.Errorf("Expected false got true\n")
	}
	Debug(true)
	if !TgetDebug() {
		t.Errorf("Expected true got false\n")
	}
}

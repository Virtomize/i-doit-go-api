package goidoit

import (
	"testing"
)

func TestNewAPI(t *testing.T) {
	var tests = []constructortest{
		{[]string{"", "test"}, false},
		{[]string{"test", ""}, false},
		{[]string{"test", "test"}, true},
	}
	for i := range tests {
		_, err := NewAPI(tests[i].input[0], tests[i].input[1])
		if tests[i].expected {
			if err != nil {
				t.Errorf("Expected nil got %v\n", err)
			}
		} else {
			if err == nil {
				t.Errorf("Expected some error got %v\n", err)
			}
		}
	}
}

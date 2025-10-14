package main

import (
	"testing"
)

func TestLargestNormal(t *testing.T) {
	actual := Largest(1, 2, 3)
	expected := 3

	if actual != expected {
		t.Errorf("Expected 3 but got %v", actual)
	}

	actual = Largest(5, 6, 7)
	expected = 7

	if actual != expected {
		t.Errorf("Expected 3 but got %v", actual)
	}

}

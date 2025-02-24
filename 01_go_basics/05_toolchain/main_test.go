package main

import "testing"

func TestAdd(t *testing.T) {
	result := Add(1, 2)
	if result != 3 {
		t.Errorf("Add(1,2) = %d, want 3", result)
	}
}

func TestSubtract(t *testing.T) {
	result := Subtract(1, 2)
	if result != -1 {
		t.Errorf("Subtract(1,2) = %d, want -1", result)
	}
}

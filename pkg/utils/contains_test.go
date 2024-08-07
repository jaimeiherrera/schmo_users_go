package utils

import "testing"

func TestStringContains(t *testing.T) {
	slice := []string{"a", "b", "c"}
	if !StringContains(slice, "a") {
		t.Errorf("Expected slice to contain 'a'")
	}
	if StringContains(slice, "d") {
		t.Errorf("Expected slice to not contain 'd'")
	}
}

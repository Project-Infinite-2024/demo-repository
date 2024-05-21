package main

import "testing"

// Test for the Add function with an intentional error
func TestAdd(t *testing.T) {
    result := Add(2, 3)
    expected := 6  // Intentional error: 2 + 3 should be 5, not 6

    if result != expected {
        t.Errorf("Add(2, 3) = %d; want %d", result, expected)
    }
}

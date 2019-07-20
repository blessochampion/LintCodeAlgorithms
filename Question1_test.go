package LintCode

import "testing"

func TestAplusB(t *testing.T) {
	var a, b int
	t.Run("Add two numbers whose sum is less than two digits", func(t *testing.T) {
		a = 4
		b = 5
		result := AplusB(a, b)
		expected := 9
		if result != expected {
			t.Errorf("Expected the sum of %d and %d to be %d, got %d", a, b, expected, result)
		}
	})

	t.Run("Add two numbers whose sum is greater than 1 digit", func(t *testing.T) {
		a = 40
		b = 5
		result := AplusB(a, b)
		expected := 45
		if result != expected {
			t.Errorf("Expected the sum of %d and %d to be %d, got %d", a, b, expected, result)
		}
	})

	t.Run("Add two numbers whose sum is greater than 2 digit", func(t *testing.T) {
		a = 99
		b = 5
		result := AplusB(a, b)
		expected := 104
		if result != expected {
			t.Errorf("Expected the sum of %d and %d to be %d, got %d", a, b, expected, result)
		}
	})
}

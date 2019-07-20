package LintCode

import "testing"

func TestTrailingZeros(t *testing.T) {

	var n int64
	var expected int64
	var result int64
	t.Run("n with factorial less than 10", func(t *testing.T) {
		n = 1
		result = TrailingZeros(n)
		expected = 0
		if result != expected {
			t.Errorf("Expected %d got %d", expected, result)
		}
	})

	t.Run("n with factorial greater than 10 but does not end with zero", func(t *testing.T) {
		n = 3
		result = TrailingZeros(n)
		expected = 0
		if result != expected {
			t.Errorf("Expected %d got %d", expected, result)
		}
	})

	t.Run("n with factorial with 1 leading zero", func(t *testing.T) {
		n = 5
		result = TrailingZeros(n)
		expected = 1
		if result != expected {
			t.Errorf("Expected %d got %d", expected, result)
		}
	})

	t.Run("n with factorial with  2 leading zero", func(t *testing.T) {
		n = 105
		result = TrailingZeros(n)
		expected = 25
		if result != expected {
			t.Errorf("Expected %d got %d", expected, result)
		}
	})
}

package calc

import "testing"

func TestAddition(t *testing.T) {
	a, b := 1, 3
	result := Add(a, b)
	if result < 0 {
		t.Error("Value is negative")
	}
	if result != 7 {
		t.Errorf("%d + %d = %d, Value is not 7", a, b, result)
	}
}

func TestSubtraction(t *testing.T) {
	a, b := 2, 3
	result := Add(a, b)
	if result < 0 {
		t.Error("Value is negative")
	}
	if result != 1 {
		t.Errorf("%d + %d = %d, Value is not 1", a, b, result)
	}
}

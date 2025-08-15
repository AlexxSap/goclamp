package goclamp

import "testing"

func TestIntClamp(t *testing.T) {
	tests := []struct {
		descr    string
		min      int
		max      int
		value    int
		expected int
	}{
		{"simpleInto", 2, 7, 5, 5},
		{"simpleMin", 2, 7, 1, 2},
		{"simpleMax", 2, 7, 9, 7},
	}

	for _, test := range tests {
		t.Run(test.descr, func(t *testing.T) {
			act := clamp(test.min, test.max, test.value)
			if act != test.expected {
				t.Errorf("act: %d, exp:%d", act, test.expected)
			}
		})
	}
}

func TestFloatClamp(t *testing.T) {
	tests := []struct {
		descr    string
		min      float32
		max      float32
		value    float32
		expected float32
	}{
		{"simpleInto", 2.2, 7.7, 5.5, 5.5},
		{"simpleMin", 2.2, 7.7, 1.1, 2.2},
		{"simpleMax", 2.2, 7.7, 9.9, 7.7},
	}

	for _, test := range tests {
		t.Run(test.descr, func(t *testing.T) {
			act := clamp(test.min, test.max, test.value)
			if act != test.expected {
				t.Errorf("act: %v, exp:%v", act, test.expected)
			}
		})
	}
}

func TestIntWithFunc(t *testing.T) {
	actual := clampFunc(
		3,
		5,
		7,
		func(a, b int) bool {
			return true
		})
	expected := 3
	if actual != expected {
		t.Errorf("act: %v, exp:%v", actual, expected)
	}
}

func TestWithFuncOdd(t *testing.T) {
	actual := clampFunc(
		4,
		7,
		2,
		func(a, b int) bool {
			return a%b == 0
		})
	expected := 2
	if actual != expected {
		t.Errorf("act: %v, exp:%v", actual, expected)
	}
}

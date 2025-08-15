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

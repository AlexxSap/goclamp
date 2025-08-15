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

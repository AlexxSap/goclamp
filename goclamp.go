package goclamp

import (
	"cmp"
)

// clamp function that limits a value to a specified range [min, max]
func clamp[T cmp.Ordered](min, max, value T) T {
	if cmp.Less(value, min) {
		return min
	} else if cmp.Less(max, value) {
		return max
	}
	return value
}

// clampFunc function that limits a value to a specified range [min, max] with specified less function
func clampFunc[T any](min, max, value T, less func(a, b T) bool) T {
	if less(value, min) {
		return min
	} else if less(max, value) {
		return max
	}

	return value
}

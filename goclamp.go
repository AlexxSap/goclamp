package goclamp

import (
	"cmp"
)

// Lesser is an interface that defines function Less. Less function reports whether the a less than b
type Lesser interface {
	Less(a, b any) bool
}

// clamp function that limits a value to a specified range [min, max]
func clamp[T cmp.Ordered](min, max, value T) T {
	if cmp.Less(value, min) {
		return min
	} else if cmp.Less(max, value) {
		return max
	}
	return value
}

// clampFunc function that limits a value to a specified range [min, max] with specified less function (implementation of Lesser)
func clampFunc[T any](min, max, value T, less func(a, b T) bool) T {
	if less(value, min) {
		return min
	} else if less(max, value) {
		return max
	}

	return value
}

package goclamp

import (
	"cmp"
)

type Lesser interface {
	Less(a, b any) bool
}

func clamp[T cmp.Ordered](min, max, value T) T {
	if cmp.Less(value, min) {
		return min
	} else if cmp.Less(max, value) {
		return max
	}
	return value
}

func clampFunc[T any](min, max, value T, less func(a, b T) bool) T {
	if less(value, min) {
		return min
	} else if less(max, value) {
		return max
	}

	return value
}

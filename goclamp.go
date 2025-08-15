package goclamp

import (
	"cmp"
)

func clamp[T cmp.Ordered](min, max, value T) T {
	if cmp.Less(value, min) {
		return min
	} else if cmp.Less(max, value) {
		return max
	}
	return value
}

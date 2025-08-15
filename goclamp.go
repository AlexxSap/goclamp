package goclamp

func clamp(min, max, value int) int {
	if value < min {
		return min
	} else if value > max {
		return max
	}
	return value
}

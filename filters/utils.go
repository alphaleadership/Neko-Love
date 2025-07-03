package filters

// clamp restricts the integer value v to be within the range [min, max].
// If v is less than min, min is returned. If v is greater than max, max is returned.
// Otherwise, v is returned unchanged.
func clamp(v, min, max int) int {
	if v < min {
		return min
	}
	if v > max {
		return max
	}
	return v
}

// clamp8 limits the input integer v to the range [0, 255] and returns it as a uint8.
// If v is less than 0, it returns 0. If v is greater than 255, it returns 255.
// Otherwise, it returns v converted to uint8.
func clamp8(v int) uint8 {
	if v < 0 {
		return 0
	}
	if v > 255 {
		return 255
	}
	return uint8(v)
}

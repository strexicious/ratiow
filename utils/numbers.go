package utils

import "math"

func Clamp(x, min, max float64) float64 {
	return math.Min(max, math.Max(x, min))
}

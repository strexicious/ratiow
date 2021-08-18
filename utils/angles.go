package utils

import "math"

func Deg2Radians(degrees float64) float64 {
	return degrees * math.Pi / 180.0
}

package util

import "math"

const (
	resizeMax = 20
	resizeMin = 2
)

func Interpolate(value uint32) float64 {
	var max uint32 = 1000000
	var maxv float64 = resizeMin
	var minv float64 = resizeMax
	if value >= max {
		return maxv
	}

	minLog := math.Log10(float64(max))
	logFactor := 0 - minLog

	logValue := (math.Log10(float64(value)) - minLog) / logFactor

	linearFactor := (maxv - minv) / (1.0 - 0.0)
	res := minv + (linearFactor * (1.0 - logValue))
	return res
}

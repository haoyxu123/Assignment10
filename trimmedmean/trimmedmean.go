package trimmedmean

import (
	"sort"
)

func ComputeTrimmedMean(data []float64, lowerTrim, upperTrim float64) float64 {
	sort.Float64s(data)
	n := len(data)
	lowerCount := int(float64(n) * lowerTrim)
	upperCount := int(float64(n) * upperTrim)

	if upperCount == 0 {
		upperCount = lowerCount
	}

	trimmedData := data[lowerCount : n-upperCount]
	sum := 0.0
	for _, v := range trimmedData {
		sum += v
	}
	mean := sum / float64(len(trimmedData))
	return mean
}

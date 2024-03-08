package trimmedmean

import (
	"math/rand"
	"sort"
	"testing"
)

func TestTrimmedMean(t *testing.T) {
	rand.Seed(123)

	floatData := make([]float64, 100)
	for i := range floatData {
		floatData[i] = rand.Float64() * 100
	}

	intData := make([]float64, 100)
	for i := range intData {
		intData[i] = float64(rand.Intn(100))
	}

	trim := 0.05

	sort.Float64s(floatData)
	floatMean := ComputeTrimmedMean(floatData, trim, trim)

	sort.Float64s(intData)
	intMean := ComputeTrimmedMean(intData, trim, trim)

	t.Logf("Trimmed mean of float data: %f", floatMean)
	t.Logf("Trimmed mean of integer data: %f", intMean)
}

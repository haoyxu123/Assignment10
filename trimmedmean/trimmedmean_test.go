package trimmedmean

import (
	"reflect"
	"testing"
)

func TestComputeTrimmedMean(t *testing.T) {
	tests := []struct {
		name      string
		data      []float64
		lowerTrim float64
		upperTrim float64
		want      float64
	}{
		{"SymmetricTrim", []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 0.1, 0.1, 5.5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ComputeTrimmedMean(tt.data, tt.lowerTrim, tt.upperTrim); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ComputeTrimmedMean() = %v, want %v", got, tt.want)
			}
		})
	}
}

package main

import (
	"fmt"

	"example.com/mymodule/trimmedmean"
)

func main() {
	data := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	lowerTrim := 0.1
	upperTrim := 0.1
	mean := trimmedmean.ComputeTrimmedMean(data, lowerTrim, upperTrim)
	fmt.Println("Trimmed mean:", mean)
}

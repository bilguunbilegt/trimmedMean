package trimmedMean

import (
	"fmt"
	"sort"
)

func trimmedMean(numbers []float64, trim ...int) (float64, error) {
	if len(trim) == 0 || len(trim) > 2 {
		return 0, fmt.Errorf("invalid number of trimming arguments")
	}

	// Sort the numbers in ascending order
	sort.Float64s(numbers)

	// Determine the number of elements to trim from the lower and upper extremes of the sorted numbers array
	lowerTrim := trim[0]
	upperTrim := lowerTrim
	if len(trim) == 2 {
		upperTrim = trim[1]
	}

	// Calculate indices for slicing the array to remove the lower and upper extremes
	n := len(numbers)
	lowerIndex := lowerTrim
	upperIndex := n - upperTrim

	if lowerIndex >= upperIndex || lowerIndex < 0 || upperIndex > n {
		return 0, fmt.Errorf("invalid trimming values, resulting in no elements or out-of-bounds indices")
	}

	// Check if at least one element remains after trimming the array and return an error if not the case (i.e. the array is empty)
	if upperIndex-lowerIndex < 1 {
		return 0, fmt.Errorf("trimming results in fewer than one element")
	}

	// Slice the array to remove the lower and upper extremes of the sorted numbers array and calculate the mean
	trimmed := numbers[lowerIndex:upperIndex]

	// Calculate the mean of the trimmed slice of numbers and return it
	sum := 0.0
	for _, num := range trimmed {
		sum += num
	}
	mean := sum / float64(len(trimmed))

	return mean, nil
}

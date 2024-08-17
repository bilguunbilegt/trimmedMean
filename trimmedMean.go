package trimmedMean

import (
	"fmt"
	"sort"
)

// TrimmedMean calculates the mean after trimming a percentage of elements from both ends.
func TrimmedMean(numbers []float64, trim ...int) (float64, error) {
	if len(trim) == 0 || len(trim) > 2 {
		return 0, fmt.Errorf("invalid number of trimming arguments")
	}

	// Sort the numbers in ascending order
	sort.Float64s(numbers)

	// Determine the percentage of elements to trim from the lower and upper extremes of the sorted numbers array
	lowerTrim := float64(trim[0])
	upperTrim := lowerTrim
	if len(trim) == 2 {
		upperTrim = float64(trim[1])
	}

	// Validate trimming percentages
	if lowerTrim < 0 || lowerTrim >= 50 || upperTrim < 0 || upperTrim >= 50 {
		return 0, fmt.Errorf("trimming percentages must be between 0 and less than 50")
	}

	// Calculate number of elements to trim
	n := len(numbers)
	lowerCount := int(float64(n) * (lowerTrim / 100))
	upperCount := int(float64(n) * (upperTrim / 100))

	// Calculate indices for slicing the array to remove the lower and upper extremes
	lowerIndex := lowerCount
	upperIndex := n - upperCount

	if lowerIndex >= upperIndex || lowerIndex < 0 || upperIndex > n {
		return 0, fmt.Errorf("invalid trimming values, resulting in no elements or out-of-bounds indices")
	}

	// Check if at least one element remains after trimming the array
	if upperIndex-lowerIndex < 1 {
		return 0, fmt.Errorf("trimming results in fewer than one element")
	}

	// Slice the array to remove the lower and upper extremes and calculate the mean
	trimmed := numbers[lowerIndex:upperIndex]

	// Calculate the mean of the trimmed slice of numbers
	sum := 0.0
	for _, num := range trimmed {
		sum += num
	}
	mean := sum / float64(len(trimmed))

	return mean, nil
}

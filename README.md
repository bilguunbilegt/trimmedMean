# trimmedMean

The `trimmedMean` package provides a function to calculate the trimmed mean of a slice of floating-point numbers. The trimmed mean is a measure of central tendency that excludes a specified number of elements from both the lower and upper ends of the data set before computing the mean.

## Installation

To use this package in your Go project, you can simply import it:

```go
import "github.com/bilguunbilegt/trimmedMean"
```

## Usage

The `trimmedMean` function calculates the trimmed mean of a given slice of `float64` numbers. You can specify the number of elements to trim from the lower and upper ends of the sorted list.

### Function Signature

```go
func trimmedMean(numbers []float64, trim ...int) (float64, error)
```

### Parameters

- `numbers`: A slice of `float64` representing the data set.
- `trim`: A variable number of integer arguments. 
  - The first argument specifies the number of elements to trim from the lower end.
  - The second argument (optional) specifies the number of elements to trim from the upper end. If not provided, the function will trim the same number of elements from the upper end as from the lower end.

### Return Values

- `float64`: The calculated trimmed mean of the remaining elements.
- `error`: An error is returned if the trimming values result in an invalid state (e.g., too few elements remaining).

### Example

```go
package main

import (
    "fmt"
    "github.com/bilguunbilegt/trimmedMean"
)

func main() {
    numbers := []float64{1, 2, 3, 4, 5, 6, 7}
    mean, err := trimmedMean.TrimmedMean(numbers, 10, 10)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Trimmed Mean:", mean)
    }
}
```

This example trims 1 element from both the lower and upper ends of the sorted list, resulting in the trimmed mean of `4.0`.

## Error Handling

The `trimmedMean` function returns an error in the following cases:

- If no trimming arguments are provided.
- If more than two trimming arguments are provided.
- If trimming results in fewer than one element remaining.

### Example of Error Handling

```go
numbers := []float64{1, 2, 3, 4, 5}
mean, err := trimmedMean.TrimmedMean(numbers, 3, 3)
if err != nil {
    fmt.Println("Error:", err)
} else {
    fmt.Println("Trimmed Mean:", mean)
}
```

In this example, trimming 3 elements from both ends would result in fewer than one element remaining, so an error is returned.

## Testing

Unit tests are included to ensure the correctness of the `trimmedMean` function. The tests cover a variety of scenarios, including normal cases and edge cases.

### Running Tests

To run the tests, use the following command:

```bash
go test
```

This will execute the tests defined in `trimmedMean_test.go` and report any failures.

## GitHub Repository

You can find the source code on GitHub: [https://github.com/bilguunbilegt/trimmedMean](https://github.com/bilguunbilegt/trimmedMean)

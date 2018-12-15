//Package hamming finds the count of different characters in 2 strings
package hamming

import (
	"errors"
)

//Distance compares 2 string inputs in equal length to find number of difference characters between them.
func Distance(a, b string) (int, error) {

	if len(a) != len(b) {
		return 0, errors.New("uneven inputs")
	}

	count := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			count++
		}
	}

	return count, nil
}

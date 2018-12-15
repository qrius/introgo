//Package hamming finds the count of different characters in 2 strings
package hamming

import (
	"errors"
	"strings"
)

//Distance compares 2 string inputs in equal length to find number of difference characters between them.
func Distance(a, b string) (int, error) {

	count := 0
	sa := strings.Split(a, "")
	sb := strings.Split(b, "")

	if len(a) != len(b) {
		return -1, errors.New("uneven inputs")
	}

	for i := range sa {
		if sa[i] != sb[i] {
			count++
		}
	}

	return count, nil
}

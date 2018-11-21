//Package twofer is a simple exercise in switch/case statement and string building
package twofer

import (
	"fmt"
)

// ShareWith function returns a simple count for sharing scheme
func ShareWith(name string) string {

	if len(name) == 0 {
		name = "you"
	}

	return fmt.Sprintf("%s%s%s", "One for ", name, ", one for me.")
}

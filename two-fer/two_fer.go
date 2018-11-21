//Package twofer is a simple exercise in switch/case statement and string building
package twofer

import (
	"fmt"
)

// ShareWith function returns a simple count for sharing scheme
func ShareWith(name string) string {
	var v string
	// var b bytes.Buffer

	l := len(name)

	switch {
	case l == 0:
		v = "you"
	case l > 0:
		v = name
	default:
		v = "string not recognizable"
	}

	// b.WriteString("One for ")
	// b.WriteString(v)
	// b.WriteString(", one for me.")
	// return (b.String())
	return (fmt.Sprintf("%s%s%s", "One for ", v, ", one for me."))
}

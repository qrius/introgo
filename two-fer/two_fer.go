//Package twofer is a simple exercise in switch/case statement and string building
package twofer

// ShareWith should have a comment documenting it.
func ShareWith(name string) string {
	l := len(name)

	switch {
	case l == 0:
		return ("One for you, one for me.")
	case l > 0:
		return ("One for " + name + ", one for me.")
	default:
		return ("string not recognible")
	}
}

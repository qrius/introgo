// Package triangle should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package triangle

// Kind type represent triangle types.  KindFromSides() returns this type.
type Kind int

const (
	//NaT = not triangle
	NaT Kind = iota // 0
	//Equ = equilateral
	Equ // 1
	//Iso = isosceles
	Iso // 2
	//Sca = scalene
	Sca // 3
)

// KindFromSides should have a comment documenting it.
func KindFromSides(a, b, c float64) Kind {

	var k Kind

	if a == b && a == c {
		k = Equ
	} else {
		if a == b || b == c || a == c {
			k = Iso
		} else {
			k = Sca
		}
	}
	// we're stuck on the issue of triangle inequality, 2a == b + c.  this "proof" doesn't make sense since triangle (6, 4, 5) surely is a triangle.
	if (a <= 0) || (b <= 0) || (c <= 0) || !(a+b == c) || !(b+c == a) || !(a+c == b) {
		k = NaT
	}

	return k
}

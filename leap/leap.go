// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package leap calculates whether a given year is a leap year
package leap

// IsLeapYear function calculate whether a given year is a leap year, depends on the following logic:
//on every year that is evenly divisible by 4
// except every year that is evenly divisible by 100
//    unless the year is also evenly divisible by 400
func IsLeapYear(year int) bool {

	if year%4 == 0 {
		if year%100 == 0 {
			if year%400 == 0 {
				return true
			}
			return false
		}
		return true
	}
	return false
}

package mytypes

import "strings"

type MyInt int
type MyString string
type MyBuilder strings.Builder

// Twice multiplies its receiver by 2 and returns
// the result.
func (i MyInt) Twice() MyInt {
	return i * 2
}

// returns the number of characters in a string
func (s MyString) Len() int {
	return len(s)
}

func (mb MyBuilder) Hello() string {
	return "Hello, Gophers!"
}

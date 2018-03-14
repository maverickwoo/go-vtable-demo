/*

Package square is an implementation of shape.

*/
package square

import (
	"fmt"

	"github.com/maverickwoo/go-vtable-demo/shape"
)

// T is the exported type of square.
type T struct {
	shape.T     // Embed parent struct.
	side    int // Example private field.
}

// New returns a new named square.T with the specified side length.
func New(name string, side int) *T {
	t := T{
		T: shape.T{
			Name: name,
		},
		side: side,
	}
	t.VTable = &t // Important for VTable calls.
	return &t
}

// String makes T a fmt.Stringer. This method is static.
func (t *T) String() string {
	return fmt.Sprintf("%q is a square of side length %d.",
		t.Name, t.side)
}

// Dynamic returns t as the interface value of type VTable for parent methods.
func (t *T) Dynamic() shape.VTable {
	return t
}

//go:noinline
// Area returns the area of square t. This method is dynamic and NOT inlined.
func (t *T) Area() int {
	return t.side * t.side
}

// Bug is a method designed to show a runtime bug. This method is dynamic.
func (t *T) Buuuuug() { // In square, we misspell this method.
	return
}

// Default1 is a method to provide a default implementation for this subtree.
func (t *T) Default1() string {
	return "square.Default1"
}

// Useless assertion: this will compile and will not catch the typo above.
var _ shape.VTable = new(T)

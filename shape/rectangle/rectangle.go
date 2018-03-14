/*

Package rectangle is an implementation of shape.

*/
package rectangle

import (
	"fmt"

	"github.com/maverickwoo/go-vtable-demo/shape"
)

// T is the exported type of rectangle.
type T struct {
	shape.T           // Embed parent struct.
	Length, Width int // Example fields (exported for wide).
}

// New returns a new named rectangle.T of dimensions length and width.
func New(name string, length, width int) *T {
	t := T{
		T: shape.T{
			Name: name,
		},
		Length: length,
		Width:  width,
	}
	t.VTable = &t // Important for VTable calls.
	return &t
}

// String makes T a fmt.Stringer. This method is static.
func (t *T) String() string {
	return fmt.Sprintf("%q is a rectangle of dimensions %d x %d.",
		t.Name, t.Length, t.Width)
}

// Dynamic returns t as the interface value of type VTable for ancestor methods.
func (t *T) Dynamic() shape.VTable {
	return t
}

// Area returns the area of rectangle t. This method is dynamic.
func (t *T) Area() int {
	return t.Length * t.Width
}

// Bug is a method designed to show a runtime bug. This method is dynamic.
func (t *T) Bug() { // In rectangle, we spell this method correctly.
	return
}

// Default0 is a method to override the default implementation in shape.
func (t *T) Default0() string {
	return "rectangle.Default0->" + t.T.Default0()
}

// Default1 is a method to provide a default implementation for this subtree.
func (t *T) Default1() string {
	return "rectangle.Default1"
}

// Useless assertion: see square.
var _ shape.VTable = new(T)

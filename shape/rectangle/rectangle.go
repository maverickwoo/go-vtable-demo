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
	length, width int // Example private fields.
}

// New returns a new named rectangle.T of dimensions length and width.
func New(name string, length, width int) *T {
	t := T{
		T: shape.T{
			Name: name,
		},
		length: length,
		width:  width,
	}
	t.VTable = &t // Important for VTable calls.
	return &t
}

// String makes T a fmt.Stringer. This method is static.
func (t *T) String() string {
	return fmt.Sprintf("%q is a rectangle of dimensions %d x %d.",
		t.Name, t.length, t.width)
}

// Dynamic returns t as the interface value of type VTable for parent methods.
func (t *T) Dynamic() shape.VTable {
	return t
}

// Area returns the area of rectangle t. This method is dynamic.
func (t *T) Area() int {
	return t.length * t.width
}

// Bug is a method designed to show a runtime bug. This method is dynamic.
func (t *T) Bug() { // In rectangle, we spell this method correctly.
	return
}

// Useless assertion: see square.
var _ shape.VTable = new(T)

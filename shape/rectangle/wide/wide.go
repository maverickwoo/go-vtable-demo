/*

Package wide is a rectangle where the width is at least the length.

*/
package wide

import (
	"unsafe"

	"github.com/maverickwoo/go-vtable-demo/shape"
	"github.com/maverickwoo/go-vtable-demo/shape/rectangle"
)

// T is the exported type of a wide rectangle.
type T struct {
	rectangle.T
}

// New returns a new named wide.T of dimensions length <= width.
func New(name string, length, width int) *T {
	if !(length <= width) {
		width, length = length, width
	}
	// Same shape; different types. Just a toy example to demo invariant.
	t := (*T)(unsafe.Pointer(rectangle.New(name, length, width)))
	t.VTable = t // Important for VTable calls.
	return t
}

// Dynamic returns an interface value pointing to the ancestor, which has
// intentionally the same shape as wide.T (but not the same type). However,
// given our present hierarchy, we may as well just let rectangle.Dynamic be
// promoted here. So we have deliberately misspelled this method to imitate its
// absence.
func (t *T) Dynamic_not_present() shape.VTable {
	return &t.T
}

// Default0 is a method to override the default implementation in rectangle.
// Imagine this is a more efficient implementation given the invariant.
func (t *T) Default0() string {
	return "wide.Default0->" + t.T.Default0()
}

// Default1 is a method to override the default implementation in rectangle.
// Imagine this is a more efficient implementation given the invariant.
func (t *T) Default1() string {
	return "wide.Default1->" + t.T.Default1()
}

// String makes T a fmt.Stringer. This method is static.
func (t *T) String() string {
	return t.T.String() + " (With width invariant.)"
}

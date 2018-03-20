/*

Package shape is a virtual package designed to work with concrete
implementations.

*/
package shape

import "fmt"

// T is the exported type of shape.
type T struct {
	VTable        // Embedded interface value to implement dynamic dispatch.
	Name   string // Example exported field shared by all implementations.
	border string // Example unexported field shared by all implementations.
	fill   string // Ditto.
}

// VTable is the exported interface type for implementing dynamic dispatch.
type VTable interface {
	Dynamic() VTable         // Controlled approach to get t's dynamic type.
	Area() int               // Example virtual method that just works.
	Bug()                    // Example virtual method to demo runtime bug.
	Default0() string        // Example virtual method to demo override.
	Default1() string        // Example virtual method to demo override.
	SetBorder(string) VTable // Example virtual method to demo chaining.
	SetFill(string) VTable   // Example virtual method to demo chaining.
}

var Bug = false

// Print is a method of shape to demo dynamic dispatch.
func (t *T) Print() {
	fmt.Printf("----------------------------\n")

	// Static type is *shape.T.
	fmt.Printf("t in shape:\t\t%#v\n", t)

	// Dynamic type is great for reflection, e.g., template.Execute().
	fmt.Printf("t.Dynamic() in shape:\t%#v\n", t.Dynamic())

	// In usual cases one may just use t.VTable for the same effect.
	fmt.Printf("t.VTable in shape:\t%#v\n", t.VTable)

	// However, a more nuanced override can be achieved.
	fmt.Printf("----------------------------\n")
	fmt.Printf("t.Default0() in shape:\t%s\n", t.Default0())
	fmt.Printf("t.Dynamic().Default0():\t%s\n", t.Dynamic().Default0())
	fmt.Printf("t.VTable.Default0():\t%s\n", t.VTable.Default0())
	fmt.Printf("----------------------------\n")
	fmt.Printf("t.Default1() in shape:\t%s\n", t.Default1())
	fmt.Printf("t.Dynamic().Default1():\t%s\n", t.Dynamic().Default1())
	fmt.Printf("t.VTable.Default1():\t%s\n", t.VTable.Default1())
	fmt.Printf("----------------------------\n")

	// Other virtual methods in VTable can be used in a shape method too.
	fmt.Printf("t.Area() in shape:\t%d\n", t.Area())

	// Demonstrate a misspelled method name (and thus a missing method) can
	// lead to a runtime crash. See square.
	if Bug {
		fmt.Print("Calling t.Bug():\t")
		t.Bug()
		fmt.Print("Success!\n")
	}
}

// Default0 is a method of shape to let us demo overriding a default
// implementation. Note that t.Default0() in shape is a static call to this
// method and so one should be very careful when providing a default
// implementation at the same depth as its usage since it can give rise to
// surprises. OTOH, we specifically do not define Default1 at this depth and so
// t.Default1() in shape is a dynamic call without surprise.
func (t *T) Default0() string {
	return "shape.Default0"
}

// SetBorder is a method demonstrating how to set an unexported field in a
// chainable method.
func (t *T) SetBorder(s string) VTable {
	t.border = s
	return t
}

// SetFill is a method demonstrating how to set an unexported field in a
// chainable method.
func (t *T) SetFill(s string) VTable {
	t.fill = s
	return t
}

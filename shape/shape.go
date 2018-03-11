/*

Package shape is a virtual package designed to work with concrete
implementations.

*/
package shape

import "fmt"

// T is the exported type of shape.
type T struct {
	VTable        // Embedded interface value to implement dynamic dispatch.
	Name   string // Example struct field shared by all implementations.
}

// VTable is the exported interface type for implementing dynamic dispatch.
type VTable interface {
	Dynamic() VTable // Let methods in shape get t's dynamic type.
	Area() int       // Example virtual method that works.
	Bug()            // Example virtual method to demo runtime bug.
}

var Bug = false

// Print is a method of shape for demonstrating dynamic dispatch.
func (t *T) Print() {
	// Static type is *shape.T.
	fmt.Printf("t in shape:\t\t%#v\n", t)

	// Dynamic type is great for reflection, e.g., template.Execute().
	fmt.Printf("t.Dynamic() in shape:\t%#v\n", t.Dynamic())

	// Other virtual methods in VTable can be used in a shape method too.
	fmt.Printf("t.Area() in shape:\t%d\n", t.Area())

	// Demonstrate a misspelled method name can lead to a runtime crash.
	// See square.
	if Bug {
		fmt.Print("Calling t.Bug():\t")
		t.Bug()
		fmt.Print("Success!\n")
	}
}

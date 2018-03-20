package main

import (
	"flag"
	"fmt"

	"github.com/maverickwoo/go-vtable-demo/shape"
	"github.com/maverickwoo/go-vtable-demo/shape/rectangle"
	"github.com/maverickwoo/go-vtable-demo/shape/rectangle/wide"
	"github.com/maverickwoo/go-vtable-demo/shape/square"
)

func main() {
	flag.BoolVar(&shape.Bug, "b", false, "Set to show runtime bug.")
	flag.Parse()

	// Manufacture values; we spell out their types for clarity.
	var s1 *rectangle.T = rectangle.New("s1", 6, 7)
	var s2 *square.T = square.New("s2", 8)
	var s3 *wide.T = wide.New("s3", 10, 2) // Force invariant swap.

	// Print (static) and Area (dynamic) demos.
	fmt.Printf("VTable demo:\t\t%s\n", s1) // Static rectangle.String().
	s1.Print()                             // Static rectangle.Print(), promoted.
	_ = s1.Area()                          // Static rectangle.Area().
	fmt.Println()
	//
	fmt.Printf("VTable demo:\t\t%s\n", s2) // Static square.String().
	s2.Print()                             // Static square.Print(), promoted.
	_ = s2.Area()                          // Static square.Area().
	fmt.Println()
	//
	fmt.Printf("VTable demo:\t\t%s\n", s3) // Static wide.String().
	s3.Print()                             // Static wide.Print(), promoted.
	_ = s3.Area()                          // Static wide.Area().
	fmt.Println()

	// Chaining demo; note the the type changes after first chained call.
	var s4 shape.VTable = rectangle.
		New("s4", 8, 9).   // Returns rectangle.T.
		SetBorder("blue"). // Returns shape.VTable.
		SetFill("fuchsia") // Returns shape.VTable.
	// s4.Print() and s4.String() are illegal since they are not in VTable.
	fmt.Printf("Chaining demo:\n")
	fmt.Printf("----------------------------\n")
	fmt.Printf("t in main:\t\t%#v\n", s4)            // Reflection yields shape.T.
	fmt.Printf("t.Area() in main:\t%d\n", s4.Area()) // VTable works.
	fmt.Println()

	fmt.Println("Try running with -b for a crash.")
}

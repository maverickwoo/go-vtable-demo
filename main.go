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

	// Manufacture two values; we spell out their types for clarity.
	var s1 *rectangle.T = rectangle.New("s1", 6, 7)
	var s2 *square.T = square.New("s2", 8)
	var s3 *wide.T = wide.New("s3", 10, 2) // Force invariant swap.

	fmt.Printf("Demo:\t\t\t%s\n", s1) // Static rectangle.String().
	s1.Print()                        // Static rectangle.Print(), promoted.
	_ = s1.Area()                     // Static rectangle.Area().
	fmt.Println()

	fmt.Printf("Demo:\t\t\t%s\n", s2) // Static square.String().
	s2.Print()                        // Static square.Print(), promoted.
	_ = s2.Area()                     // Static square.Area().
	fmt.Println()

	fmt.Printf("Demo:\t\t\t%s\n", s3) // Static wide.String().
	s3.Print()                        // Static wide.Print(), promoted.
	_ = s3.Area()                     // Static wide.Area().
	fmt.Println()

	fmt.Println("Try running with -b for a crash.")
}

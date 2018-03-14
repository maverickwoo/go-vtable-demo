# Purpose

This repository is a demonstration of one approach to implement dynamic dispatch
in Go 1 using an explicit `VTable` interface value embedded in structs. I have
extracted this pattern from several of my own golang projects that implemented
dynamic dispatch in less explicit but similar ways. I believe spelling out the
pattern and naming the `VTable` interface value explicitly help me clarify my
thoughts and also make the pattern easier to execute in the future.

Please see the comment in the source code and
[![GoDoc](https://godoc.org/github.com/maverickwoo/go-vtable-demo?status.svg)](https://godoc.org/github.com/maverickwoo/go-vtable-demo)
for more information.

Also, please see [this Reddit
post](https://www.reddit.com/r/golang/comments/83qboe/the_vtable_pattern_in_go/)
and [this golang-nuts
thread](https://groups.google.com/forum/#!topic/golang-nuts/5FFiKxzDCDc) for
more discussion.

# Design

Suppose we want to implement a base `shape` package that

1. implements some base methods within the `shape` package, e.g., `Print()`, but
2. leaves some other methods for a "descendant package" such as `rectangle` to
provide, e.g., `Area()`, while
3. wanting to allow the base methods in `shape` to call descendant-provided
methods.

How would we do this? Go 1 does not have inheritance but it does have automatic
method promotion for embedded fields. So if only conditions 1 and 2 are needed,
this can be solved by struct embedding in the straightforward manner. However,
this approach does **not** achieve condition 3.

One way to achieve condition 3 is to make `shape.Print` takes an interface value
of type say `Areaer` (for the lack of a better name). For example:

```go
package shape

// [...]

type Areaer interface {
	Area() int
}

func (t *T) Print(a Areaer) {
	fmt.Printf("%s has area %d.", t.Name, a.Area())
}

// [...]
```

But this is clunky to use since, assuming `s` is the descendant under
consideration, we need to call `s.Print(s)` at a call site and we need to make a
distinction between `t` and `a` inside `Print`.

## `VTable` implementation of condition 3

The design in this demo project solves this problem by embedding an interface
value spelled `VTable` in `shape.T`, which is in turn embedded by value in
`rectangle.T`. By putting `Area()` in `shape.VTable` and making sure that
`rectangle.T.VTable` has the dynamic type `rectangle.T` (see `rectangle.New`), a
call to `t.Area()` in `shape.Print` would be dispatched to `rectangle.Area()`,
thus achieving conditions 1 through 3 in a seamless manner.

However, note that the embedded interface value adds two words. For small
structs, say the nodes in an AST, this amount of space overhead may be
prohibitively expensive.

## `Dynamic()` wrapper vs. `VTable`

A further enhancement is to introduce the `Dynamic()` method in `VTable`. This
simple wrapper wraps a descendant value in an interface value of type `VTable`
in `shape`. This enables any reflection facility to retrieve the actual type. In
this demo project, this is demonstrated by the use of the `%#v` verb. But one
may also consider Go's own `template.Execute()` when implementing an MVC where
`shape` stands in for an abstract model and `rectangle` and `square` stand in
for concrete models.

In the common cases where each descendant has its `Dynamic` method returns the
receiver, such as in `rectangle.Dynamic`, the `Dynamic()` declaration in the
interface is redundant: in `shape`, we can just replace `t.Dynamic()` with
`t.VTable`.

However, there may also be niche cases where we may prefer the `Dynamic()`
method of a descendant to return a `VTable` pointing to an ancestor of it. One
plausible example is demonstrated in `rectangle/wide`, where the `wide`
descendant wants to always expose its identically-shaped data in the type of
`rectangle.T`. This may be desirable to clients that performs serialization.

(As for the structural reason of why `wide` exists in its location in the
hierarchy, notice that the factory method of `wide` ensures an invariant, and we
may imagine that some methods of `rectangle` may have a more efficient
implementation given the invariant and so we may want to override them in
`wide`.)

## Typos can lead to a runtime crash

This demo project also demonstrates what happens if one of the descendants do
not provide a method in `VTable`. In our example, the `square.Bug()` method is
missing due to a "typo": `square.Buuuuug()`. The result is a runtime crash when
`shape.Print` calls `VTable.Bug()`. Note that the "assertion" at the end of the
`square` package is statically true because `square.T` embeds `shape.VTable` and
is thus unable to catch the "typo".

# Performance

For the trivial function `Area`, there is a ~8 times slowdown on my laptop when
there is inlining (as in rectangle). But when there is no inlining (as in
square), the overhead is very modest.

```
$ go version
go version go1.10 darwin/amd64

$ go test -bench=. -benchmem -cpu 1
goos: darwin
goarch: amd64
pkg: github.com/maverickwoo/go-vtable-demo
BenchmarkRectangleAreaStatic  	2000000000	         0.33 ns/op	       0 B/op	       0 allocs/op
BenchmarkRectangleAreaDynamic 	1000000000	         2.59 ns/op	       0 B/op	       0 allocs/op
BenchmarkSquareAreaStatic     	1000000000	         2.07 ns/op	       0 B/op	       0 allocs/op
BenchmarkSquareAreaDynamic    	1000000000	         2.68 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/maverickwoo/go-vtable-demo	8.782s
```

# Notes

* Please remember to run this with `-b` to see the runtime crash.

* This `VTable` pattern is intended for a very specific dynamic-dispatch
  scenario as explained above. Briefly, it is when a base package has a need to
  call descendant-provided methods and the application does not mind the space
  overhead. For all other cases, exposing functions that accept an interface
  argument should be preferred.

* I have not studied the performance implications of this pattern at the
  instruction level yet, but the overhead in benchmark seems reasonable. For
  simple methods, the overhead is big and thus harder to justify; but for more
  sophisticated methods, this may be acceptable.

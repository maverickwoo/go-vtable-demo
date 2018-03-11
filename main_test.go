package main

import (
	"testing"

	"github.com/maverickwoo/go-vtable-demo/shape"
	"github.com/maverickwoo/go-vtable-demo/shape/rectangle"
	"github.com/maverickwoo/go-vtable-demo/shape/square"
)

func BenchmarkRectangleAreaStatic(b *testing.B) {
	b.StopTimer()
	var concrete *rectangle.T = rectangle.New("sq", 6, 7)
	b.StartTimer()
	for n := 0; n < b.N; n++ {
		_ = concrete.Area()
	}
}

func BenchmarkRectangleAreaDynamic(b *testing.B) {
	b.StopTimer()
	var concrete *rectangle.T = rectangle.New("sq", 6, 7)
	var virtual *shape.T = &concrete.T
	b.StartTimer()
	for n := 0; n < b.N; n++ {
		_ = virtual.Area()
	}
}

func BenchmarkSquareAreaStatic(b *testing.B) {
	b.StopTimer()
	var concrete *square.T = square.New("sq", 8)
	b.StartTimer()
	for n := 0; n < b.N; n++ {
		_ = concrete.Area()
	}
}

func BenchmarkSquareAreaDynamic(b *testing.B) {
	b.StopTimer()
	var concrete *square.T = square.New("sq", 8)
	var virtual *shape.T = &concrete.T
	b.StartTimer()
	for n := 0; n < b.N; n++ {
		_ = virtual.Area()
	}
}

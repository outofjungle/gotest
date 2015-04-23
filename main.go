package main

import (
	"fmt"

	"github.com/outofjungle/gotest/vector"
)

func main() {
	p1 := vector.Point{
		X: -1,
		Y: 2,
	}

	p2 := vector.Point{
		X: -3,
		Y: 6,
	}

	scalar := p1.Dot(p2)
	fmt.Println(scalar)
}

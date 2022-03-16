package geom_test

import (
	"fmt"

	"github.com/eihigh/geom"
)

func ExampleRect() {
	r1 := geom.Recti{10, 10, 20, 20}
	r2 := geom.XYWHToRect(10, 10, 10, 10)
	r3 := geom.PosSizeToRect(geom.Vec2i{10, 10}, geom.Vec2i{10, 10})
	fmt.Println(r1.Eq(r2))
	fmt.Println(r1.Eq(r3))

	// Output:
	// true
	// true
}

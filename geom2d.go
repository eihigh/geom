package geom

import (
	"image"
	"math"
)

type Vec2[T intfloat] struct {
	X, Y T
}

type Vec2f = Vec2[float64]

type Vec2i = Vec2[int]

func PolarToVec2(r, theta float64) Vec2f {
	s, c := math.Sincos(theta)
	return Vec2f{c * r, s * r}
}

func (v Vec2[T]) Add(other Vec2[T]) Vec2[T] {
	return Vec2[T]{v.X + other.X, v.Y + other.Y}
}

func (v Vec2[T]) Eq(other Vec2[T]) bool {
	return v == other
}

func (v Vec2[T]) LengthSq() T {
	return v.X*v.X + v.Y*v.Y
}

func (v Vec2[T]) Length() float64 {
	return math.Sqrt(float64(v.LengthSq()))
}

func (v Vec2[T]) Angle() float64 {
	return math.Atan2(float64(v.Y), float64(v.X))
}

func (v Vec2[T]) Polar() (r, theta float64) {
	return v.Length(), v.Angle()
}

func (v Vec2[T]) XY() (x, y T) {
	return v.X, v.Y
}

func (v Vec2[T]) Complex64() complex64 {
	return complex(float32(v.X), float32(v.Y))
}

func (v Vec2[T]) Complex128() complex128 {
	return complex(float64(v.X), float64(v.Y))
}

type Rect[T intfloat] struct {
	X0, Y0, X1, Y1 T
}

type Recti = Rect[int]

type Rectf = Rect[float64]

func XYWHToRect[T intfloat](x, y, w, h T) Rect[T] {
	return Rect[T]{x, y, x + w, y + h}
}

func MinMaxToRect[T intfloat](min, max Vec2[T]) Rect[T] {
	return Rect[T]{min.X, min.Y, max.X, max.Y}
}

func PosSizeToRect[T intfloat](pos, size Vec2[T]) Rect[T] {
	return Rect[T]{pos.X, pos.Y, pos.X + size.X, pos.Y + size.Y}
}

func (r Rect[T]) Min() Vec2[T] {
	return Vec2[T]{r.X0, r.Y0}
}

func (r Rect[T]) Max() Vec2[T] {
	return Vec2[T]{r.X1, r.Y1}
}

func (r Rect[T]) Dx() T {
	return r.X1 - r.X0
}

func (r Rect[T]) Dy() T {
	return r.Y1 - r.Y0
}

func (r Rect[T]) Size() Vec2[T] {
	return Vec2[T]{r.X1 - r.X0, r.Y1 - r.Y0}
}

func (r Rect[T]) Empty() bool {
	return r.X0 >= r.X1 || r.Y0 >= r.Y1
}

func (r Rect[T]) Eq(other Rect[T]) bool {
	return r == other || r.Empty() && other.Empty()
}

func (r Rect[T]) Image() image.Rectangle {
	return image.Rect(
		int(math.Round(float64(r.X0))),
		int(math.Round(float64(r.Y0))),
		int(math.Round(float64(r.X1))),
		int(math.Round(float64(r.Y1))),
	)
}

type Line2[T intfloat] struct {
	X0, Y0, X1, Y1 T
}

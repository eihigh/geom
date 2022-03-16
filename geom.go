package geom

import "math"

const (
	Tau = math.Pi * 2
)

type intfloat interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~float32 | ~float64
}

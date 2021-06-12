package constraints

import (
	"image"

	"github.com/cybriq/giocore/f32"
)

// Constraints define the minimum and maximum dimensions of a layout area
type Constraints struct {
	Min, Max image.Point
}

// Exact returns the Constraints with the minimum and maximum size set to size.
func Exact(size image.Point) Constraints {
	return Constraints{
		Min: size, Max: size,
	}
}

// FPt converts an point to a f32.Point.
func FPt(p image.Point) f32.Point {
	return f32.Point{
		X: float32(p.X), Y: float32(p.Y),
	}
}

// FRect converts a rectangle to a f32.Rectangle.
func FRect(r image.Rectangle) f32.Rectangle {
	return f32.Rectangle{
		Min: FPt(r.Min), Max: FPt(r.Max),
	}
}

// Constrain a size so each dimension is in the range [min;max].
func (c Constraints) Constrain(size image.Point) image.Point {
	if min := c.Min.X; size.X < min {
		size.X = min
	}
	if min := c.Min.Y; size.Y < min {
		size.Y = min
	}
	if max := c.Max.X; size.X > max {
		size.X = max
	}
	if max := c.Max.Y; size.Y > max {
		size.Y = max
	}
	return size
}

package dim

import (
	"image"
)

// Constraints define the minimum and maximum dimensions of a layout area
type Constraints struct {
	Min, Max image.Point
}

// Exact returns the Constraints with the minimum and maximum size set to size
func Exact(size image.Point) Constraints {
	return Constraints{
		Min: size, Max: size,
	}
}

// Constrain a size so each dimension is in the range [min;max]
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

package canvas

import (
	"image"
)

// Lim defines the minimum and maximum dimensions of a layout area
type Lim struct {
	Min, Max image.Point
}

// Exact returns the Lim with the minimum and maximum size set to size
func Exact(size image.Point) Lim {
	return Lim{
		Min: size, Max: size,
	}
}

// Constrain a size so each dimension is in the range [min;max]
func (c Lim) Constrain(size image.Point) image.Point {
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

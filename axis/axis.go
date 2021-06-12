package axis

import (
	"image"

	"github.com/cybriq/pokaz/constraints"
)

// Axis defines the direction of a widget with multiple elements
type Axis uint8

const (
	Horizontal Axis = iota
	Vertical
)

// Convert a point in (x, y) coordinates to (main, cross) coordinates,
// or vice versa. Specifically, Convert((x, y)) returns (x, y) unchanged
// for the horizontal axis, or (y, x) for the vertical axis.
func (a Axis) Convert(pt image.Point) image.Point {
	if a == Horizontal {
		return pt
	}
	return image.Pt(pt.Y, pt.X)
}

// mainConstraint returns the min and max main constraints for axis a.
func (a Axis) mainConstraint(cs constraints.Constraints) (int, int) {
	if a == Horizontal {
		return cs.Min.X, cs.Max.X
	}
	return cs.Min.Y, cs.Max.Y
}

// crossConstraint returns the min and max cross constraints for axis a.
func (a Axis) crossConstraint(cs constraints.Constraints) (int, int) {
	if a == Horizontal {
		return cs.Min.Y, cs.Max.Y
	}
	return cs.Min.X, cs.Max.X
}

// constraints returns the constraints for axis a.
func (a Axis) constraints(mainMin, mainMax, crossMin, crossMax int) constraints.Constraints {
	if a == Horizontal {
		return constraints.Constraints{Min: image.Pt(mainMin, crossMin), Max: image.Pt(mainMax, crossMax)}
	}
	return constraints.Constraints{Min: image.Pt(crossMin, mainMin), Max: image.Pt(crossMax, mainMax)}
}

func (a Axis) String() string {
	switch a {
	case Horizontal:
		return "Horizontal"
	case Vertical:
		return "Vertical"
	default:
		panic("unreachable")
	}
}

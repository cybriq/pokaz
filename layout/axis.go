package layout

import (
	"image"
)

// Axis defines the dir of a widget with multiple elements
type Axis uint8

const (
	Horizontal Axis = iota
	Vertical
)

// Convert a point in (x, y) coordinates to (main, cross) coordinates, or vice
// versa. Specifically, Convert((x, y)) returns (x, y) unchanged for the
// horizontal axis, or (y, x) for the vertical axis.
func (a Axis) Convert(pt image.Point) image.Point {
	if a == Horizontal {
		return pt
	}
	return image.Pt(pt.Y, pt.X)
}

// MainConstraint returns the min and max main Lim for axis a
func (a Axis) MainConstraint(cs Lim) (int, int) {
	if a == Horizontal {
		return cs.Min.X, cs.Max.X
	}
	return cs.Min.Y, cs.Max.Y
}

// CrossConstraint returns the min and max cross Lim for axis a
func (a Axis) CrossConstraint(cs Lim) (int, int) {
	if a == Horizontal {
		return cs.Min.Y, cs.Max.Y
	}
	return cs.Min.X, cs.Max.X
}

// Lim return the Lim for axis a
func (a Axis) Constraints(mainMin, mainMax, crossMin, crossMax int) Lim {
	if a == Horizontal {
		return Lim{
			Min: image.Pt(mainMin, crossMin),
			Max: image.Pt(mainMax, crossMax),
		}
	}
	return Lim{
		Min: image.Pt(crossMin, mainMin),
		Max: image.Pt(crossMax, mainMax),
	}
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

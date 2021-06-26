package layout

import (
	"image"

	"github.com/cybriq/giocore/f32"
)

// Point converts an image.Point to a f32.Point
func Point(p image.Point) f32.Point {
	return f32.Point{
		X: float32(p.X), Y: float32(p.Y),
	}
}

// Rect converts an image.Rectangle to a f32.Rectangle
func Rect(r image.Rectangle) f32.Rectangle {
	return f32.Rectangle{
		Min: Point(r.Min), Max: Point(r.Max),
	}
}

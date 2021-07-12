package canvas

import (
	"image"

	"github.com/cybriq/giocore/f32"
)

// ToPoint converts an image.Point to a f32.Point
func ToPoint(p image.Point) f32.Point {
	return f32.Point{
		X: float32(p.X), Y: float32(p.Y),
	}
}

// ToRect converts an image.Rectangle to a f32.Rectangle
func ToRect(r image.Rectangle) f32.Rectangle {
	return f32.Rectangle{
		Min: ToPoint(r.Min), Max: ToPoint(r.Max),
	}
}

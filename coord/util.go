package coord

import (
	"image"

	"github.com/cybriq/giocore/f32"
)

// FPt converts an image.Point to a f32.Point
func FPt(p image.Point) f32.Point {
	return f32.Point{
		X: float32(p.X), Y: float32(p.Y),
	}
}

// FRect converts an image.Rectangle to a f32.Rectangle
func FRect(r image.Rectangle) f32.Rectangle {
	return f32.Rectangle{
		Min: FPt(r.Min), Max: FPt(r.Max),
	}
}


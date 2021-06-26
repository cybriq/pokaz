package layout

import (
	"image"
	
	"github.com/cybriq/giocore/unit"
)

// spacer adds space between widgets.
type spacer struct {
	Width, Height unit.Value
}

// Spacer makes a new spacer
func Spacer(width, height float32) spacer {
	return spacer{unit.Sp(width), unit.Sp(height)}
}

// Fn lays out a spacer
func (s spacer) Fn(gtx Ctx) Dims {
	return Dims{
		Size: image.Point{
			X: gtx.Px(s.Width),
			Y: gtx.Px(s.Height),
		},
	}
}

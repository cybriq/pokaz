package layout

import (
	"image"
	
	"github.com/cybriq/giocore/unit"
)

// Spacer adds space between widgets.
type Spacer struct {
	Width, Height unit.Value
}

// NewSpacer makes a new Spacer
func NewSpacer(width, height float32) Spacer {
	return Spacer{unit.Sp(width), unit.Sp(height)}
}

// Fn lays out a Spacer
func (s Spacer) Fn(gtx Ctx) Dims {
	return Dims{
		Size: image.Point{
			X: gtx.Px(s.Width),
			Y: gtx.Px(s.Height),
		},
	}
}

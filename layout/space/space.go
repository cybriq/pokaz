package space

import (
	"image"

	"github.com/cybriq/giocore/unit"
	"github.com/cybriq/pokaz/layout/ctx"
	"github.com/cybriq/pokaz/layout/dim"
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
func (s spacer) Fn(gtx ctx.Context) dim.Dimensions {
	return dim.Dimensions{
		Size: image.Point{
			X: gtx.Px(s.Width),
			Y: gtx.Px(s.Height),
		},
	}
}

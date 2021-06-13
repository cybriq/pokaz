package space

import (
	"image"

	"github.com/cybriq/giocore/unit"
	"github.com/cybriq/pokaz/ctx"
	"github.com/cybriq/pokaz/dims"
)

// spacer adds space between widgets.
type spacer struct {
	Width, Height unit.Value
}

func (s spacer) Fn(gtx ctx.Context) dims.Dimensions {
	return dims.Dimensions{
		Size: image.Point{
			X: gtx.Px(s.Width),
			Y: gtx.Px(s.Height),
		},
	}
}

func New(width, height float32) spacer {
	return spacer{unit.Sp(width), unit.Sp(height)}
}

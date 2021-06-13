package inset

import (
	"image"

	"github.com/cybriq/giocore/op"
	"github.com/cybriq/giocore/unit"
	"github.com/cybriq/pokaz/coord"
	"github.com/cybriq/pokaz/ctx"
	"github.com/cybriq/pokaz/dims"
	"github.com/cybriq/pokaz/wdg"
)

// inset adds space around a widget by decreasing its maximum constraints. The
// minimum constraints will be adjusted to ensure they do not exceed the
// maximum.
type inset struct {
	Top, Right, Bottom, Left unit.Value
}

// layout an inset.
func (in inset) layout(gtx ctx.Context, w wdg.Widget) dims.Dimensions {
	top := gtx.Px(in.Top)
	right := gtx.Px(in.Right)
	bottom := gtx.Px(in.Bottom)
	left := gtx.Px(in.Left)
	mcs := gtx.Constraints
	mcs.Max.X -= left + right
	if mcs.Max.X < 0 {
		left = 0
		right = 0
		mcs.Max.X = 0
	}
	if mcs.Min.X > mcs.Max.X {
		mcs.Min.X = mcs.Max.X
	}
	mcs.Max.Y -= top + bottom
	if mcs.Max.Y < 0 {
		bottom = 0
		top = 0
		mcs.Max.Y = 0
	}
	if mcs.Min.Y > mcs.Max.Y {
		mcs.Min.Y = mcs.Max.Y
	}
	stack := op.Save(gtx.Ops)
	op.Offset(coord.FPt(image.Point{X: left, Y: top})).Add(gtx.Ops)
	gtx.Constraints = mcs
	dm := w(gtx)
	stack.Load()
	return dims.Dimensions{
		Size:     dm.Size.Add(image.Point{X: right + left, Y: top + bottom}),
		Baseline: dm.Baseline + bottom,
	}
}

// uniform returns an inset with a single inset applied to all edges.
func uniform(v unit.Value) inset {
	return inset{Top: v, Right: v, Bottom: v, Left: v}
}

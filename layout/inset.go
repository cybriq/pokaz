package layout

import (
	"image"

	"github.com/cybriq/giocore/op"
	"github.com/cybriq/giocore/unit"
)

// NewInset creates a padded empty space around a widget
func NewInset(pad float32, embed Widget) (out *Inset) {
	out = &Inset{
		insetSpec: uniform(unit.Sp(pad)),
		w:         embed,
	}
	return
}

// Embed sets the widget that will be inside the Inset
func (in *Inset) Embed(w Widget) *Inset {
	in.w = w
	return in
}

type insetSpec struct {
	Top, Right, Bottom, Left unit.Value
}

// Inset adds space around a widget by decreasing its maximum constraints. The
// minimum constraints will be adjusted to ensure they do not exceed the
// maximum.
type Inset struct {
	insetSpec
	w Widget
}

// Fn lays out an Inset.
func (in Inset) Fn(gtx Ctx) Dims {
	top := gtx.Px(in.Top)
	right := gtx.Px(in.Right)
	bottom := gtx.Px(in.Bottom)
	left := gtx.Px(in.Left)
	mcs := gtx.Lim
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
	op.Offset(ToPoint(image.Point{X: left, Y: top})).Add(gtx.Ops)
	gtx.Lim = mcs
	dm := in.w(gtx)
	stack.Load()
	return Dims{
		Size:     dm.Size.Add(image.Point{X: right + left, Y: top + bottom}),
		Baseline: dm.Baseline + bottom,
	}
}

// uniform returns an Inset with a single Inset applied to all edges.
func uniform(v unit.Value) insetSpec {
	return insetSpec{Top: v, Right: v, Bottom: v, Left: v}
}

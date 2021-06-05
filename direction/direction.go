package direction

import (
	"github.com/l0k18/giocore/op"
	"github.com/l0k18/pokaz/constraints"
	"github.com/l0k18/pokaz/outline"
	"image"
)

type Direction uint8

const (
	NW Direction = iota
	N
	NE
	E
	SE
	S
	SW
	W
	Center
	endDirections
)

var Directions = []string{
	"NW", "N", "NE", "E", "SE", "S", "SW", "W", "Center",
}

// String returns the name of the direction in string form
func (d Direction) String() string {
	if d < 0 || d >= endDirections {
		panic("direction is out of bounds")
	}
	return Directions[d]
}

// Layout a widget according to the direction.
// The widget is called with the context constraints minimum cleared.
func (d Direction) Layout(
	gtx outline.Context, w outline.Widget,
) outline.Dimensions {
	macro := op.Record(gtx.Ops)
	cs := gtx.Constraints
	gtx.Constraints.Min = image.Point{}
	dims := w(gtx)
	call := macro.Stop()
	sz := dims.Size
	if sz.X < cs.Min.X {
		sz.X = cs.Min.X
	}
	if sz.Y < cs.Min.Y {
		sz.Y = cs.Min.Y
	}

	defer op.Save(gtx.Ops).Load()
	p := d.Position(dims.Size, sz)
	op.Offset(constraints.FPt(p)).Add(gtx.Ops)
	call.Add(gtx.Ops)

	return outline.Dimensions{
		Size:     sz,
		Baseline: dims.Baseline + sz.Y - dims.Size.Y - p.Y,
	}
}

// Position calculates widget position according to the direction.
func (d Direction) Position(widget, bounds image.Point) image.Point {
	var p image.Point

	switch d {
	case N, S, Center:
		p.X = (bounds.X - widget.X) / 2
	case NE, SE, E:
		p.X = bounds.X - widget.X
	}

	switch d {
	case W, Center, E:
		p.Y = (bounds.Y - widget.Y) / 2
	case SW, S, SE:
		p.Y = bounds.Y - widget.Y
	}

	return p
}

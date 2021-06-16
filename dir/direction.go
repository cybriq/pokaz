package dir

import (
	"image"

	"github.com/cybriq/giocore/op"
	"github.com/cybriq/pokaz/coord"
	"github.com/cybriq/pokaz/ctx"
	"github.com/cybriq/pokaz/dim"
	"github.com/cybriq/pokaz/wdg"
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
		panic("dir is out of bounds")
	}
	return Directions[d]
}

// Fn lays out a widget according to the dir. The widget is called with the
// context constraints minimum cleared.
func (d Direction) Fn(
	gtx ctx.Context, w wdg.Widget,
) dim.Dimensions {
	macro := op.Record(gtx.Ops)
	cs := gtx.Constraints
	gtx.Constraints.Min = image.Point{}
	dim := w(gtx)
	call := macro.Stop()
	sz := dim.Size
	if sz.X < cs.Min.X {
		sz.X = cs.Min.X
	}
	if sz.Y < cs.Min.Y {
		sz.Y = cs.Min.Y
	}

	defer op.Save(gtx.Ops).Load()
	p := d.Position(dim.Size, sz)
	op.Offset(coord.FPt(p)).Add(gtx.Ops)
	call.Add(gtx.Ops)

	return dim.Dimensions{
		Size:     sz,
		Baseline: dim.Baseline + sz.Y - dim.Size.Y - p.Y,
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

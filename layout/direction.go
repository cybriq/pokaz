package layout

import (
	"image"
	
	"github.com/cybriq/giocore/op"
)

type Dir uint8

const (
	NW Dir = iota
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

var directions = []string{
	"NW", "N", "NE", "E", "SE", "S", "SW", "W", "Center",
}

// String returns the name of the direction in string form
func (d Dir) String() string {
	if d < 0 || d >= endDirections {
		panic("dir is out of bounds")
	}
	return directions[d]
}

// Fn lays out a widget according to the dir. The widget is called with the
// context constraints minimum cleared.
func (d Dir) Fn(
	gtx Ctx, w Widget,
) Dims {
	macro := op.Record(gtx.Ops)
	cs := gtx.Lim
	gtx.Lim.Min = image.Point{}
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
	op.Offset(Point(p)).Add(gtx.Ops)
	call.Add(gtx.Ops)

	return Dims{
		Size:     sz,
		Baseline: dims.Baseline + sz.Y - dims.Size.Y - p.Y,
	}
}

// Position calculates widget position according to the direction.
func (d Dir) Position(widget, bounds image.Point) image.Point {
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

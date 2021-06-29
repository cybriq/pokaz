package layout

import (
	"image"

	"github.com/cybriq/giocore/op"
)

type Dir uint8

type Direction struct {
	Dir
	Widget
}

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

// String returns the name of the direction insetSpec string form
func (d Direction) String() string {
	if d.Dir < NW || d.Dir >= endDirections {
		panic("dir is out of bounds")
	}
	return directions[d.Dir]
}

// NewDirection creates a new Direction
func NewDirection(d Dir, w Widget) Direction {
	return Direction{
		Dir:    d,
		Widget: w,
	}
}

// Fn lays out a widget according to the dir. The widget is called with the
// context constraints minimum cleared.
func (d Direction) Fn(gtx Ctx) Dims {
	macro := op.Record(gtx.Ops)
	cs := gtx.Lim
	gtx.Lim.Min = image.Point{}
	dims := d.Widget(gtx)
	call := macro.Stop()
	sz := dims.Size
	if sz.X < cs.Min.X {
		sz.X = cs.Min.X
	}
	if sz.Y < cs.Min.Y {
		sz.Y = cs.Min.Y
	}

	defer op.Save(gtx.Ops).Load()
	p := d.position(dims.Size, sz)
	op.Offset(ToPoint(p)).Add(gtx.Ops)
	call.Add(gtx.Ops)

	return Dims{
		Size:     sz,
		Baseline: dims.Baseline + sz.Y - dims.Size.Y - p.Y,
	}
}

// position calculates widget position according to the direction.
func (d Direction) position(widget, bounds image.Point) image.Point {
	var p image.Point

	switch d.Dir {
	case N, S, Center:
		p.X = (bounds.X - widget.X) / 2
	case NE, SE, E:
		p.X = bounds.X - widget.X
	}

	switch d.Dir {
	case W, Center, E:
		p.Y = (bounds.Y - widget.Y) / 2
	case SW, S, SE:
		p.Y = bounds.Y - widget.Y
	}

	return p
}

// SPDX-License-Identifier: Unlicense OR MIT

package stak

import (
	"image"
	
	"github.com/cybriq/giocore/op"
	"github.com/cybriq/pokaz/layout/conv"
	"github.com/cybriq/pokaz/layout/ctx/proto"
	"github.com/cybriq/pokaz/layout/dim"
	"github.com/cybriq/pokaz/layout/dir"
	"github.com/cybriq/pokaz/layout/wdg"
)

// stack lays out child elements on top of each other, according to an alignment
// dir.
type stack struct {
	// alignment is the dir to align children smaller than the available space.
	alignment dir.Direction
}

// child represents a child for a stack layout.
type child struct {
	expanded bool
	widget   wdg.Widget

	// Scratch space.
	call op.CallOp
	dims dim.Dimensions
}

// stacked returns a stack child that is laid out with no minimum constraints and
// maximum constraints passed to stack.layout.
func stacked(w wdg.Widget) child {
	return child{
		widget: w,
	}
}

// expanded returns a stack child with the minimum constraints set to the largest
// stacked child. The maximum constraints are set to the same as passed to
// stack.layout.
func expanded(w wdg.Widget) child {
	return child{
		expanded: true,
		widget:   w,
	}
}

// layout a stack of children. The position of the children are determined by the
// specified order, but stacked children are laid out before expanded children.
func (s stack) layout(
	gtx proto.Context,
	children ...child,
) dim.Dimensions {
	var maxSZ image.Point
	// First lay out stacked children.
	ct := gtx
	cs := ct.Constraints()
	cs.Min = image.Point{}
	for i, w := range children {
		if w.expanded {
			continue
		}
		macro := op.Record(gtx.Ops())
		d := w.widget(ct)
		call := macro.Stop()
		if w := d.Size.X; w > maxSZ.X {
			maxSZ.X = w
		}
		if h := d.Size.Y; h > maxSZ.Y {
			maxSZ.Y = h
		}
		children[i].call = call
		children[i].dims = d
	}
	// Then lay out expanded children.
	for i, w := range children {
		if !w.expanded {
			continue
		}
		macro := op.Record(gtx.Ops())
		cs.Min = maxSZ
		d := w.widget(ct)
		call := macro.Stop()
		if w := d.Size.X; w > maxSZ.X {
			maxSZ.X = w
		}
		if h := d.Size.Y; h > maxSZ.Y {
			maxSZ.Y = h
		}
		children[i].call = call
		children[i].dims = d
	}

	maxSZ = cs.Constrain(maxSZ)
	var baseline int
	for _, ch := range children {
		sz := ch.dims.Size
		var p image.Point
		switch s.alignment {
		case dir.N, dir.S, dir.Center:
			p.X = (maxSZ.X - sz.X) / 2
		case dir.NE, dir.SE, dir.E:
			p.X = maxSZ.X - sz.X
		}
		switch s.alignment {
		case dir.W, dir.Center, dir.E:
			p.Y = (maxSZ.Y - sz.Y) / 2
		case dir.SW, dir.S, dir.SE:
			p.Y = maxSZ.Y - sz.Y
		}
		stack := op.Save(gtx.Ops())
		op.Offset(conv.Point(p)).Add(gtx.Ops())
		ch.call.Add(gtx.Ops())
		stack.Load()
		if baseline == 0 {
			if b := ch.dims.Baseline; b != 0 {
				baseline = b + maxSZ.Y - sz.Y - p.Y
			}
		}
	}
	return dim.Dimensions{
		Size:     maxSZ,
		Baseline: baseline,
	}
}

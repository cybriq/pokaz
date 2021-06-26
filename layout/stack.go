package layout

import (
	"image"
	
	"github.com/cybriq/giocore/op"
)

type _stack struct {
	*stack
	children []child
}

// Stack starts a chain of widgets to compose into a stack
func Stack() (out *_stack) {
	out = &_stack{stack: &stack{}}
	return
}

func (s *_stack) Alignment(alignment Direction) *_stack {
	s.alignment = alignment
	return s
}

// Stacked appends a widget to the stack, the stack's dimensions will be
// computed from the largest widget in the stack
func (s *_stack) Stacked(w Widget) (out *_stack) {
	s.children = append(s.children, stacked(w))
	return s
}

// Expanded lays out a widget with the same max constraints as the stack
func (s *_stack) Expanded(w Widget) (out *_stack) {
	s.children = append(s.children, expanded(w))
	return s
}

// Fn runs the ops queue configured in the stack
func (s *_stack) Fn(c Context) Dimensions {
	return s.stack.layout(c, s.children...)
}

// stack lays out child elements on top of each other, according to an alignment
// dir.
type stack struct {
	// alignment is the dir to align children smaller than the available space.
	alignment Direction
}

// child represents a child for a stack layout.
type child struct {
	expanded bool
	widget   Widget

	// Scratch space.
	call op.CallOp
	dims Dimensions
}

// stacked returns a stack child that is laid out with no minimum constraints and
// maximum constraints passed to stack.layout.
func stacked(w Widget) child {
	return child{
		widget: w,
	}
}

// expanded returns a stack child with the minimum constraints set to the largest
// stacked child. The maximum constraints are set to the same as passed to
// stack.layout.
func expanded(w Widget) child {
	return child{
		expanded: true,
		widget:   w,
	}
}

// layout a stack of children. The position of the children are determined by the
// specified order, but stacked children are laid out before expanded children.
func (s stack) layout(
	gtx Context,
	children ...child,
) Dimensions {
	var maxSZ image.Point
	// First lay out stacked children.
	ct := gtx
	ct.Constraints.Min = image.Point{}
	for i, w := range children {
		if w.expanded {
			continue
		}
		macro := op.Record(gtx.Ops)
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
		macro := op.Record(gtx.Ops)
		ct.Constraints.Min = maxSZ
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

	maxSZ = gtx.Constraints.Constrain(maxSZ)
	var baseline int
	for _, ch := range children {
		sz := ch.dims.Size
		var p image.Point
		switch s.alignment {
		case N, S, Center:
			p.X = (maxSZ.X - sz.X) / 2
		case NE, SE, E:
			p.X = maxSZ.X - sz.X
		}
		switch s.alignment {
		case W, Center, E:
			p.Y = (maxSZ.Y - sz.Y) / 2
		case SW, S, SE:
			p.Y = maxSZ.Y - sz.Y
		}
		stack := op.Save(gtx.Ops)
		op.Offset(Point(p)).Add(gtx.Ops)
		ch.call.Add(gtx.Ops)
		stack.Load()
		if baseline == 0 {
			if b := ch.dims.Baseline; b != 0 {
				baseline = b + maxSZ.Y - sz.Y - p.Y
			}
		}
	}
	return Dimensions{
		Size:     maxSZ,
		Baseline: baseline,
	}
}

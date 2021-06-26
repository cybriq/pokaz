package layout

import (
	"image"
	
	"github.com/cybriq/giocore/op"
)

type stackSpec struct {
	*stack
	stackChildren []stackChild
}

// Stack starts a chain of widgets to compose into a stack
func Stack() (out *stackSpec) {
	out = &stackSpec{stack: &stack{}}
	return
}

func (s *stackSpec) Alignment(alignment dirSpec) *stackSpec {
	s.alignment = alignment
	return s
}

// Stacked appends a widget to the stack, the stack's dimensions will be
// computed from the largest widget insetSpec the stack
func (s *stackSpec) Stacked(w Widget) (out *stackSpec) {
	s.stackChildren = append(s.stackChildren, stacked(w))
	return s
}

// Expanded lays out a widget with the same max constraints as the stack
func (s *stackSpec) Expanded(w Widget) (out *stackSpec) {
	s.stackChildren = append(s.stackChildren, expanded(w))
	return s
}

// Fn runs the ops queue configured insetSpec the stack
func (s *stackSpec) Fn(c Ctx) Dims {
	return s.layout(c, s.stackChildren...)
}

// stack lays out stackChild elements on top of each other, according to an alignment
// dir.
type stack struct {
	// alignment is the dir to align stackChildren smaller than the available space.
	alignment dirSpec
}

// stackChild represents a stackChild for a stack layout.
type stackChild struct {
	expanded bool
	widget   Widget

	// Scratch space.
	call op.CallOp
	dims Dims
}

// stacked returns a stack stackChild that is laid out with no minimum constraints and
// maximum constraints passed to stack.layout.
func stacked(w Widget) stackChild {
	return stackChild{
		widget: w,
	}
}

// expanded returns a stack stackChild with the minimum constraints set to the largest
// stacked stackChild. The maximum constraints are set to the same as passed to
// stack.layout.
func expanded(w Widget) stackChild {
	return stackChild{
		expanded: true,
		widget:   w,
	}
}

// layout a stack of stackChildren. The position of the stackChildren are determined by the
// specified order, but stacked stackChildren are laid out before expanded stackChildren.
func (s stack) layout(
	gtx Ctx,
	stackChildren ...stackChild,
) Dims {
	var maxSZ image.Point
	// First lay out stacked stackChildren.
	ct := gtx
	ct.Lim.Min = image.Point{}
	for i, w := range stackChildren {
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
		stackChildren[i].call = call
		stackChildren[i].dims = d
	}
	// Then lay out expanded stackChildren.
	for i, w := range stackChildren {
		if !w.expanded {
			continue
		}
		macro := op.Record(gtx.Ops)
		ct.Lim.Min = maxSZ
		d := w.widget(ct)
		call := macro.Stop()
		if w := d.Size.X; w > maxSZ.X {
			maxSZ.X = w
		}
		if h := d.Size.Y; h > maxSZ.Y {
			maxSZ.Y = h
		}
		stackChildren[i].call = call
		stackChildren[i].dims = d
	}

	maxSZ = gtx.Lim.Constrain(maxSZ)
	var baseline int
	for _, ch := range stackChildren {
		sz := ch.dims.Size
		var p image.Point
		switch s.alignment.Dir {
		case N, S, Center:
			p.X = (maxSZ.X - sz.X) / 2
		case NE, SE, E:
			p.X = maxSZ.X - sz.X
		}
		switch s.alignment.Dir {
		case W, Center, E:
			p.Y = (maxSZ.Y - sz.Y) / 2
		case SW, S, SE:
			p.Y = maxSZ.Y - sz.Y
		}
		stack := op.Save(gtx.Ops)
		op.Offset(ToPoint(p)).Add(gtx.Ops)
		ch.call.Add(gtx.Ops)
		stack.Load()
		if baseline == 0 {
			if b := ch.dims.Baseline; b != 0 {
				baseline = b + maxSZ.Y - sz.Y - p.Y
			}
		}
	}
	return Dims{
		Size:     maxSZ,
		Baseline: baseline,
	}
}

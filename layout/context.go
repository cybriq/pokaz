package layout

import (
	"time"

	"github.com/cybriq/giocore/f32"
	"github.com/cybriq/giocore/io/event"
	"github.com/cybriq/giocore/io/system"
	"github.com/cybriq/giocore/op"
	"github.com/cybriq/giocore/unit"
)

// Ctx carries the state needed by almost all layouts and widgets. A zero value
// Ctx never returns events, map units to pixels with a scale of 1.0, and
// returns the zero time from Time.
type Ctx struct {
	// Lim sets constraints to track the constraints for the active widget
	// or layout.
	Lim
	// Metric is the physical pixels per unit.Dp
	unit.Metric
	// By convention, a nil Queue is a signal to widgets to draw themselves
	// in a disabled state.
	event.Queue
	// Time is the animation time.
	time.Time
	// Ops are the collection of ops that accumulates during composition
	*op.Ops
	// SkipOps indicates this context is only for deriving a map of boxes
	SkipOps bool
}

// NewCtx is a shorthand for
//
//   Ctx{
//     Ops: ops,
//     Time: e.Time,
//     Queue: e.Queue,
//     Config: e.Config,
//     Lim: Exact(e.Size),
//   }
//
// NewCtx calls ops.Reset and adjusts ops for e.Insets.
func NewCtx(ops *op.Ops, e system.FrameEvent) Ctx {
	ops.Reset()

	size := e.Size

	if e.Insets != (system.Insets{}) {
		left := e.Metric.Px(e.Insets.Left)
		top := e.Metric.Px(e.Insets.Top)
		op.Offset(
			f32.Point{
				X: float32(left),
				Y: float32(top),
			},
		).Add(ops)

		size.X -= left + e.Metric.Px(e.Insets.Right)
		size.Y -= top + e.Metric.Px(e.Insets.Bottom)
	}

	return Ctx{
		Ops:    ops,
		Time:   e.Now,
		Queue:  e.Queue,
		Metric: e.Metric,
		Lim:    Exact(size),
	}
}

// Px maps the value to pixels.
func (c Ctx) Px(v unit.Value) int {
	return c.Metric.Px(v)
}

// Events return the events available for the key. If no queue is configured,
// Events returns nil.
func (c Ctx) Events(k event.Tag) []event.Event {
	if c.Queue == nil {
		return nil
	}
	return c.Queue.Events(k)
}

// Disabled returns a copy of this ctx with a nil Queue, blocking events to
// widgets using it.
//
// By convention, a nil Queue is a signal to widgets to draw themselves
// in a disabled state.
func (c Ctx) Disabled() Ctx {
	c.Queue = nil
	return c
}

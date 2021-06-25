package ctx

import (
	"time"

	"github.com/cybriq/giocore/f32"
	"github.com/cybriq/giocore/io/event"
	"github.com/cybriq/giocore/io/system"
	"github.com/cybriq/giocore/op"
	"github.com/cybriq/giocore/unit"
	"github.com/cybriq/pokaz/layout/ctx/proto"
	"github.com/cybriq/pokaz/layout/dim"
)

// Context carries the state needed by almost all layouts and widgets. A zero
// value Context never returns events, map units to pixels with a scale of 1.0,
// and returns the zero time from Now.
//
// This implementation is intentionally using value receivers as each time it
// is passed through it defines the base context for the next widget.
type Context struct {
	// Constraints track the constraints for the active widget or layout.
	constraints dim.Constraints

	metric unit.Metric
	// By convention, a nil Queue is a signal to widgets to draw themselves in a
	// disabled state.
	queue event.Queue

	// if the queue is enabled (causes widgets to render greyed out)
	enabled bool

	// Now is the animation time.
	now time.Time

	ops *op.Ops
}

var _ proto.Context = Context{}

func (c Context) Constraints() dim.Constraints {
	return c.constraints
}

func (c Context) SetConstraints(constraints dim.Constraints) {
	c.constraints = constraints
}

func (c Context) Metric() unit.Metric {
	return c.metric
}

func (c Context) Queue() event.Queue {
	return c.queue
}

func (c Context) Now() time.Time {
	return c.now
}

func (c Context) Ops() *op.Ops {
	return c.ops
}

// New is a shorthand for
//
//   Context{
//     Ops: ops,
//     Now: e.Now,
//     Queue: e.Queue,
//     Config: e.Config,
//     Constraints: Exact(e.Size),
//   }
//
// New calls ops.Reset and adjusts ops for e.Insets.
func New(ops *op.Ops, e system.FrameEvent) Context {
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

	return Context{
		ops:         ops,
		now:         e.Now,
		queue:       e.Queue,
		metric:      e.Metric,
		constraints: dim.Exact(size),
	}
}

// Px maps the value to pixels.
func (c Context) Px(v unit.Value) int {
	return c.metric.Px(v)
}

// Events return the events available for the key. If no queue is configured,
// Events returns nil.
func (c Context) Events(k event.Tag) []event.Event {
	if c.queue == nil {
		return nil
	}
	return c.queue.Events(k)
}

// Disabled returns a copy of this ctx with a nil Queue, blocking events to
// widgets using it.
//
// By convention, a nil Queue is a signal to widgets to draw themselves in a
// disabled state.
func (c Context) Disabled() proto.Context {
	c.queue = nil
	return c
}

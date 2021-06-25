package proto

import (
	"time"
	
	"github.com/cybriq/giocore/io/event"
	"github.com/cybriq/giocore/op"
	"github.com/cybriq/giocore/unit"
	"github.com/cybriq/pokaz/layout/dim"
)

// Context is an interface for Gio contexts to expand their capabilities.
//
// The base context as found in Gio's version of layout.
// Context is intentionally pass by value as there is no other data that is
// important for this layout paint and input route process.
//
// However, in order to enable such things as viewports,
// a context can serve as a mechanism to abstract the interface for the
// client side while allowing other widgets to inspect some derived value
// formed out of the composition of the layout tree,
// by capturing its dimensions results in a form that can be searched for
// presence inside a given viewport box.
type Context interface {
	Constraints() dim.Constraints
	Metric() unit.Metric
	Queue() event.Queue
	Now() time.Time
	Ops() *op.Ops
	Px(v unit.Value) int
	Events(k event.Tag) []event.Event
	Disabled() Context
}

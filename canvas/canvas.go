package canvas

import (
	"github.com/cybriq/giocore/io/system"
	"github.com/cybriq/giocore/op"
)

// A Canvas is a terminal point in a tree of Widget.
// It is used to represent both a Window and a Viewport frame.
type Canvas struct {
	// A Ctx defines the context of the top level of the canvas.
	// Canvases are controlled either by a window or a viewport widget
	Ctx
	// GetLimits is used to define the initial maximum dimensions in the
	// context
	GetLimits func() Lim
	// Root is the top widget containing the rest of the widgets to be
	// displayed on the canvas
	Root *Widget
}

// New makes a new canvas with the given ops queue, event queue and boundaries
// function
func New(ops *op.Ops, e system.FrameEvent, getLimits func() Lim) *Canvas {
	c := &Canvas{
		Ctx:       NewCtx(ops, e),
		GetLimits: getLimits,
	}
	// initialize with the passed limit return function,
	// for a window this is the window dimensions at frame time,
	// for a viewport this is the viewport clip boundary at frame time.
	c.Ctx.Lim = getLimits()
	return c
}

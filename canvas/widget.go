package canvas

import (
	"image"
	
	"github.com/cybriq/giocore/op"
)

// A Widget is a rectangular area of some specified dimensions in which there
// is scripting to read inputs and return paint operations
type Widget struct {
	// Parent widget that contains this widget.
	Parent *Widget
	// Context is the context for this widget
	// each widget has its own via value copy semantics in the original
	// implementation, so instead this will be stored in the widget also
	Ctx
	// Zero is the absolute position of the top left of the widget's
	// dimensions, if Zero is nil, Map needs to be called again
	Zero *image.Point
	// Dimensions are the dimensions of this widget, filled in by Map(). If
	// nil, Map() has not been run or is invalidated - set to nil to
	// invalidate it - painter will force rerunning of mapper
	Dimensions *Dims
	// mapper is the function loaded that computes the size of the
	// widget by running its child widgets mapper functions
	mapper func(w *Widget) (d Dims)
	// painter returns the ops defined by the map and set by each widget in
	// its rendering function
	painter func(w *Widget)
	// Children are the widgets referred to in the functions above
	Children WidgetChildren
	// State is a type specific data structure used in specification and
	// state functionality
	State interface{}
}

type WidgetChildren map[string]*Widget

func NewWidget(
	mapper func(w *Widget) (d Dims),
	painter func(w *Widget),
) *Widget {
	return &Widget{
		Children: make(WidgetChildren),
		mapper:   mapper,
		painter:  painter,
	}
}

// Map calls map on all the children and computes a widget dimension. This calls
// Map on all WidgetChildren and their WidgetChildren etc., and sets it in
// Dimensions above
func (w *Widget) Map() (d *Dims) {
	for i := range w.Children {
		w.Children[i].Map()
	}
	dd := w.mapper(w)
	w.Dimensions = &dd
	return w.Dimensions
}

// Paint returns the ops defined by the Widget and set by each Widget in its
// rendering function
func (w *Widget) Paint() (ops *op.Ops) {
	if w.Zero == nil || w.Dimensions == nil {
		w.Map()
	}
	w.painter(w)
	return w.Ops
}

// Invalidate sets the paint function to regenerate completely next time it is
// called.
func (w *Widget) Invalidate() {
	// nil all the cached values
	w.Zero = nil
	w.Dimensions = nil
	// propagate the invalidation down through the widget's tree
	for i := range w.Children {
		w.Children[i].Invalidate()
	}
}

// ResetContext copies a new context to all child widgets, this needs to be
// called prior to calling Map or Paint at the beginning of a frame
func (w *Widget) ResetContext(ctx Ctx) {
	w.Ctx = ctx
	for i := range w.Children {
		w.Children[i].ResetContext(ctx)
	}
	
}

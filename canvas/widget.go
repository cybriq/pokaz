package canvas

import "github.com/cybriq/giocore/op"

// A Widget is a rectangular area of some specified dimensions in which there
// is scripting to read inputs and return paint operations
type Widget struct {
	// Parent widget that contains this widget.
	Parent *Widget
	// Context is the context for this widget
	// each widget has its own via value copy semantics in the original
	// implementation, so instead this will be stored in the widget also
	Ctx
	// Dimensions are the dimensions of this widget, filled in by Map(). If
	// nil, Map() has not been run or is invalidated - set to nil to
	// invalidate it - Render will force rerunning of Outlines
	Dimensions *Dims
	// Outlines is the function loaded that computes the size of the
	// widget by running its child widgets Outlines functions
	Outlines func(w *Widget) (d Dims)
	// Render returns the ops defined by the map and set by each widget in
	// its rendering function
	Render func(w *Widget) (ops *op.Ops)
	// Children are the widgets referred to in the functions above
	Children []*Widget
}

// Map calls map on all the children and computes a widget dimension.
// This calls Map on all Children and their Children etc, and sets it in
// Dimensions above
func (w *Widget) Map() (d Dims) {
	return w.Outlines(w)
}

// Paint returns the ops defined by the Widget and set by each Widget in
// its rendering function
func (w *Widget) Paint() (ops *op.Ops) {
	return w.Render(w)
}

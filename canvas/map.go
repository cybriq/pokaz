package canvas

import "github.com/cybriq/giocore/op"

type Widget struct {
	// Parent widget that contains this widget. If the parent is nil,
	// this widget's boundaries are applied as a clip over every child
	// Paint() call
	Parent *Widget
	// Children are the widgets inside this widget
	Children []*Widget
	// Dimensions are the dimensions of this widget, filled in by Map()
	Dimensions Dims
	// Map calls map on all the children and computes a widget dimension.
	// This calls Map on all Children and their Children etc, and sets it in
	// Dims above
	Map func() Dims
	// Paint returns the ops defined by the map and set by each widget in
	// its rendering function
	Paint func() (ops *op.Ops)
}

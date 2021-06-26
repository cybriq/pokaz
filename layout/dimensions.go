package layout

import "image"

// Dims define the dimensions of a widget, including a baseline for lining
// up text
type Dims struct {
	Size     image.Point
	Baseline int
}

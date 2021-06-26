package layout

import "image"

// Dimensions define the dimensions of a widget, including a baseline for lining
// up text
type Dimensions struct {
	Size     image.Point
	Baseline int
}

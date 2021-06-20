package flex

import (
	"github.com/cybriq/pokaz/layout/align"
	"github.com/cybriq/pokaz/layout/axis"
	"github.com/cybriq/pokaz/layout/ctx"
	"github.com/cybriq/pokaz/layout/dim"
	"github.com/cybriq/pokaz/layout/wdg"
)

// Layout is a horizontal or vertical stack of widgets with fixed and expanding
// boxes
type Layout struct {
	flex
	children []child
}

// Flex creates a new flex.Layout
func Flex() (out *Layout) {
	return new(Layout)
}

// VFlex creates a new vertical flex layout
func VFlex() (out *Layout) {
	return Flex().Vertical()
}

// Vertical sets the axis to vertical, otherwise it is horizontal
func (f *Layout) Vertical() (out *Layout) {
	f.flex.axis = axis.Vertical
	return f
}

// Align sets the alignment to use on each box in the flex
func (f *Layout) Align(alignment align.Alignment) (out *Layout) {
	f.flex.alignment = alignment
	return f
}

// Space sets the spacing for the flex
func (f *Layout) Space(spc Spacing) (out *Layout) {
	f.flex.spacing = spc
	return f
}

// Rigid inserts a string of rigid widget into the flex
func (f *Layout) Rigid(w ...wdg.Widget) (out *Layout) {
	for i := range w {
		f.children = append(f.children, rigid(w[i]))
	}
	return f
}

// Flexed inserts a string of flexed widgets into the flex
func (f *Layout) Flexed(weight float32, w ...wdg.Widget) (out *Layout) {
	for i := range w {
		f.children = append(f.children, flexed(weight/float32(len(w)), w[i]))
	}
	return f
}

// Fn runs the ops in the context using the FlexChildren inside it
func (f *Layout) Fn(c ctx.Context) dim.Dimensions {
	return f.flex.layout(c, f.children...)
}

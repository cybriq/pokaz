package flex

import (
	"github.com/cybriq/pokaz/align"
	"github.com/cybriq/pokaz/axis"
	"github.com/cybriq/pokaz/ctx"
	"github.com/cybriq/pokaz/dims"
	"github.com/cybriq/pokaz/wdg"
)

type Flex struct {
	flex
	children []child
}

// New creates a new flex layout
func New() (out *Flex) {
	return new(Flex)
}

// VFlex creates a new vertical flex layout
func (f *Flex) VFlex() (out *Flex) {
	return new(Flex).Vertical()
}

// alignment setters

func (f *Flex) Align(alignment align.Alignment) (out *Flex) {
	f.flex.alignment = alignment
	return f
}

// axis setters

// Vertical sets the axis to vertical, otherwise it is horizontal
func (f *Flex) Vertical() (out *Flex) {
	f.flex.axis = axis.Vertical
	return f
}

// spacing setters

func (f *Flex) Space(spc Spacing) (out *Flex) {
	f.flex.spacing = spc
	return f
}

// Rigid inserts a rigid widget into the flex
func (f *Flex) Rigid(w wdg.Widget) (out *Flex) {
	f.children = append(f.children, rigid(w))
	return f
}

// Flexed inserts a flexed widget into the flex
func (f *Flex) Flexed(wgt float32, w wdg.Widget) (out *Flex) {
	f.children = append(f.children, flexed(wgt, w))
	return f
}

// Fn runs the ops in the context using the FlexChildren inside it
func (f *Flex) Fn(c ctx.Context) dims.Dimensions {
	return f.flex.layout(c, f.children...)
}

package flex

import (
	"github.com/cybriq/pokaz/align"
	"github.com/cybriq/pokaz/axis"
	"github.com/cybriq/pokaz/ctx"
	"github.com/cybriq/pokaz/dims"
	"github.com/cybriq/pokaz/wdg"
)

type Flex struct {
	flex     flex
	ctx      *ctx.Context
	children []child
}

// New creates a new flex layout
func New() (out *Flex) {
	return new(Flex)
}

// VFlex creates a new vertical flex layout
func (th *Flex) VFlex() (out *Flex) {
	return new(Flex).Vertical()
}

// alignment setters

// AlignStart sets alignment for layout from Start
func (f *Flex) AlignStart() (out *Flex) {
	f.flex.alignment = align.Start
	return f
}

// AlignEnd sets alignment for layout from End
func (f *Flex) AlignEnd() (out *Flex) {
	f.flex.alignment = align.End
	return f
}

// AlignMiddle sets alignment for layout from Middle
func (f *Flex) AlignMiddle() (out *Flex) {
	f.flex.alignment = align.Middle
	return f
}

// AlignBaseline sets alignment for layout from Baseline
func (f *Flex) AlignBaseline() (out *Flex) {
	f.flex.alignment = align.Baseline
	return f
}

// axis setters

// Vertical sets axis to vertical, otherwise it is horizontal
func (f *Flex) Vertical() (out *Flex) {
	f.flex.axis = axis.Vertical
	return f
}

// spacing setters

// SpaceStart sets the corresponding flex spacing parameter
func (f *Flex) SpaceStart() (out *Flex) {
	f.flex.spacing = SpaceStart
	return f
}

// SpaceEnd sets the corresponding flex spacing parameter
func (f *Flex) SpaceEnd() (out *Flex) {
	f.flex.spacing = SpaceEnd
	return f
}

// SpaceSides sets the corresponding flex spacing parameter
func (f *Flex) SpaceSides() (out *Flex) {
	f.flex.spacing = SpaceSides
	return f
}

// SpaceAround sets the corresponding flex spacing parameter
func (f *Flex) SpaceAround() (out *Flex) {
	f.flex.spacing = SpaceAround
	return f
}

// SpaceBetween sets the corresponding flex spacing parameter
func (f *Flex) SpaceBetween() (out *Flex) {
	f.flex.spacing = SpaceBetween
	return f
}

// SpaceEvenly sets the corresponding flex spacing parameter
func (f *Flex) SpaceEvenly() (out *Flex) {
	f.flex.spacing = SpaceEvenly
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

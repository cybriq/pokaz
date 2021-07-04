package canvas

import "github.com/cybriq/giocore/op"

type Flex struct {
	*Widget
	// Axis is the main axis, either Horizontal or Vertical.
	Axis Axis
	// Spacing controls the distribution of space left after
	// layout.
	Spacing Spacing
	// Alignment is the alignment in the cross axis.
	Alignment Align
	// WeightSum is the sum of weights used for the weighted
	// size of Flexed children. If WeightSum is zero, the sum
	// of all Flexed weights is used.
	WeightSum float32
	// FlexChildren is the list of child widgets of the Flex
	FlexChildren []FlexChild
}

// FlexChild is the descriptor for a Flex child.
type FlexChild struct {
	*Widget
	weight float32
	flex   bool

	// Scratch space.
	call op.CallOp
	dims Lim
}

// FlexMap computes the dimensions of a Flex
func FlexMap(w *Widget) (d Dims) {
	fl, ok := w.State.(*Flex)
	if !ok {
		return
	}
	_ = fl
	return Dims{}
}

func FlexPaint(w *Widget) {
	fl, ok := w.State.(*Flex)
	if !ok {
		return
	}
	// do the paint thing
	_ = fl
}

func HFlex() *Flex {
	wd := NewWidget(FlexMap, FlexPaint)
	fl := &Flex{
		Axis:      Horizontal,
		Spacing:   SpaceEnd,
		Alignment: Middle,
	}
	wd.State = fl
	fl.Widget = wd
	return fl
}

func VFlex() *Flex {
	wd := NewWidget(FlexMap, FlexPaint)
	fl := &Flex{
		Axis:      Vertical,
		Spacing:   SpaceEnd,
		Alignment: Middle,
	}
	wd.State = fl
	fl.Widget = wd
	return fl

}

// alignment setters

// AlignStart sets alignment for layout from Start
func (f *Flex) AlignStart() (out *Flex) {
	f.Alignment = Start
	return f
}

// AlignEnd sets alignment for layout from End
func (f *Flex) AlignEnd() (out *Flex) {
	f.Alignment = End
	return f
}

// AlignMiddle sets alignment for layout from Middle
func (f *Flex) AlignMiddle() (out *Flex) {
	f.Alignment = Middle
	return f
}

// AlignBaseline sets alignment for layout from Baseline
func (f *Flex) AlignBaseline() (out *Flex) {
	f.Alignment = Baseline
	return f
}

// Axis setters

// Vertical sets axis to vertical, otherwise it is horizontal
func (f *Flex) Vertical() (out *Flex) {
	f.Axis = Vertical
	return f
}

// Spacing setters

// SpaceStart sets the corresponding flex spacing parameter
func (f *Flex) SpaceStart() (out *Flex) {
	f.Spacing = SpaceStart
	return f
}

// SpaceEnd sets the corresponding flex spacing parameter
func (f *Flex) SpaceEnd() (out *Flex) {
	f.Spacing = SpaceEnd
	return f
}

// SpaceSides sets the corresponding flex spacing parameter
func (f *Flex) SpaceSides() (out *Flex) {
	f.Spacing = SpaceSides
	return f
}

// SpaceAround sets the corresponding flex spacing parameter
func (f *Flex) SpaceAround() (out *Flex) {
	f.Spacing = SpaceAround
	return f
}

// SpaceBetween sets the corresponding flex spacing parameter
func (f *Flex) SpaceBetween() (out *Flex) {
	f.Spacing = SpaceBetween
	return f
}

// SpaceEvenly sets the corresponding flex spacing parameter
func (f *Flex) SpaceEvenly() (out *Flex) {
	f.Spacing = SpaceEvenly
	return f
}

// Rigid inserts a rigid widget into the flex
func (f *Flex) Rigid(w *Widget) (out *Flex) {
	f.FlexChildren = append(f.FlexChildren, FlexChild{
		Widget: w,
	})
	return f
}

// Flexed inserts a flexed widget into the flex
func (f *Flex) Flexed(wgt float32, w *Widget) (out *Flex) {
	f.FlexChildren = append(f.FlexChildren, FlexChild{
		Widget: w,
		weight: wgt,
		flex:   true,
	})
	return f
}

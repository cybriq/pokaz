package layout

import (
	"image"
	
	"github.com/cybriq/giocore/op"
)

// Flex is a horizontal or vertical stack of widgets with fixed and expanding
// boxes
type Flex struct {
	flex
	children []child
}

// HFlex creates a new Flex
func HFlex() (out *Flex) {
	return new(Flex)
}

// VFlex creates a new vertical flex layout
func VFlex() (out *Flex) {
	return HFlex().Vertical()
}

// Vertical sets the axis to vertical, otherwise it is horizontal
func (f *Flex) Vertical() (out *Flex) {
	f.flex.axis = Vertical
	return f
}

// Align sets the alignment to use on each box in the flex
func (f *Flex) Align(alignment Align) (out *Flex) {
	f.flex.alignment = alignment
	return f
}

// Spacing sets the spacing for the flex
func (f *Flex) Spacing(spc Spacing) (out *Flex) {
	f.flex.spacing = spc
	return f
}

// Rigid inserts a string of rigid widget into the flex
func (f *Flex) Rigid(w ...Widget) (out *Flex) {
	for i := range w {
		f.children = append(f.children, rigid(w[i]))
	}
	return f
}

// Flexed inserts a string of flexed widgets into the flex
func (f *Flex) Flexed(weight float32, w ...Widget) (out *Flex) {
	for i := range w {
		f.children = append(
			f.children,
			flexed(weight/float32(len(w)), w[i]),
		)
	}
	return f
}

// Fn runs the ops in the context using the FlexChildren inside it
func (f *Flex) Fn(c Ctx) Dims {
	return f.layout(c, f.children...)
}

// flex lays out child elements along an axis, according to alignment and
// weights.
type flex struct {
	// axis is the main axis, either Horizontal or Vertical.
	axis Axis
	// spacing controls the distribution of space left after layout.
	spacing Spacing
	// alignment is the alignment in the cross axis.
	alignment Align
	// weightSum is the sum of weights used for the weighted size of flexed
	// children. If weightSum is zero, the sum of all flexed weights is
	// used.
	weightSum float32
}

// child is the descriptor for a flex child.
type child struct {
	flex   bool
	weight float32
	
	widget Widget
	
	// Scratch space.
	call op.CallOp
	dims Dims
}

// Spacing determine the spacing mode for a flex.
type Spacing uint8

const (
	// SpaceEnd leaves space at the end.
	SpaceEnd Spacing = iota
	// SpaceStart leaves space at the start.
	SpaceStart
	// SpaceSides share space between the start and end.
	SpaceSides
	// SpaceAround distributes space evenly between children, with half as
	// much space at the start and end.
	SpaceAround
	// SpaceBetween distributes space evenly between children, leaving no
	// space at the start and end.
	SpaceBetween
	// SpaceEvenly distributes space evenly between children and at the
	// start and end.
	SpaceEvenly
)

// rigid returns a flex child with a maximal constraint of the remaining space.
func rigid(widget Widget) child {
	return child{
		widget: widget,
	}
}

// flexed returns a flex child forced to take up weight fraction of the space
// left over from rigid children. The fraction is weight divided by either the
// weight sum of all flexed children, or the flex weightSum if it is nonzero.
func flexed(weight float32, widget Widget) child {
	return child{
		flex:   true,
		weight: weight,
		widget: widget,
	}
}

// layout a list of children. The position of the children are determined by the
// specified order, but rigid children are laid out before flexed children.
func (f flex) layout(gtx Ctx, children ...child) Dims {
	size := 0
	cs := gtx.Lim
	mainMin, mainMax := f.axis.MainConstraint(cs)
	crossMin, crossMax := f.axis.CrossConstraint(cs)
	remaining := mainMax
	var totalWeight float32
	cgtx := gtx
	// Lay out rigid children.
	for i, child := range children {
		if child.flex {
			totalWeight += child.weight
			continue
		}
		macro := op.Record(gtx.Ops)
		cgtx.Lim = f.axis.
			Constraints(0, remaining, crossMin, crossMax)
		dm := child.widget(cgtx)
		c := macro.Stop()
		sz := f.axis.Convert(dm.Size).X
		size += sz
		remaining -= sz
		if remaining < 0 {
			remaining = 0
		}
		children[i].call = c
		children[i].dims = dm
	}
	if w := f.weightSum; w != 0 {
		totalWeight = w
	}
	// fraction is the rounding error from a flex weighting.
	var fraction float32
	flexTotal := remaining
	// Lay out flexed children.
	for i, child := range children {
		if !child.flex {
			continue
		}
		var flexSize int
		if remaining > 0 && totalWeight > 0 {
			// Apply weight and add any leftover fraction from a
			// previous flexed.
			childSize := float32(flexTotal) * child.weight / totalWeight
			flexSize = int(childSize + fraction + .5)
			fraction = childSize - float32(flexSize)
			if flexSize > remaining {
				flexSize = remaining
			}
		}
		macro := op.Record(gtx.Ops)
		cgtx.Lim = f.axis.Constraints(flexSize, flexSize, crossMin,
			crossMax)
		dm := child.widget(cgtx)
		c := macro.Stop()
		sz := f.axis.Convert(dm.Size).X
		size += sz
		remaining -= sz
		if remaining < 0 {
			remaining = 0
		}
		children[i].call = c
		children[i].dims = dm
	}
	var maxCross int
	var maxBaseline int
	for _, child := range children {
		if c := f.axis.Convert(child.dims.Size).Y; c > maxCross {
			maxCross = c
		}
		if b := child.dims.Size.Y - child.dims.Baseline; b > maxBaseline {
			maxBaseline = b
		}
	}
	var space int
	if mainMin > size {
		space = mainMin - size
	}
	var mainSize int
	switch f.spacing {
	case SpaceSides:
		mainSize += space / 2
	case SpaceStart:
		mainSize += space
	case SpaceEvenly:
		mainSize += space / (1 + len(children))
	case SpaceAround:
		if len(children) > 0 {
			mainSize += space / (len(children) * 2)
		}
	}
	for i, child := range children {
		dm := child.dims
		b := dm.Size.Y - dm.Baseline
		var cross int
		switch f.alignment {
		case End:
			cross = maxCross - f.axis.Convert(dm.Size).Y
		case Middle:
			cross = (maxCross - f.axis.Convert(dm.Size).Y) / 2
		case Baseline:
			if f.axis == Horizontal {
				cross = maxBaseline - b
			}
		}
		stack := op.Save(gtx.Ops)
		pt := f.axis.Convert(image.Pt(mainSize, cross))
		op.Offset(ToPoint(pt)).Add(gtx.Ops)
		child.call.Add(gtx.Ops)
		stack.Load()
		mainSize += f.axis.Convert(dm.Size).X
		if i < len(children)-1 {
			switch f.spacing {
			case SpaceEvenly:
				mainSize += space / (1 + len(children))
			case SpaceAround:
				if len(children) > 0 {
					mainSize += space / len(children)
				}
			case SpaceBetween:
				if len(children) > 1 {
					mainSize += space / (len(children) - 1)
				}
			}
		}
	}
	switch f.spacing {
	case SpaceSides:
		mainSize += space / 2
	case SpaceEnd:
		mainSize += space
	case SpaceEvenly:
		mainSize += space / (1 + len(children))
	case SpaceAround:
		if len(children) > 0 {
			mainSize += space / (len(children) * 2)
		}
	}
	sz := f.axis.Convert(image.Pt(mainSize, maxCross))
	return Dims{Size: sz, Baseline: sz.Y - maxBaseline}
}

func (s Spacing) String() string {
	switch s {
	case SpaceEnd:
		return "End"
	case SpaceStart:
		return "Start"
	case SpaceSides:
		return "Sides"
	case SpaceAround:
		return "Around"
	case SpaceBetween:
		return "Around"
	case SpaceEvenly:
		return "Evenly"
	default:
		panic("unreachable")
	}
}

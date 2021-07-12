package canvas

import (
	"image"
	
	"github.com/cybriq/giocore/op"
)

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
	call  op.CallOp
	dims   Dims
	offset image.Point
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

// Rigid inserts rigid widgets into the flex
func (f *Flex) Rigid(w ...*Widget) (out *Flex) {
	for i := range w {
		f.FlexChildren = append(f.FlexChildren, FlexChild{
			Widget: w[i],
		})
	}
	return f
}

// Flexed inserts flexed widgets into the flex
func (f *Flex) Flexed(wgt float32, w ...*Widget) (out *Flex) {
	for i := range w {
		f.FlexChildren = append(f.FlexChildren, FlexChild{
			Widget: w[i],
			weight: wgt,
			flex:   true,
		})
	}
	return f
}

// FlexMap computes the dimensions of a Flex
func FlexMap(w *Widget) (d Dims) {
	fl, ok := w.State.(*Flex)
	if !ok {
		return
	}
	size := 0
	cs := fl.Ctx.Lim
	mainMin, mainMax := fl.Axis.MainConstraint(cs)
	crossMin, crossMax := fl.Axis.CrossConstraint(cs)
	remaining := mainMax
	var totalWeight float32
	// Lay out rigid children.
	for _, i := range fl.FlexChildren {
		if i.flex {
			totalWeight += i.weight
			continue
		}
		fl.Ctx.Lim = fl.Axis.
			Constraints(0, remaining, crossMin, crossMax)
		dm := i.Map()
		sz := fl.Axis.Convert(dm.Size).X
		size += sz
		remaining -= sz
		if remaining < 0 {
			remaining = 0
		}
		i.dims = *dm
	}
	if fl.WeightSum != 0 {
		totalWeight = fl.WeightSum
	}
	// fraction is the rounding error from a flex weighting.
	var fraction float32
	flexTotal := remaining
	// Lay out flexed children.
	for i, child := range fl.FlexChildren {
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
		fl.Ctx.Lim = fl.Axis.
			Constraints(flexSize, flexSize, crossMin, crossMax)
		dm := child.Map()
		sz := fl.Axis.Convert(dm.Size).X
		size += sz
		remaining -= sz
		if remaining < 0 {
			remaining = 0
		}
		fl.FlexChildren[i].dims = *dm
	}
	var maxCross int
	var maxBaseline int
	for _, child := range fl.FlexChildren {
		if c := fl.Axis.Convert(child.dims.Size).Y; c > maxCross {
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
	switch fl.Spacing {
	case SpaceSides:
		mainSize += space / 2
	case SpaceStart:
		mainSize += space
	case SpaceEvenly:
		mainSize += space / (1 + len(fl.FlexChildren))
	case SpaceAround:
		if len(fl.FlexChildren) > 0 {
			mainSize += space / (len(fl.FlexChildren) * 2)
		}
	}
	for i, child := range fl.FlexChildren {
		dm := child.dims
		b := dm.Size.Y - dm.Baseline
		var cross int
		switch fl.Alignment {
		case End:
			cross = maxCross - fl.Axis.Convert(dm.Size).Y
		case Middle:
			cross = (maxCross - fl.Axis.Convert(dm.Size).Y) / 2
		case Baseline:
			if fl.Axis == Horizontal {
				cross = maxBaseline - b
			}
		}
		child.offset = fl.Axis.Convert(image.Pt(mainSize, cross))
		mainSize += fl.Axis.Convert(dm.Size).X
		if i < len(fl.FlexChildren)-1 {
			switch fl.Spacing {
			case SpaceEvenly:
				mainSize += space / (1 + len(fl.FlexChildren))
			case SpaceAround:
				if len(fl.FlexChildren) > 0 {
					mainSize += space / len(fl.FlexChildren)
				}
			case SpaceBetween:
				if len(fl.FlexChildren) > 1 {
					mainSize += space / (len(fl.FlexChildren) - 1)
				}
			}
		}
	}
	switch fl.Spacing {
	case SpaceSides:
		mainSize += space / 2
	case SpaceEnd:
		mainSize += space
	case SpaceEvenly:
		mainSize += space / (1 + len(fl.FlexChildren))
	case SpaceAround:
		if len(fl.FlexChildren) > 0 {
			mainSize += space / (len(fl.FlexChildren) * 2)
		}
	}
	sz := fl.Axis.Convert(image.Pt(mainSize, maxCross))
	return Dims{Size: sz, Baseline: sz.Y - maxBaseline}
}

func FlexPaint(w *Widget) {
	fl, ok := w.State.(*Flex)
	if !ok {
		return
	}
	// do the paint thing
	_ = fl
}

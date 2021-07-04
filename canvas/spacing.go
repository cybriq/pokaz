package canvas

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

func (s Spacing) String() string {
	switch s {
	case SpaceEnd:
		return "SpaceEnd"
	case SpaceStart:
		return "SpaceStart"
	case SpaceSides:
		return "SpaceSides"
	case SpaceAround:
		return "SpaceAround"
	case SpaceBetween:
		return "SpaceAround"
	case SpaceEvenly:
		return "SpaceEvenly"
	default:
		panic("unreachable")
	}
}

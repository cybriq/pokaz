package layout

type Align uint8

const (
	Start Align = iota
	End
	Middle
	Baseline
	endAlign
)

var aligners = []string{
	"Start", "End", "Middle", "Baseline",
}

func (a Align) String() string {
	if a < 0 || a >= endAlign {
		panic("alignment value out of range")
	}
	return aligners[a]
}

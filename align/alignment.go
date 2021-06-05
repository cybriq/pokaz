package align

type Alignment uint8

const (
	Start Alignment = iota
	End
	Middle
	Baseline
	endAlign
)

var aligners = []string{
	"Start", "End", "Middle", "Baseline",
}

func (a Alignment) String() string {
	if a < 0 || a >= endAlign {
		panic("alignment value out of range")
	}
	return aligners[a]
}

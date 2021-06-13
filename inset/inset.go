package inset

import (
	"github.com/cybriq/giocore/unit"
	"github.com/cybriq/pokaz/ctx"
	"github.com/cybriq/pokaz/dims"
	"github.com/cybriq/pokaz/wdg"
)

type Inset struct {
	in    inset
	w     wdg.Widget
}

// New creates a padded empty space around a widget
func New(pad float32, embed wdg.Widget) (out *Inset) {
	out = &Inset{
		in:     uniform(unit.Sp(pad)),
		w:      embed,
	}
	return
}

// Embed sets the widget that will be inside the inset
func (in *Inset) Embed(w wdg.Widget) *Inset {
	in.w = w
	return in
}

// Fn lays out the given widget with the configured context and padding
func (in *Inset) Fn(c ctx.Context) dims.Dimensions {
	return in.in.layout(c, in.w)
}


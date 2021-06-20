package inset

import (
	"github.com/cybriq/giocore/unit"
	"github.com/cybriq/pokaz/ctx"
	"github.com/cybriq/pokaz/dim"
	"github.com/cybriq/pokaz/wdg"
)

type _inset struct {
	in inset
	w  wdg.Widget
}

// New creates a padded empty space around a widget
func New(pad float32, embed wdg.Widget) (out *_inset) {
	out = &_inset{
		in: uniform(unit.Sp(pad)),
		w:  embed,
	}
	return
}

// Embed sets the widget that will be inside the inset
func (in *_inset) Embed(w wdg.Widget) *_inset {
	in.w = w
	return in
}

// Fn lays out the given widget with the configured context and padding
func (in *_inset) Fn(c ctx.Context) dim.Dimensions {
	return in.in.layout(c, in.w)
}

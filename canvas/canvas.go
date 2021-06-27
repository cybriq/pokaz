package canvas

import (
	"github.com/cybriq/giocore/io/system"
	"github.com/cybriq/giocore/op"
)

type Canvas struct {
	GetLimits func() Lim
	Ctx
	Root *Widget
}

func New(ops *op.Ops, e system.FrameEvent, getLimits func() Lim) *Canvas {
	c := &Canvas{
		Ctx: NewCtx(ops, e),
		GetLimits: getLimits,
	}
	return c
}

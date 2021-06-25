// Package view implements a context which is used to generate and filter a
// map of the widget boxes inside a widget
package view

import (
	"time"
	
	"github.com/cybriq/giocore/io/event"
	"github.com/cybriq/giocore/op"
	"github.com/cybriq/giocore/unit"
	"github.com/cybriq/pokaz/layout/ctx"
	"github.com/cybriq/pokaz/layout/ctx/proto"
	"github.com/cybriq/pokaz/layout/dim"
)

type Port struct{
	ctx.Context
}

func NewViewport() *Port {
	return &Port{}
}

func (p *Port) Constraints() dim.Constraints {
	panic("implement me")
}

func (p *Port) Metric() unit.Metric {
	panic("implement me")
}

func (p *Port) Queue() event.Queue {
	panic("implement me")
}

func (p *Port) Now() time.Time {
	panic("implement me")
}

func (p *Port) Ops() *op.Ops {
	panic("implement me")
}

func (p *Port) Px(v unit.Value) int {
	panic("implement me")
}

func (p *Port) Events(k event.Tag) []event.Event {
	panic("implement me")
}

func (p *Port) Disabled() proto.Context {
	panic("implement me")
}

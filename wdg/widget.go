package wdg

import (
	"github.com/cybriq/pokaz/ctx"
	"github.com/cybriq/pokaz/dim"
)

// Widget is a generic function that adds some kind of ops to a layout rectangle
type Widget func(gtx ctx.Context) dim.Dimensions


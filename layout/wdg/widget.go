package wdg

import (
	"github.com/cybriq/pokaz/layout/ctx"
	"github.com/cybriq/pokaz/layout/dim"
)

// Widget is a generic function that adds some kind of ops to a layout rectangle
type Widget func(gtx ctx.Context) dim.Dimensions

package wdg

import (
	"github.com/cybriq/pokaz/ctx"
	"github.com/cybriq/pokaz/dims"
)

// Widget is a generic function that adds some kind of ops to a layout rectangle
type Widget func(gtx ctx.Context) dims.Dimensions


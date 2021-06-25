package wdg

import (
	"github.com/cybriq/pokaz/layout/ctx/proto"
	"github.com/cybriq/pokaz/layout/dim"
)

// Widget is a generic function that adds some kind of ops to a layout rectangle
type Widget func(gtx proto.Context) dim.Dimensions

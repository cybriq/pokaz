package widget

import (
	"github.com/cybriq/pokaz/context"
	"github.com/cybriq/pokaz/dimensions"
)

type Widget func(gtx context.Context) dimensions.Dimensions

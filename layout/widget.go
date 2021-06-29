package layout

// Widget is a generic function that adds some kind of ops to a layout rectangle
type Widget func(gtx Ctx) Dims

// Widgeter defines a widget interface
type Widgeter interface {
	Fn(Ctx) Dims
}


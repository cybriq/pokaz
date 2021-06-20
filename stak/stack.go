package stak

import (
	"github.com/cybriq/pokaz/ctx"
	"github.com/cybriq/pokaz/dim"
	"github.com/cybriq/pokaz/dir"
	"github.com/cybriq/pokaz/wdg"
)

type _stack struct {
	*stack
	children []child
}

// Stack starts a chain of widgets to compose into a stack
func Stack() (out *_stack) {
	out = &_stack{stack: &stack{}}
	return
}

func (s *_stack) Alignment(alignment dir.Direction) *_stack {
	s.alignment = alignment
	return s
}

// Stacked appends a widget to the stack, the stack's dimensions will be
// computed from the largest widget in the stack
func (s *_stack) Stacked(w wdg.Widget) (out *_stack) {
	s.children = append(s.children, stacked(w))
	return s
}

// Expanded lays out a widget with the same max constraints as the stack
func (s *_stack) Expanded(w wdg.Widget) (out *_stack) {
	s.children = append(s.children, expanded(w))
	return s
}

// Fn runs the ops queue configured in the stack
func (s *_stack) Fn(c ctx.Context) dim.Dimensions {
	return s.stack.layout(c, s.children...)
}

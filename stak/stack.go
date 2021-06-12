package stak

import (
	"github.com/cybriq/pokaz/ctx"
	"github.com/cybriq/pokaz/dims"
	"github.com/cybriq/pokaz/dir"
	"github.com/cybriq/pokaz/wdg"
)

type Stack struct {
	*stack
	children []child
}

// New starts a chain of widgets to compose into a stack
func New() (out *Stack) {
	out = &Stack{stack: &stack{}}
	return
}

func (s *stack) Alignment(alignment dir.Direction) *stack {
	s.alignment = alignment
	return s
}

// functions to chain widgets to stack (first is lowest last highest)

// Stacked appends a widget to the stack, the stack's dimensions will be
// computed from the largest widget in the stack
func (s *Stack) Stacked(w wdg.Widget) (out *Stack) {
	s.children = append(s.children, stacked(w))
	return s
}

// Expanded lays out a widget with the same max constraints as the stack
func (s *Stack) Expanded(w wdg.Widget) (out *Stack) {
	s.children = append(s.children, expanded(w))
	return s
}

// Fn runs the ops queue configured in the stack
func (s *Stack) Fn(c ctx.Context) dims.Dimensions {
	return s.stack.layout(c, s.children...)
}

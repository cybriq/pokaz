// Package canvas provides an integrated system for tracking a render queue
// that keeps a map of the boundaries of the widgets to enable creating
// viewports.
//
// General Principles
//
// In the Gio GUI programming system, there is context and the context is
// populated by running a closure containing a composition tree.
//
// What canvas does is roll the two things together so that a single common
// embedded command adds to the map after the dimensions are computed; and
// furthermore, the implementation code is revised to separate generating the
// dimensions and appending the ops to the op tree filtering the render tree.
//
// A canvas has a function that tells it its bounds, which for a window would be
// loaded from the driver, for a viewport, there can be rigid and fixed display
// pixel width, and automatically ignoring in-axis expansion directives in child
// widgets, something that is also only possible with a controlling type such as
// canvas.
//
// Implementation
//
// Widgets are thus a tree structure and to specify, one uses generator
// functions that inject a pair of functions that compute dimensions and
// assemble the ops required.
//
// These widget generator functions form a systematic ordering that specifies
// what the hierarchy of widget's shape is thus making it possible to derive the
// bounds of a widget by walking the widget tree, as well as using this tree to
// filter a region of the tree to derive a render queue for a viewport,
// defining a new root and clip boundary as specified by the viewport widget.
//
// This design also makes possible and preferable the implementation of text
// widgets simpler by making each widget glyph a first class citizen also
// thus potentially simplifying rich text layout with a wrapper type that
// contains the text flowed into the rectangle.
//
// A note to remind when implementing text shaping there will be a higher
// level abstraction to cover strings of text including text flow, left,
// right and center positioning within the box,
// and RTL encoding of course.
//
// Likely it will be better to flow text as word-based units so kerning is
// done correctly, splicing punctuation and spacing between them as required.
package canvas

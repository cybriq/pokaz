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
// loaded from the driver, for a viewport, there can be rigid (expand to minimum
// of all elements) and fixed display pixel width, and automatically ignoring
// in-axis expansion directives in child widgets, something that is also only
// possible with a controlling type such as canvas.
package canvas


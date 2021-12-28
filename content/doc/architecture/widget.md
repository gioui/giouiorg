---
title: Widget
subtitle: Reusable and composable parts
---

We've been mentioning widgets for quite a while now. In principle widgets are composable and drawable UI elements that may react to input. More concretely:

* They get input from an [`Queue`](https://gioui.org/io/system#FrameEvent.Queue).
* They might hold some state.
* They calculate their size given constraints.
* They draw themselves to an [`op.Ops`](https://gioui.org/op#Ops) list.

By convention, widgets have a `Layout` method that does all of the above. Some widgets have separate methods for querying their state or to [pass events back to the program](https://gioui.org/widget#Clickable.Clicked).

Some widgets have several visual representations. For example, the stateful [Clickable](https://gioui.org/widget#Clickable) is used as basis for [buttons](https://gioui.org/widget/material#ButtonStyle.Layout) and [icon buttons](https://gioui.org/widget/material#IconButtonStyle.Layout). In fact, the [material package](https://gioui.org/widget/material) implements only the [Material Design](https://material.io) and is intended to be supplemented by other packages implementing different designs.


## Context

To build out more complex UI from these primitives we need some structure that describes the layout in a composable way.

It's possible to specify a layout statically, but display sizes vary greatly, so we need to be able to calculate the layout dynamically - that is constrain the available display size and then calculate the rest of the layout. We also need a comfortable way of passing events through the composed structure and similarly we need a way to pass [`op.Ops`](https://gioui.org/op#Ops) through the system.

[`layout.Context`](https://gioui.org/layout#Context) conveniently bundles these aspects together. It carries the state that is needed by almost all layouts and widgets.

To summarise the terminology:

* [`Constraints`](https://gioui.org/layout#Context.Constraints) are an "incoming" parameter to a widget. The constraints hold a widget's maximum (and minimum) size.
* [`Ops`](https://gioui.org/layout#Context.Ops) holds the generated draw operations.
* [`Events`](https://gioui.org/layout#Context.Events) holds events generated since the last drawing operation.

By convention, functions that accept a `layout.Context` return [`layout.Dimensions`](https://gioui.org/layout#Dimensions) which provides both the dimensions of the laid-out widget and the baseline of any text content within that widget.

<{{files/architecture/main.go}}[/START CONTEXTLOOP OMIT/,/END CONTEXTLOOP OMIT/]


## Custom

As an example, here is how to implement a very simple button.

Let's start by drawing it:

<{{files/architecture/button.go}}[/START VISUAL OMIT/,/END VISUAL OMIT/]

<pre style="min-height: 100px" data-run="wasm" data-pkg="architecture" data-args="button-visual" data-size="200x100"></pre>

Then handle pointer clicks:

<{{files/architecture/button.go}}[/START FINAL OMIT/,/END FINAL OMIT/]

<pre style="min-height: 100px" data-run="wasm" data-pkg="architecture" data-args="button" data-size="200x100"></pre>

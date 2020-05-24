---
title: Architecture
---

## Introduction

Gio is a library for implementing [immediate mode user interfaces](https://eliasnaur.com/blog/immediate-mode-gui-programming). This approach can be implemented in multiple ways, however the overarching similarity is that the program:

1. Listens for events such as mouse or keyboard input.
2. Updates its internal state based on the event.
3. Runs code that lays out and redraws the user interface state.

A minimal immediate mode command-line UI in pseudo-code:

```
main() {
	checked = false
	for every keypress {
		clear screen
		layoutCheckbox(keypress, &checked)
		if checked {
			print("info")
		}
	}
}

layoutCheckbox(keypress, checked) {
	if keypress == SPACE {
		*checked = !*checked
	}

	if *checked {
		print("[x]")
	} else {
		print("[ ]")
	}
}
```

In the immediate mode model, the program is in control of clearing and updating the display, and directly draws widgets and handle input during the updates.

In contrast, traditional "retained mode" libraries own the widgets through implicit library-managed state, typically arranged in a tree-like structure such as a
browser's [DOM](https://en.wikipedia.org/wiki/Document_Object_Model). As a result, the program must use the facilities given by the library to manipulate
its widgets.

Actual GUI programming has several concerns in addition to the simple example above:

1. How to get the events?
2. When to redraw the state?
3. What do the widget structures look like?
4. How to track the focus?
5. How to structure the events?
6. How to communicate with the graphics card?
7. How to handle input?
8. How to draw text?
9. Where does the widget state belong?
10. And many more.

The rest of this document tries to answer how Gio does it. If you wish to know more about immediate mode UI, these references are a good start:

* https://caseymuratori.com/blog_0001
* http://sol.gfxile.net/imgui/
* http://www.johno.se/book/imgui.html
* https://github.com/ocornut/imgui
* https://eliasnaur.com/blog/immediate-mode-gui-programming


## Window

Since a GUI library needs to talk to some sort of display system to display information:

<{{files/architecture/main.go}}[/START DRAWLOOP OMIT/,/END DRAWLOOP OMIT/]

[`app.NewWindow`](http://gioui.org/app#NewWindow) chooses the appropriate "driver" depending on the environment and build context. It might choose Wayland, Win32, Cocoa among several others.

An `app.Window` sends events from the display system to the [`windows.Events()`](https://gioui.org/app#Window.Events) channel. The system events are listed in [`gioui.org/io/system`](https://gioui.org/io/system). The input events, such as [`gioui.org/io/pointer`](https://gioui.org/io/pointer) and [`gioui.org/io/key`](https://gioui.org/io/key), are also sent into that channel.


## Operations

All UI libraries need a way for the program to specify what to display and how to handle events. Gio programs use operations, serialized into one or more [`op.Ops`](https://gioui.org/op#Ops) operation lists. Operation lists are in turn passed to the window driver through the [`FrameEvent.Frame`](https://gioui.org/io/system#FrameEvent.Frame) function.

By convention, each operation kind is represented by a Go type with an `Add` method that records the operation into the `Ops` argument. Like any Go struct literal, zero-valued fields can be useful to represent optional values.

For example, recording an operation that sets the current color to red:

<{{files/architecture/draw.go}}[/START OPERATIONS OMIT/,/END OPERATIONS OMIT/]

You might be thinking that it would be more usual to have an `ops.Add(ColorOp{Color: red})` method instead of using `op.ColorOp{Color: red}.Add(ops)`. It's like this so that the `Add` method doesn't have to take an interface-typed argument, which would often require an allocation to call. This is a key aspect of Gio's "zero allocation" design.


## Drawing

The [`paint`](https://gioui.org/op/paint) package provides operations for drawing graphics.

Coordinates are based on the top-left corner, although it's possible to [transform the coordinate system](https://gioui.org/op#TransformOp). This means `f32.Point{X:0, Y:0}` is the top left corner of the window. All drawing operations use pixel units, see [Units](#Units) section for more information.

For example, the following code will draw a 10x10 pixel colored rectangle at the top level corner of the window:

<{{files/architecture/draw.go}}[/START DRAWING OMIT/,/END DRAWING OMIT/]

<pre style="min-height: 100px" data-run="wasm" data-pkg="architecture" data-args="draw-paint" data-size="200x100"></pre>

### Transformation

Operation [`op.TransformOp`](https://gioui.org/op#TransformOp) translates the position of the operations that come after it.

For example, the following function offsets the red rectangle 100 pixels to the right:

<{{files/architecture/draw.go}}[/START TRANSFORMATION OMIT/,/END TRANSFORMATION OMIT/]

<pre style="min-height: 100px" data-run="wasm" data-pkg="architecture" data-args="draw-transformation" data-size="200x100"></pre>

Note: in the future, TransformOp will allow other transformations such as scaling and rotation.

### Clipping

In some cases we want the drawing to confined to a non-rectangular shape, for example to avoid overlapping drawings. Package [`gioui.org/op/clip`](https://gioui.org/op/clip) provides exactly that.

[`clip.Rect`](https://gioui.org/op/clip#Rect) clips all subsequent drawing operations to a rectangle with rounded corners. This is useful as a basis for a button background:

<{{files/architecture/draw.go}}[/START CLIPPING OMIT/,/END CLIPPING OMIT/]

<pre style="min-height: 100px" data-run="wasm" data-pkg="architecture" data-args="draw-clip" data-size="200x100"></pre>

Note: that we first need to get the actual operation for the clipping with `Op` before calling `Add`. This level of indirection is useful if we want to use the same clipping operation multiple times. Under the hood, Op records a [macro](#Macros) that encodes the clipping path.

For more complex clipping [`clip.Path`](https://gioui.org/op/clip#Path) can express shapes built from lines and bézier curves. This example draws a triangle with a curved edge:

<{{files/architecture/draw.go}}[/START CLIP TRIANGLE OMIT/,/END CLIP TRIANGLE OMIT/]

<pre style="min-height: 100px" data-run="wasm" data-pkg="architecture" data-args="draw-clip-triangle" data-size="200x100"></pre>

### Push and Pop

Some operations affect all operations that follow them. For example, [`paint.ColorOp`](https://gioui.org/op/paint#ColorOp) sets the "brush" color that is used in subsequent [`op.PaintOp`](https://gioui.org/op/paint#PaintOp) operations. This drawing context also includes coordinate transformation (set by [`op.TransformOp`](https://gioui.org/op#TransformOp)) and clipping (set by [`clip.ClipOp`](https://gioui.org/op/clip#ClipOp)).

We often need to set up some drawing context and then restore it to its previous state, leaving later operations unaffected. We can use [`op.StackOp`](https://gioui.org/op#StackOp) to do this. A Push operation saves the current drawing context; a Pop operation restores it.

For example, the `clipButtonOutline` function in the previous section has the unfortunate side-effect of clipping all later operations to the outline of the button background! Let's make a version of it that doesn't affect any callers:

<{{files/architecture/draw.go}}[/START STACK OMIT/,/END STACK OMIT/]

<pre style="min-height: 100px" data-run="wasm" data-pkg="architecture" data-args="draw-stack" data-size="200x100"></pre>

### Drawing Order and Macros

Drawing happens from back to front. In this function the green rectangle is drawn on top of red rectangle:

<{{files/architecture/draw.go}}[/START DRAWORDER OMIT/,/END DRAWORDER OMIT/]

<pre style="min-height: 100px" data-run="wasm" data-pkg="architecture" data-args="draw-draworder" data-size="200x100"></pre>

Sometimes you may want to change this order. For example, you may want to delay drawing to apply a transform that is calculated during drawing, or you may want to perform a list of operations several times. For this purpose there is [op.MacroOp](https://gioui.org/op#MacroOp).

<!-- TODO: Use a more practical example here. -->

<{{files/architecture/draw.go}}[/START MACRO OMIT/,/END MACRO OMIT/]

<pre style="min-height: 100px" data-run="wasm" data-pkg="architecture" data-args="draw-macro" data-size="200x100"></pre>

### Animation

Gio only issues FrameEvents when the window is resized or the user interacts with the window. However, animation requires continuous redrawing until the animation is completed. For that there is [`op.InvalidateOp`](https://gioui.org/op#InvalidateOp).

The following code will animate a green "progress bar" that fills up from left to right over 5 seconds from when the program starts:

<{{files/architecture/draw.go}}[/START ANIMATION OMIT/,/END ANIMATION OMIT/]

<pre style="min-height: 100px" data-run="wasm" data-pkg="architecture" data-args="draw-animation" data-size="200x100"></pre>

### Reusing operations with CallOp

While `op.MacroOp` allows you to record and replay operations on a single operation list, [`op.CallOp`](https://gioui.org/op#CallOp) allows for reuse of a separate operation list. This is useful for caching operations that are expensive to re-create, or for animating the disappearance of otherwise removed widgets:

<{{files/architecture/draw.go}}[/START CACHE OMIT/,/END CACHE OMIT/]

<pre style="min-height: 100px" data-run="wasm" data-pkg="architecture" data-args="draw-cache" data-size="200x100"></pre>

### Images

[`paint.ImageOp`](https://gioui.org/op/paint#ImageOp) is used to draw images. Like [`paint.ColorOp`](https://gioui.org/op/paint#ColorOp), it sets part of the drawing context (the "brush") that's used for subsequent [`PaintOp`](https://gioui.org/op/paint#PaintOp). [`ImageOp`](https://gioui.org/op/paint#ImageOp) is used similarly to [`ColorOp`](https://gioui.org/op/paint#ColorOp).

Note that [`image.RGBA`](https://golang.org/pkg/image#RGBA) and [`image.Uniform`](https://golang.org/pkg/image#Uniform) images are efficient and treated specially. Other [`Image`](https://golang.org/pkg/image#Image) implementations will undergo a more expensive copy and conversion to the underlying image model.

<{{files/architecture/draw.go}}[/START IMAGE OMIT/,/END IMAGE OMIT/]

<pre style="min-height: 100px" data-run="wasm" data-pkg="architecture" data-args="draw-image" data-size="200x100"></pre>

The image must not be mutated until another [`FrameEvent`](https://gioui.org/io/system#FrameEvent) happens, because the image may be read asynchronously while the frame is being drawn.


### Fonts

[`gioui.org/font`](https://gioui.org/font) contains the central registry for fonts. This helps to reduce passing fonts around throughout the program.

There is one font bundled in package [`gioui.org/font/gofont`](https://gioui.org/font/gofont), you can use [`gofont.Register()`](https://gioui.org/font/gofont#Register) to register it to the global registry.

For loading other fonts there is [`gioui.org/font/opentype`](https://gioui.org/font/opentype). After parsing the font(s) using [`opentype.Parse`](https://gioui.org/font/opentype#Parse) or [`opentype.ParseCollection`](https://gioui.org/font/opentype#ParseCollection) they can be registered with [`font.Register`](https://gioui.org/font#Register).

<!-- TODO: code example. -->


### Text

For converting strings to clip shapes there is the [`gioui.org/text`](https://gioui.org/text) package.

It contains [`text.FontRegistry`](https://gioui.org/text#FontRegistry) that implements cached string to shape conversion, with appropriate fallbacks.

In most cases you can use [`widget.Label`](https://gioui.org/widget#Label) which handles wrapping and layout constraints. Or when you are using material design then [`material.LabelStyle`](https://gioui.org/widget/material#LabelStyle).

<!-- TODO: code example. -->


## Input

Input is delivered to the widgets via a [`system.FrameEvent`](https://gioui.org/io/system#FrameEvent) through the [`Queue`](https://gioui.org/io/system#FrameEvent.Queue) field.

Some of the most common events in `FrameEvent.Queue` are:

* [`key.Event`](https://gioui.org/io/key#Event), [`key.Focus`](https://gioui.org/io/key#Focus) - for keyboard input.
* [`key.EditEvent`](https://gioui.org/io/key#EditEvent) - for text editing.
* [`pointer.Event`](https://gioui.org/io/pointer#Event) - for mouse and touch input.

The program can respond to these events however it likes - for example, by updating its local data structures or running a user-triggered action. The [`FrameEvent`](https://gioui.org/io/system#FrameEvent) is special - when the program receives a [`FrameEvent`](https://gioui.org/io/system#FrameEvent), it is responsible for updating the display by calling the [`e.Frame`](https://gioui.org/io/system#FrameEvent.Frame) function with an operation list representing the new state. These operations are generated immediately in response to the [`FrameEvent`](https://gioui.org/io/system#FrameEvent) which is the main reason that Gio is known as an "immediate mode" GUI.

Event-processors, such as [`Click`](https://gioui.org/gesture#Click) and [`Scroll`](https://gioui.org/gesture#Scroll) from [package `gioui.org/gesture`](https://gioui.org/gesture) detect higher-level actions from individual click events.

To distribute input among multiple different widgets, Gio needs to know about event handlers and their configuration. However, since the Gio framework is stateless, there's no direct way for the program to specify that.

Instead, some operations associate input event types (for example, keyboard presses) with arbitrary [tags](https://gioui.org/io/event#Tag) (interface{} values) chosen by the program. A program creates these operations when it's processing the [`FrameEvent`](https://gioui.org/io/system#FrameEvent)- input operations are operations like any other. In return, an [event.Queue](https://gioui.org/io/event#Queue) supplies the events that arrived since the last frame, separated by tag.

The following example demonstrates pointer input handling:

<{{files/architecture/button.go}}[/START LOWLEVEL OMIT/,/END LOWLEVEL OMIT/]

<pre style="min-height: 100px" data-run="wasm" data-pkg="architecture" data-args="button-low" data-size="200x100"></pre>

It's convenient to use a Go pointer value for the input tag, as it's cheap to convert a pointer to an interface{} and it's easy to make the value specific to a local data structure, which avoids the risk of tag conflict.

For more details take a look at [`gioui.org/io/pointer`](https://gioui.org/io/pointer) (pointer/mouse events) and [`gioui.org/io/key`](https://gioui.org/io/key) (keyboard events).


## Handling external state changes

A single frame consists of getting input, registering for input and drawing the new state:

<{{files/architecture/main.go}}[/START DRAWQUEUELOOP OMIT/,/END DRAWQUEUELOOP OMIT/]

Let's make the button change it's position every second. We can use a select to wait for events from the window and the external source at the same time. We'll use a [`Ticker`](https://golang.org/pkg/time#Ticker) as an example external change. Once we have modified the state we need to notify the window to retrigger rendering with [`w.Invalidate()`](https://gioui.org/app#Window.Invalidate).

<{{files/architecture/external.go}}[/START LOOP OMIT/,/END LOOP OMIT/]

<pre style="min-height: 100px" data-run="wasm" data-pkg="architecture" data-args="external-changes" data-size="200x100"></pre>

Writing a program using these concepts could get really verbose, which is why Gio provides standard widgets for common look and behaviour. Most programs end up using widgets primarily and few low-level operations.


## Widget

We've been mentioning widgets quite a while now. In principle widgets are composable and drawable UI elements that may react to input. Or to put more concretely.

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
* [`Dimensions`](https://gioui.org/layout#Context.Dimensions) are an "outgoing" return value from a widget, used for tracking or returning the most recent layout size.
* [`Ops`](https://gioui.org/layout#Context.Ops) holds the generated draw operations.
* [`Events`](https://gioui.org/layout#Context.Events) holds events generated since the last drawing operation.

<{{files/architecture/main.go}}[/START CONTEXTLOOP OMIT/,/END CONTEXTLOOP OMIT/]


## Units

Drawing operations use pixel coordinates, ignoring any transformation applied. However, for most use-cases you don't want to tie your user-interface sizes and positions to screen pixels. People may have screen-scaling enabled and pixel densities vary significantly between devices.

In addition to the physical pixel, package [`gioui.org/unit`](https://gioui.org/unit) implements device independent units:

* [`Px`](https://gioui.org/unit#Px) - device dependent pixel. One Px is a pixel on the screen.
* [`Dp`](https://gioui.org/unit#Dp) - device independent pixel. Takes into account screen-density and the screen-scaling settings.
* [`Sp`](https://gioui.org/unit#Sp) - device independent pixel for text. An Sp is like a Dp but adjusted for font-scaling.

[`layout.Context`](https://gioui.org/layout#Context) has method [`Px`](https://gioui.org/layout#Context.Px) to convert from [`unit.Value`](https://gioui.org/unit#Value) to pixels

<!-- TODO: example -->

For more information on pixel-density see:

* https://material.io/design/layout/pixel-density.html.
* https://webplatform.github.io/docs/tutorials/understanding-css-units/

### Coordinate systems

You may have noticed that widget constraints and dimensions sizes are in integer units, while drawing commands such as [`PaintOp`](https://gioui.org/op/paint#PaintOp) use floating point units. That's because they refer to two distinct coordinate systems, the layout coordinate system and the drawing coordinate system. The distinction is subtle, but important.

The layout coordinate system is in integer pixels, because it's important that widgets never unintentionally overlap in the middle of a physical pixel. In fact, the decision to use integer coordinates was motivated by [conflation issues](https://github.com/flutter/flutter/issues/15035) in other UI libraries caused by allowing fractional layouts.

As a bonus, integer coordinates are perfectly deterministic across all platforms which leads to easier debugging and testing of layouts.

On the other hand, drawing commands need the generality of floating point coordinates for smooth animation and for expression inherently fractional shapes such as bézier curves.

It's possible to draw shapes that overlap at fractional pixel coordinates, but only intentionally: drawing commands directly derived from layout constraints have integer coordinates by construction.


## Custom Widget

As an example, here is how to implement a very simple button.

Let's start by drawing it:

<{{files/architecture/button.go}}[/START VISUAL OMIT/,/END VISUAL OMIT/]

<pre style="min-height: 100px" data-run="wasm" data-pkg="architecture" data-args="button-visual" data-size="200x100"></pre>

Then handle pointer clicks:

<{{files/architecture/button.go}}[/START FINAL OMIT/,/END FINAL OMIT/]

<pre style="min-height: 100px" data-run="wasm" data-pkg="architecture" data-args="button" data-size="200x100"></pre>


## Layout

Package [`gioui.org/layout`](https://gioui.org/layout) provides support for common layout operations such as spacing, lists and stacks.

In the layout examples we'll use this `ColorBox` widget to visualize layouts:

<{{files/architecture/colorbox.go}}[/START WIDGET OMIT/,/END WIDGET OMIT/]


### Inset

[`layout.Inset`](https://gioui.org/layout#Inset) adds space around a widget.

<{{files/architecture/layout.go}}[/START INSET OMIT/,/END INSET OMIT/]

<pre style="min-height: 100px" data-run="wasm" data-pkg="architecture" data-args="layout-inset" data-size="200x100"></pre>


### Stack

[`layout.Stack`](https://gioui.org/layout#Stack) lays out child elements on top of each other, according to the alignment direction. The child of a stack layout can be:

* [`Stacked`](https://gioui.org/layout#Stacked) - which doesn't have minimum constraints and the maximum constraints passed to Stack.Layout.
* [`Expanded`](https://gioui.org/layout#Expanded) - which uses the largest Stacked item as the minimum constraint and maximum is the maximum constraints passed to Stack.Layout.

For example, this draws green and blue rectangles on top of a red background:

<{{files/architecture/layout.go}}[/START STACK OMIT/,/END STACK OMIT/]

<pre style="min-height: 100px" data-run="wasm" data-pkg="architecture" data-args="layout-stack" data-size="200x100"></pre>


### List

[`layout.List`](https://gioui.org/layout#List) can display a potentially large list of items. Since `List` also handles scrolling it must be persisted across layouts, otherwise the scrolling position is lost.

<{{files/architecture/layout.go}}[/START LIST OMIT/,/END LIST OMIT/]

<pre style="min-height: 100px" data-run="wasm" data-pkg="architecture" data-args="layout-list" data-size="200x100"></pre>


### Flex

[`layout.Flex`](https://gioui.org/layout#List) lays out children according to their weights or rigid constraints. First the rigid elements are used to determine the remaining space and then the remaining space is divided among flexed children according to weights.

The children can be:

* [`Rigid`](https://gioui.org/layout#Rigid) - are laid out with as much space left over from other rigid children.
* [`Flexed`](https://gioui.org/layout#Flexed) - children are sized according to their weights and thespace left over from rigid children.

<{{files/architecture/layout.go}}[/START FLEX OMIT/,/END FLEX OMIT/]

<pre style="min-height: 100px" data-run="wasm" data-pkg="architecture" data-args="layout-flex" data-size="200x100"></pre>


## Themes

The same abstract widget can have many visual representations, ranging from simple color changes to entirely custom graphics. To give an application a consistent appearance it is useful to have an abstraction that represents a particular "theme".

Package [`gioui.org/widget/material`](https://gioui.org/widget/material) implements a theme based on the [Material Design](https://material.io/design), and the [`Theme`](https://gioui.org/widget/material#Theme) struct encapsulates the parameters for varying colors, sizes and fonts.

To use a theme, you must first initialize in your application loop:

<{{files/architecture/main.go}}[/START THEMELOOP OMIT/,/END THEMELOOP OMIT/]

Then in your application use the provided widgets:

<{{files/architecture/theme.go}}[/START EXAMPLE OMIT/,/END EXAMPLE OMIT/]

<pre style="min-height: 100px" data-run="wasm" data-pkg="architecture" data-args="theme" data-size="200x100"></pre>

[Kitchen example](https://git.sr.ht/~eliasnaur/gio/tree/master/example/kitchen/kitchen.go) shows all the different widgets available.

## Custom Layout

Sometimes the builtin layouts are not sufficient. To create a custom layout for widgets there are special functions and structures to manipulate layout.Context. In general, layouting code performs the following steps for each sub-widget:

* Use `StackOp.Push`.
* Set `layout.Context.Constraints`.
* Set `op.TransformOp`.
* Call `widget.Layout(gtx, ...)`.
* Use dimensions returned by widget.
* Use `StackOp.Pop`.

For complicated layouts you would also need to use macros. As an example take a look at [layout.Flex](https://gioui.org/layout#Flex). Which roughly implements:

1. Record widgets in macros.
2. Calculate sizes for non-rigid widgets.
3. Draw widgets based on the calculated sizes by replaying their macros.

### Example: Split View Widget

As an example, the following layout displays two widgets side-by-side:

<{{files/architecture/split-visual.go}}[/START WIDGET OMIT/,/END WIDGET OMIT/]

The usage code would look like:

<{{files/architecture/split-visual.go}}[/START EXAMPLE OMIT/,/END EXAMPLE OMIT/]

<pre style="min-height: 100px" data-run="wasm" data-pkg="architecture" data-args="split-visual" data-size="200x100"></pre>

#### Interactivity

To make it more useful we could make the split draggable.

First let's make the ratio adjustable. We should try to make zero values useful, in this case `0` could mean that it's split in the center.

<{{files/architecture/split-ratio.go}}[/START WIDGET OMIT/,/END WIDGET OMIT/]

The usage code would look like:

<{{files/architecture/split-ratio.go}}[/START EXAMPLE OMIT/,/END EXAMPLE OMIT/]

<pre style="min-height: 100px" data-run="wasm" data-pkg="architecture" data-args="split-ratio" data-size="200x100"></pre>

Because we also need to have an area designated for moving the split, let's add a bar into the center:

<{{files/architecture/split-interactive.go}}[/START BAR OMIT/,/END BAR OMIT/]

Now we need to store our interactive state:

<{{files/architecture/split-interactive.go}}[/START INPUTSTATE OMIT/,/END INPUTSTATE OMIT/]

And then we need to handle input events:

<{{files/architecture/split-interactive.go}}[/START INPUTCODE OMIT/,/END INPUTCODE OMIT/]

Putting the whole widget together:

<{{files/architecture/split-interactive.go}}[/START WIDGET OMIT/,/END WIDGET OMIT/]

And an example:

<{{files/architecture/split-interactive.go}}[/START EXAMPLE OMIT/,/END EXAMPLE OMIT/]

<pre style="min-height: 100px" data-run="wasm" data-pkg="architecture" data-args="split-interactive" data-size="200x100"></pre>


## Common errors

### The system is drawing on top of my custom widget, or otherwise ignoring its size.

The problem: You've created a nice new widget. You lay it out, say, in a Flex Rigid. The next Rigid draws on top of it.

The explanation: Gio communicates the size of widgets dynamically via layout.Context.Dimensions (commonly "gtx.Dimensions"). High level widgets (such as Labels) "return" or pass on their dimensions in gtx.Dimensions, but lower-level operations, such as paint.PaintOp, do not set Dimensions.

The solution: Update gtx.Dimensions in your widget's Layout function before you return.

### My list.List won't scroll

The problem: You lay out a list and then it just sits there and doesn't scroll.

The explanation: A lot of widgets in Gio are context free -- you can and should declare them every time through your Layout function. Lists are not like that. They record their scroll position internally, and that needs to persist between calls to Layout.

The solution: Declare your List once outside the event handling loop and reuse it across frames.

### The system is ignoring updates to a widget

The problem: You define a field in your widget struct with the widget. You update the child widget state, either implicitly or explicitly. The child widget stubbornly refuses to reflect your updates.

This is related to the problem with Lists that won't scroll.

One possible explanation: You might be seeing a common "gotcha" in Go code, where you've defined a method on a value receiver, not a pointer receiver, so all the updates you're making to your widget are only visible inside that function, and thrown away when it returns.

The solution: `Layout` and `Update` methods on stateful widgets should have pointer receivers.


---
title: Drawing
subtitle: Displaying things on the screen
---

The [`paint`](https://gioui.org/op/paint) package provides operations for drawing graphics.

Coordinates are based on the top-left corner, although it's possible to [transform the coordinate system](https://gioui.org/op#TransformOp). This means `f32.Point{X:0, Y:0}` is the top left corner of the window. All drawing operations use pixel units, see [Units](#units) section for more information.

For example, the following code will draw a 100x100 pixel colored rectangle at the top left corner of the window:

<{{files/architecture/draw.go}}[/START DRAWING OMIT/,/END DRAWING OMIT/]

<pre style="min-height: 100px" data-run="wasm" data-pkg="architecture" data-args="draw-paint" data-size="200x100"></pre>

## Offset

Operation [`op.TransformOp`](https://gioui.org/op#TransformOp) translates the position of the operations that come after it.

For example, the following function offsets the red rectangle 100 pixels to the right:

<{{files/architecture/draw.go}}[/START TRANSFORMATION OMIT/,/END TRANSFORMATION OMIT/]

<pre style="min-height: 100px" data-run="wasm" data-pkg="architecture" data-args="draw-transformation" data-size="200x100"></pre>

## Clipping

In some cases we want the drawing to be confined to a non-rectangular shape, for example to avoid overlapping drawings. Package [`gioui.org/op/clip`](https://gioui.org/op/clip) provides exactly that.

[`clip.RRect`](https://gioui.org/op/clip#RRect) clips all subsequent drawing operations to a rectangle with rounded corners. This is useful as a basis for a button background:

<{{files/architecture/draw.go}}[/START CLIPPING OMIT/,/END CLIPPING OMIT/]

<pre style="min-height: 100px" data-run="wasm" data-pkg="architecture" data-args="draw-clip" data-size="200x100"></pre>

For more complex clipping [`clip.Path`](https://gioui.org/op/clip#Path) can express shapes built from lines and bézier curves. This example draws a triangle with a curved edge:

<{{files/architecture/draw.go}}[/START CLIP TRIANGLE OMIT/,/END CLIP TRIANGLE OMIT/]

<pre style="min-height: 100px" data-run="wasm" data-pkg="architecture" data-args="draw-clip-triangle" data-size="200x100"></pre>

## Lines

To draw lines it's possible to use [`clip.Stroke`](https://gioui.org/op/clip#Stroke)
instead of [`clip.Outline`](https://gioui.org/op/clip#Outline).
We can also use [`paint.FillShape`](https://gioui.org/op/paint#FillShape) helper
to avoid managing the clip state or use `ColorOp` or `PaintOp`.

It's possible to use the predefined shapes, such as [`clip.RRect`](https://gioui.org/op/clip#RRect):

<{{files/architecture/draw.go}}[/START STROKE RECT OMIT/,/END STROKE RECT OMIT/]

<pre style="min-height: 100px" data-run="wasm" data-pkg="architecture" data-args="draw-stroke-rect" data-size="200x100"></pre>

Or use a custom shape drawn with [`clip.Path`](https://gioui.org/op/clip#Path):

<{{files/architecture/draw.go}}[/START STROKE TRIANGLE OMIT/,/END STROKE TRIANGLE OMIT/]

<pre style="min-height: 100px" data-run="wasm" data-pkg="architecture" data-args="draw-stroke-triangle" data-size="200x100"></pre>



For dashes, stroke end caps and joins, there's a separate package [gioui.org/x/stroke](https://gioui.org/x/stroke).
However, they are not as performant as `clip.Stroke`.


## Operation Stack

Some operations affect all operations that follow them. For example, [`paint.ColorOp`](https://gioui.org/op/paint#ColorOp) sets the "brush" color that is used in subsequent [`paint.PaintOp`](https://gioui.org/op/paint#PaintOp) operations. This drawing context also includes coordinate transformation (set by [`op.TransformOp`](https://gioui.org/op#TransformOp)) and clipping (set by [`clip.Op`](https://gioui.org/op/clip#Op)).

Some operations, such as clips and transformations, allow you to temporarily apply them and later restore the previous state.

For example, the `redButtonBackground` function in the previous section has the unfortunate side-effect of clipping all later operations to the outline of the button background! Let's make a version of it that doesn't affect any callers:

<{{files/architecture/draw.go}}[/START STACK OMIT/,/END STACK OMIT/]

<pre style="min-height: 100px" data-run="wasm" data-pkg="architecture" data-args="draw-stack" data-size="200x100"></pre>

## Drawing Order {#macros}

Drawing happens from back to front. In this function the green rectangle is drawn on top of red rectangle:

<{{files/architecture/draw.go}}[/START DRAWORDER OMIT/,/END DRAWORDER OMIT/]

<pre style="min-height: 100px" data-run="wasm" data-pkg="architecture" data-args="draw-draworder" data-size="200x100"></pre>

Sometimes you may want to change this order. For example, you may want to delay drawing to apply a transform that is calculated during drawing, or you may want to perform a list of operations several times. For this purpose there is [op.MacroOp](https://gioui.org/op#MacroOp).

<!-- TODO: Use a more practical example here. -->

<{{files/architecture/draw.go}}[/START MACRO OMIT/,/END MACRO OMIT/]

<pre style="min-height: 100px" data-run="wasm" data-pkg="architecture" data-args="draw-macro" data-size="200x100"></pre>

## Animation

Gio only issues FrameEvents when the window is resized or the user interacts with the window. However, animation requires continuous redrawing until the animation is completed. For that there is [`op.InvalidateOp`](https://gioui.org/op#InvalidateOp).

The following code will animate a green "progress bar" that fills up from left to right over 10 seconds from when the program starts:

<{{files/architecture/draw.go}}[/START ANIMATION OMIT/,/END ANIMATION OMIT/]

<pre style="min-height: 100px" data-run="wasm" data-pkg="architecture" data-args="draw-animation" data-size="200x100"></pre>

## Record and replay

While `op.MacroOp` allows you to record and replay operations on a single operation list, [`op.CallOp`](https://gioui.org/op#CallOp) allows for reuse of a separate operation list. This is useful for caching operations that are expensive to re-create, or for animating the disappearance of otherwise removed widgets:

<{{files/architecture/draw.go}}[/START CACHE OMIT/,/END CACHE OMIT/]

<pre style="min-height: 100px" data-run="wasm" data-pkg="architecture" data-args="draw-cache" data-size="200x100"></pre>

## Images

[`paint.ImageOp`](https://gioui.org/op/paint#ImageOp) is used to draw images. Like [`paint.ColorOp`](https://gioui.org/op/paint#ColorOp), it sets part of the drawing context (the "brush") that's used for subsequent [`PaintOp`](https://gioui.org/op/paint#PaintOp). [`ImageOp`](https://gioui.org/op/paint#ImageOp) is used similarly to [`ColorOp`](https://gioui.org/op/paint#ColorOp).

Note that [`image.NRGBA`](https://golang.org/pkg/image#NRGBA) and [`image.Uniform`](https://golang.org/pkg/image#Uniform) images are efficient and treated specially. Other [`Image`](https://golang.org/pkg/image#Image) implementations will undergo a more expensive copy and conversion to the underlying image model.

<{{files/architecture/draw.go}}[/START IMAGE OMIT/,/END IMAGE OMIT/]

<pre style="min-height: 100px" data-run="wasm" data-pkg="architecture" data-args="draw-image" data-size="200x100"></pre>

The image must not be mutated until another [`FrameEvent`](https://gioui.org/io/system#FrameEvent) happens, because the image may be read asynchronously while the frame is being drawn.

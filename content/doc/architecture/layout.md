---
title: Layout
subtitle: Putting things where they belong
---

Package [`gioui.org/layout`](https://gioui.org/layout) provides support for common layout operations such as spacing, lists and stacks of overlapping widgets.

In the layout examples we'll use this `ColorBox` widget to visualize layouts:

<{{files/architecture/colorbox.go}}[/START WIDGET OMIT/,/END WIDGET OMIT/]


## Inset

[`layout.Inset`](https://gioui.org/layout#Inset) adds space around a widget.

<{{files/architecture/layout.go}}[/START INSET OMIT/,/END INSET OMIT/]

<pre style="min-height: 100px" data-run="wasm" data-pkg="architecture" data-args="layout-inset" data-size="200x100"></pre>


## Stack

[`layout.Stack`](https://gioui.org/layout#Stack) lays out overlapping child elements according to the alignment direction. The child of a stack layout can be:

* [`Stacked`](https://gioui.org/layout#Stacked) - which doesn't have minimum constraints and the maximum constraints passed to Stack.Layout.
* [`Expanded`](https://gioui.org/layout#Expanded) - which uses the largest Stacked item as the minimum constraint and maximum is the maximum constraints passed to Stack.Layout.

For example, this draws green and blue rectangles on top of a red background:

<{{files/architecture/layout.go}}[/START STACK OMIT/,/END STACK OMIT/]

<pre style="min-height: 100px" data-run="wasm" data-pkg="architecture" data-args="layout-stack" data-size="400x100"></pre>

### Background

Because layouting a background for a widget is very frequent there is a more performant implementation for that scenario, which roughly corresponds to:

``` go
layout.Stack{Alignment: layout.C}.Layout(gtx,
	layout.Expanded(background),
	layout.Stacked(widget)
)
```

<{{files/architecture/layout.go}}[/START BACKGROUND OMIT/,/END BACKGROUND OMIT/]

<pre style="min-height: 100px" data-run="wasm" data-pkg="architecture" data-args="layout-background" data-size="400x100"></pre>


## List

[`layout.List`](https://gioui.org/layout#List) can display a potentially large list of items. Since `List` also handles scrolling it must be persisted across layouts, otherwise the scrolling position is lost. List handles large numbers of items by only laying out the visible elements. Each frame, the provided closure is invoked only for indicies visible at the current scroll position (and possibly a small number of items above and below the scroll position).

<{{files/architecture/layout.go}}[/START LIST OMIT/,/END LIST OMIT/]

<pre style="min-height: 100px" data-run="wasm" data-pkg="architecture" data-args="layout-list" data-size="400x100"></pre>


## Flex

[`layout.Flex`](https://gioui.org/layout#List) lays out children according to their weights or rigid constraints. First the rigid elements are used to determine the remaining space and then the remaining space is divided among flexed children according to weights.

The children can be:

* [`Rigid`](https://gioui.org/layout#Rigid) - are laid out with as much space left over from other rigid children.
* [`Flexed`](https://gioui.org/layout#Flexed) - children are sized according to their weights and the space left over from rigid children.

<{{files/architecture/layout.go}}[/START FLEX OMIT/,/END FLEX OMIT/]

<pre style="min-height: 100px" data-run="wasm" data-pkg="architecture" data-args="layout-flex" data-size="400x100"></pre>

## Spacer

[`layout.Spacer`](https://gioui.org/layout#Spacer) can be used together with `layout.List` or `layout.Flex` to add empty space between items.

<{{files/architecture/layout.go}}[/START SPACER OMIT/,/END SPACER OMIT/]

<pre style="min-height: 100px" data-run="wasm" data-pkg="architecture" data-args="layout-spacer" data-size="400x100"></pre>

## Custom

Sometimes the builtin layouts are not sufficient. To create a custom layout for widgets there are special functions and structures to manipulate layout.Context. In general, layout code performs the following steps for each sub-widget:

* Use `op.Save`.
* Set `layout.Context.Constraints`.
* Set `op.TransformOp`.
* Call `widget.Layout(gtx, ...)`.
* Use dimensions returned by widget.
* Use `StateOp.Load`.

For complicated layouts you would also need to use macros. As an example take a look at [layout.Flex](https://gioui.org/layout#Flex). Which roughly implements:

1. Record widgets in macros.
2. Calculate sizes for non-rigid widgets.
3. Draw widgets based on the calculated sizes by replaying their macros.

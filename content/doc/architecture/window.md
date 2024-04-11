---
title: Window
subtitle: Talking with the OS
---

Since a GUI library needs to talk to some sort of display system to display information:

<{{files/architecture/main.go}}[/START DRAWLOOP OMIT/,/END DRAWLOOP OMIT/]

[`app.Window.Run`](http://gioui.org/app#Window.Run) chooses the appropriate "driver" depending on the environment and build context. It might choose Wayland, Win32, or Cocoa among several others.

An `app.Window` allows accessing events from the display with [`window.Event()`](https://gioui.org/app#Window.Event). There are other lifecycle events in the `gioui.org/app` package such as [`app.DestroyEvent`](https://gioui.org/app#DestroyEvent) and [`app.FrameEvent`](https://gioui.org/app#FrameEvent).


## Operations

All UI libraries need a way for the program to specify what to display and how to handle events. Gio programs use operations, serialized into one or more [`op.Ops`](https://gioui.org/op#Ops) operation lists. Operation lists are in turn passed to the window driver through the [`FrameEvent.Frame`](https://gioui.org/app#FrameEvent.Frame) function.

By convention, each operation kind is represented by a Go type with an `Add` method that records the operation into the `Ops` argument. Like any Go struct literal, zero-valued fields can be useful to represent optional values.

For example, recording an operation that sets the current color to red:

<{{files/architecture/draw.go}}[/START OPERATIONS OMIT/,/END OPERATIONS OMIT/]

You might be thinking that it would be more usual to have an `ops.Add(ColorOp{Color: red})` method instead of using `op.ColorOp{Color: red}.Add(ops)`. It's like this so that the `Add` method doesn't have to take an interface-typed argument, which would often require an allocation to call. This is a key aspect of Gio's "zero allocation" design.

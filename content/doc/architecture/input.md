---
title: Input
subtitle: Reacting to a mouse and keyboard
---

Input is delivered to the widgets via a [`system.FrameEvent`](https://gioui.org/io/system#FrameEvent) through the [`Queue`](https://gioui.org/io/system#FrameEvent.Queue) field.

Some of the most common events in `FrameEvent.Queue` are:

* [`key.Event`](https://gioui.org/io/key#Event), [`key.FocusEvent`](https://gioui.org/io/key#FocusEvent) - for keyboard input.
* [`key.EditEvent`](https://gioui.org/io/key#EditEvent) - for text editing.
* [`pointer.Event`](https://gioui.org/io/pointer#Event) - for mouse and touch input.

The program can respond to these events however it likes - for example, by updating its local data structures or running a user-triggered action. The [`FrameEvent`](https://gioui.org/io/system#FrameEvent) is special - when the program receives a [`FrameEvent`](https://gioui.org/io/system#FrameEvent), it is responsible for updating the display by calling the [`e.Frame`](https://gioui.org/io/system#FrameEvent.Frame) function with an operation list representing the new state. These operations are generated immediately in response to the [`FrameEvent`](https://gioui.org/io/system#FrameEvent) which is the main reason that Gio is known as an "immediate mode" GUI.

Event-processors, such as [`Click`](https://gioui.org/gesture#Click) and [`Scroll`](https://gioui.org/gesture#Scroll) from [package `gioui.org/gesture`](https://gioui.org/gesture) detect higher-level actions from individual click events.

To distribute input among multiple different widgets, Gio needs to know about event handlers and their configuration. However, since the Gio framework is stateless, there's no direct way for the program to specify that.

Instead, some operations associate input event types (for example, keyboard presses) with arbitrary [tags](https://gioui.org/io/event#Tag) (interface{} values) chosen by the program. A program creates these operations when it's processing the [`FrameEvent`](https://gioui.org/io/system#FrameEvent) -- input operations are operations like any other. In return, an [event.Queue](https://gioui.org/io/event#Queue) supplies the events that arrived since the last frame, separated by tag.

You can think about the tag as a unique key for a given input area. The Gio event router will associate input events on in that area with the tag provided for that area. Then you can get those events the next frame by supplying the same tag to `event.Queue`. Often widgets will encapsulate this event logic by supplying a pointer to their persistent state as the tag for their input area.

The following example demonstrates pointer input handling:

<{{files/architecture/button.go}}[/START LOWLEVEL OMIT/,/END LOWLEVEL OMIT/]

<pre style="min-height: 100px" data-run="wasm" data-pkg="architecture" data-args="button-low" data-size="200x100"></pre>

It's convenient to use a Go pointer value for the input tag, as it's cheap to convert a pointer to an interface{} and it's easy to make the value specific to a local data structure, which avoids the risk of tag conflict.

For more details take a look at [`gioui.org/io/pointer`](https://gioui.org/io/pointer) (pointer/mouse events) and [`gioui.org/io/key`](https://gioui.org/io/key) (keyboard events).


## External input

A single frame consists of getting input, registering for input and drawing the new state:

<{{files/architecture/main.go}}[/START DRAWQUEUELOOP OMIT/,/END DRAWQUEUELOOP OMIT/]

Let's make the button change it's position every second. We can use a select to wait for events from the window and the external source at the same time. We'll use a [`Ticker`](https://golang.org/pkg/time#Ticker) as an example external change. Once we have modified the state we need to notify the window to retrigger rendering with [`w.Invalidate()`](https://gioui.org/app#Window.Invalidate).

<{{files/architecture/external.go}}[/START LOOP OMIT/,/END LOOP OMIT/]

<pre style="min-height: 100px" data-run="wasm" data-pkg="architecture" data-args="external-changes" data-size="200x100"></pre>

Writing a program using these concepts could get really verbose, which is why Gio provides standard widgets for common look and behaviour. Most programs end up using widgets primarily and few low-level operations.

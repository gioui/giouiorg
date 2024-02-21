---
title: Common Errors
subtitle: We've all been there
---

## My list.List won't scroll

The problem: You lay out a list and then it just sits there and doesn't scroll.

The explanation: A lot of widgets in Gio are context free -- you can and should declare them every time through your Layout function. Lists are not like that. They record their scroll position internally, and that needs to persist between calls to Layout.

The solution: Declare your List once outside the event handling loop and reuse it across frames.

## The system is ignoring updates to a widget

The problem: You define a field in your widget struct that contains one of the provided types in `gioui.org/widget`. You update the child widget state, either implicitly or explicitly. The child widget stubbornly refuses to reflect your updates.

This is related to the problem with Lists that won't scroll.

One possible explanation: You might be seeing a common "gotcha" in Go code, where you've defined a method on a value receiver, not a pointer receiver, so all the updates you're making to your widget are only visible inside that function, and thrown away when it returns.

The solution: `Layout` and `Update` methods on stateful widgets should have pointer receivers.

## Custom widget ignores size

The problem: You've created a nice new widget. You lay it out, say, in a Flex Rigid. The next Rigid draws on top of it.

The explanation: Gio communicates the size of widgets dynamically via returned `layout.Dimensions`. High level widgets (such as Labels) return or pass on their dimensions, but lower-level operations, such as paint.PaintOp, do not automatically provide their dimensions.

The solution: calculate the proper dimensions of the content you drew with your custom operations, and return that in your `layout.Dimension`.

## Dependencies don't compile any more

The problem: You've updated your Gio version with `go get -u gioui.org@latest` and things don't compile.

The explanation: In Go `go get -u` (the `-u` part) is unfortunately an [unsafe operation for pre v1.0 releases](https://github.com/golang/go/issues/64864), which includes Gio and some dependencies such as typesetting. `-u` ends up downloading the latest minor version for all dependencies, where unstable dependencies may have breaking changes.

The solution: update Gio dependencies only with `go get gioui.org@latest`. If you have ended up in a very messy situation you can first try reverting `go.mod` to your older commit.

If the suggestions above don't help, then you can try deleting all the lines from `go.mod`, except `module ...` and `go ...` lines, and running `go mod tidy`. This will end up downloading the latest direct dependencies.
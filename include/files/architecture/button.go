// SPDX-License-Identifier: Unlicense OR MIT

package main

import (
	"image"
	"image/color"

	"gioui.org/io/event"
	"gioui.org/io/pointer"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
)

// START LOWLEVEL OMIT
var tag = new(bool) // We could use &pressed for this instead.
var pressed = false

func doButton(ops *op.Ops, q event.Queue) {
	// Make sure we don’t pollute the graphics context.
	defer op.Save(ops).Load()

	// Process events that arrived between the last frame and this one.
	for _, ev := range q.Events(tag) {
		if x, ok := ev.(pointer.Event); ok {
			switch x.Type {
			case pointer.Press:
				pressed = true
			case pointer.Release:
				pressed = false
			}
		}
	}

	// Confine the area of interest to a 100x100 rectangle.
	pointer.Rect(image.Rect(0, 0, 100, 100)).Add(ops)
	// Declare the tag.
	pointer.InputOp{
		Tag:   tag,
		Types: pointer.Press | pointer.Release,
	}.Add(ops)

	clip.Rect{Max: image.Pt(100, 100)}.Add(ops)
	var c color.NRGBA
	if pressed {
		c = color.NRGBA{R: 0xFF, A: 0xFF}
	} else {
		c = color.NRGBA{G: 0xFF, A: 0xFF}
	}
	paint.ColorOp{Color: c}.Add(ops)
	paint.PaintOp{}.Add(ops)
}

// END LOWLEVEL OMIT

var buttonVisual ButtonVisual

func handleButtonVisual(gtx layout.Context) layout.Dimensions {
	return buttonVisual.Layout(gtx)
}

// START VISUAL OMIT
type ButtonVisual struct {
	pressed bool
}

func (b *ButtonVisual) Layout(gtx layout.Context) layout.Dimensions {
	col := color.NRGBA{R: 0x80, A: 0xFF}
	if b.pressed {
		col = color.NRGBA{G: 0x80, A: 0xFF}
	}
	return drawSquare(gtx.Ops, col)
}

func drawSquare(ops *op.Ops, color color.NRGBA) layout.Dimensions {
	defer op.Save(ops).Load()
	clip.Rect{Max: image.Pt(100, 100)}.Add(ops)
	paint.ColorOp{Color: color}.Add(ops)
	paint.PaintOp{}.Add(ops)
	return layout.Dimensions{Size: image.Pt(100, 100)}
}

// END VISUAL OMIT

var button Button

func handleButton(gtx layout.Context) layout.Dimensions {
	return button.Layout(gtx)
}

// START FINAL OMIT
type Button struct {
	pressed bool
}

func (b *Button) Layout(gtx layout.Context) layout.Dimensions {
	// Avoid affecting the input tree with pointer events.
	defer op.Save(gtx.Ops).Load()

	// here we loop through all the events associated with this button.
	for _, e := range gtx.Events(b) {
		if e, ok := e.(pointer.Event); ok {
			switch e.Type {
			case pointer.Press:
				b.pressed = true
			case pointer.Release:
				b.pressed = false
			}
		}
	}

	// Confine the area for pointer events.
	pointer.Rect(image.Rect(0, 0, 100, 100)).Add(gtx.Ops)
	pointer.InputOp{
		Tag:   b,
		Types: pointer.Press | pointer.Release,
	}.Add(gtx.Ops)

	// Draw the button.
	col := color.NRGBA{R: 0x80, A: 0xFF}
	if b.pressed {
		col = color.NRGBA{G: 0x80, A: 0xFF}
	}
	return drawSquare(gtx.Ops, col)
}

// END FINAL OMIT

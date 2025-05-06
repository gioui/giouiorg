// SPDX-License-Identifier: Unlicense OR MIT

package main

import (
	"image"
	"image/color"

	"gioui.org/io/event"
	"gioui.org/io/input"
	"gioui.org/io/pointer"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
)

// START LOWLEVEL OMIT
var tag = new(bool) // We could use &pressed for this instead.
var pressed = false

func doButton(ops *op.Ops, q input.Source) {
	// Confine the area of interest to a 100x100 rectangle.
	defer clip.Rect{Max: image.Pt(100, 100)}.Push(ops).Pop()

	// Declare `tag` as being one of the targets.
	event.Op(ops, tag)

	// Process events that arrived between the last frame and this one.
	for {
		ev, ok := q.Event(pointer.Filter{
			Target: tag,
			Kinds:  pointer.Press | pointer.Release,
		})
		if !ok {
			break
		}

		if x, ok := ev.(pointer.Event); ok {
			switch x.Kind {
			case pointer.Press:
				pressed = true
			case pointer.Release:
				pressed = false
			}
		}
	}

	// Draw the button.
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

// START INPUTTREE OMIT
var (
	// Declare a number of variables to use both as state
	// and input tags.
	root, child1, child2 bool
)

// displayForTag adds a pointer.InputOp interested
// in press and release events to the given op.Ops using
// the given tag. It also paints a color based on the current
// value of the tag to the current clip.
func displayForTag(ops *op.Ops, tag *bool, rect clip.Rect) {
	event.Op(ops, tag)

	// Choose a color based on whether the tag is being pressed.
	c := color.NRGBA{B: 0xFF, A: 0xFF}
	if *tag {
		c = color.NRGBA{R: 0xFF, A: 0xFF}
	}
	// Paint the current clipping area with a translucent color.
	translucent := c
	translucent.A = 0x44
	paint.ColorOp{Color: translucent}.Add(ops)
	paint.PaintOp{}.Add(ops)

	// Reduce our clipping area to the outline of the rectangle, then
	// paint that outline. This should make it easier to see overlapping
	// rectangles.
	defer clip.Stroke{
		Path:  rect.Path(),
		Width: 5,
	}.Op().Push(ops).Pop()
	paint.ColorOp{Color: c}.Add(ops)
	paint.PaintOp{}.Add(ops)
}

func doPointerTree(ops *op.Ops, q input.Source) {
	// Process events that arrived between the last frame and this one for every tag.
	for _, tag := range []*bool{&root, &child1, &child2} {
		for {
			ev, ok := q.Event(pointer.Filter{
				Target: tag,
				Kinds:  pointer.Press | pointer.Release,
			})
			if !ok {
				break
			}

			x, ok := ev.(pointer.Event)
			if !ok {
				continue
			}

			switch x.Kind {
			case pointer.Press:
				*tag = true
			case pointer.Release:
				*tag = false
			}
		}
	}

	// Confine the rootArea of interest to a 200x200 rectangle.
	rootRect := clip.Rect(image.Rect(0, 0, 200, 200))
	rootArea := rootRect.Push(ops)
	displayForTag(ops, &root, rootRect)

	// Any clip areas we add before Pop-ing the root area
	// are considered its children.
	child1Rect := clip.Rect(image.Rect(25, 25, 175, 100))
	child1Area := child1Rect.Push(ops)
	displayForTag(ops, &child1, child1Rect)
	child1Area.Pop()

	child2Rect := clip.Rect(image.Rect(100, 25, 175, 175))
	child2Area := child2Rect.Push(ops)
	displayForTag(ops, &child2, child2Rect)
	child2Area.Pop()

	rootArea.Pop()
	// Now anything we add is _not_ a child of the rootArea.
}

// END INPUTTREE OMIT

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
	defer clip.Rect{Max: image.Pt(100, 100)}.Push(ops).Pop()
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
	// Confine the area for pointer events.
	area := clip.Rect(image.Rect(0, 0, 100, 100)).Push(gtx.Ops)

	event.Op(gtx.Ops, b)

	// here we loop through all the events associated with this button.
	for ev := range gtx.Events(pointer.Filter{
		Target: b,
		Kinds:  pointer.Press | pointer.Release,
	}) {
		e, ok := ev.(pointer.Event)
		if !ok {
			continue
		}
		switch e.Kind {
		case pointer.Press:
			b.pressed = true
		case pointer.Release:
			b.pressed = false
		}
	}

	area.Pop()

	// Draw the button.
	col := color.NRGBA{R: 0x80, A: 0xFF}
	if b.pressed {
		col = color.NRGBA{G: 0x80, A: 0xFF}
	}
	return drawSquare(gtx.Ops, col)
}

// END FINAL OMIT

package main

import (
	"image"

	"gioui.org/f32"
	"gioui.org/io/pointer"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
	"gioui.org/widget/material"
)

// START EXAMPLE OMIT
var split Split

func exampleSplit(gtx layout.Context, th *material.Theme) layout.Dimensions {
	return split.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		return FillWithLabel(gtx, th, "Left", red)
	}, func(gtx layout.Context) layout.Dimensions {
		return FillWithLabel(gtx, th, "Right", blue)
	})
}

// END EXAMPLE OMIT

// START WIDGET OMIT
// START INPUTSTATE OMIT
type Split struct {
	// Ratio keeps the current layout.
	// 0 is center, -1 completely to the left, 1 completely to the right.
	Ratio float32
	// Bar is the width for resizing the layout
	Bar unit.Value

	drag   bool
	dragID pointer.ID
	dragX  float32
}

// END INPUTSTATE OMIT

var defaultBarWidth = unit.Dp(10)

func (s *Split) Layout(gtx layout.Context, left, right layout.Widget) layout.Dimensions {
	// START BAR OMIT
	bar := gtx.Px(s.Bar)
	if bar <= 1 {
		bar = gtx.Px(defaultBarWidth)
	}

	proportion := (s.Ratio + 1) / 2
	leftsize := int(proportion*float32(gtx.Constraints.Max.X) - float32(bar))

	rightoffset := leftsize + bar
	rightsize := gtx.Constraints.Max.X - rightoffset
	// END BAR OMIT

	{ // handle input
		// Avoid affecting the input tree with pointer events.
		var stack op.StackOp
		stack.Push(gtx.Ops)
		defer stack.Pop()

		// START INPUTCODE OMIT
		for _, ev := range gtx.Events(s) {
			e, ok := ev.(pointer.Event)
			if !ok {
				continue
			}

			switch e.Type {
			case pointer.Press:
				if s.drag {
					break
				}

				s.drag = true
				s.dragID = e.PointerID
				s.dragX = e.Position.X

			case pointer.Move:
				if !s.drag || s.dragID != e.PointerID {
					break
				}

				deltaX := e.Position.X - s.dragX
				s.dragX = e.Position.X

				deltaRatio := deltaX * 2 / float32(gtx.Constraints.Max.X)
				s.Ratio += deltaRatio

			case pointer.Release:
				fallthrough
			case pointer.Cancel:
				if !s.drag || s.dragID != e.PointerID {
					break
				}
				s.drag = false
			}
		}

		// register for input
		barRect := image.Rect(leftsize, 0, rightoffset, gtx.Constraints.Max.X)
		pointer.Rect(barRect).Add(gtx.Ops)
		pointer.InputOp{Tag: s, Grab: s.drag}.Add(gtx.Ops)
		// END INPUTCODE OMIT
	}

	{
		var stack op.StackOp
		stack.Push(gtx.Ops)

		gtx := gtx
		gtx.Constraints.Min = image.Pt(leftsize, leftsize)
		left(gtx)

		stack.Pop()
	}

	{
		var stack op.StackOp
		stack.Push(gtx.Ops)

		op.TransformOp{}.Offset(f32.Pt(float32(rightoffset), 0)).Add(gtx.Ops)
		gtx := gtx
		gtx.Constraints.Min = image.Pt(rightsize, rightsize)
		right(gtx)

		stack.Pop()
	}

	return layout.Dimensions{Size: gtx.Constraints.Max}
}

// END WIDGET OMIT

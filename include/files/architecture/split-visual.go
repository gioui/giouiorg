package main

import (
	"image"
	"image/color"

	"gioui.org/f32"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/widget/material"
)

// START EXAMPLE OMIT
func exampleSplitVisual(gtx layout.Context, th *material.Theme) layout.Dimensions {
	return SplitVisual{}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		return FillWithLabel(gtx, th, "Left", red)
	}, func(gtx layout.Context) layout.Dimensions {
		return FillWithLabel(gtx, th, "Right", blue)
	})
}

func FillWithLabel(gtx layout.Context, th *material.Theme, text string, backgroundColor color.RGBA) layout.Dimensions {
	ColorBox(gtx, gtx.Constraints.Max, backgroundColor)
	return layout.Center.Layout(gtx, material.H3(th, text).Layout)
}

// END EXAMPLE OMIT

// START WIDGET OMIT
type SplitVisual struct{}

func (s SplitVisual) Layout(gtx layout.Context, left, right layout.Widget) layout.Dimensions {
	leftsize := gtx.Constraints.Min.X / 2
	rightsize := gtx.Constraints.Min.X - leftsize

	{
		stack := op.Push(gtx.Ops)

		gtx := gtx
		gtx.Constraints = layout.Exact(image.Pt(leftsize, gtx.Constraints.Max.Y))
		left(gtx)

		stack.Pop()
	}

	{
		stack := op.Push(gtx.Ops)

		gtx := gtx
		gtx.Constraints = layout.Exact(image.Pt(rightsize, gtx.Constraints.Max.Y))
		op.TransformOp{}.Offset(f32.Pt(float32(leftsize), 0)).Add(gtx.Ops)
		right(gtx)

		stack.Pop()
	}

	return layout.Dimensions{Size: gtx.Constraints.Max}
}

// END WIDGET OMIT

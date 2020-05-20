package main

import (
	"image/color"

	"gioui.org/f32"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/widget/material"
)

// START EXAMPLE OMIT
func exampleSplitVisual(gtx *layout.Context, th *material.Theme) {
	SplitVisual{}.Layout(gtx, func() {
		FillWithLabel(gtx, th, "Left", red)
	}, func() {
		FillWithLabel(gtx, th, "Right", blue)
	})
}

func FillWithLabel(gtx *layout.Context, th *material.Theme, text string, backgroundColor color.RGBA) {
	ColorBox(gtx, gtx.Constraints.Max, backgroundColor)
	layout.Center.Layout(gtx, func() {
		material.H3(th, text).Layout(gtx)
	})
}

// END EXAMPLE OMIT

// START WIDGET OMIT
type SplitVisual struct{}

func (s SplitVisual) Layout(gtx *layout.Context, left, right layout.Widget) {
	savedConstraints := gtx.Constraints
	defer func() {
		gtx.Constraints = savedConstraints
		gtx.Dimensions.Size = gtx.Constraints.Max
	}()

	leftsize := gtx.Constraints.Min.X / 2
	rightsize := gtx.Constraints.Min.X - leftsize

	{
		var stack op.StackOp
		stack.Push(gtx.Ops)

		gtx.Constraints.Min.X = leftsize
		gtx.Constraints.Max.X = leftsize
		left()

		stack.Pop()
	}

	{
		var stack op.StackOp
		stack.Push(gtx.Ops)

		gtx.Constraints.Min.X = rightsize
		gtx.Constraints.Max.X = rightsize

		op.TransformOp{}.Offset(f32.Pt(float32(leftsize), 0)).Add(gtx.Ops)
		right()

		stack.Pop()
	}
}

// END WIDGET OMIT

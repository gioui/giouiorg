package main

import (
	"gioui.org/f32"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/widget/material"
)

// START EXAMPLE OMIT
func exampleSplitRatio(gtx *layout.Context, th *material.Theme) {
	SplitRatio{Ratio: -0.3}.Layout(gtx, func() {
		FillWithLabel(gtx, th, "Left", red)
	}, func() {
		FillWithLabel(gtx, th, "Right", blue)
	})
}

// END EXAMPLE OMIT

// START WIDGET OMIT
type SplitRatio struct {
	// Ratio keeps the current layout.
	// 0 is center, -1 completely to the left, 1 completely to the right.
	Ratio float32
}

func (s SplitRatio) Layout(gtx *layout.Context, left, right layout.Widget) {
	savedConstraints := gtx.Constraints
	defer func() {
		gtx.Constraints = savedConstraints
		gtx.Dimensions.Size = gtx.Constraints.Max
	}()

	proportion := (s.Ratio + 1) / 2
	leftsize := int(proportion * float32(gtx.Constraints.Max.X))

	rightoffset := leftsize
	rightsize := gtx.Constraints.Max.X - rightoffset

	{
		var stack op.StackOp
		stack.Push(gtx.Ops)

		gtx.Constraints.Min.X, gtx.Constraints.Max.X = leftsize, leftsize
		left()

		stack.Pop()
	}

	{
		var stack op.StackOp
		stack.Push(gtx.Ops)

		op.TransformOp{}.Offset(f32.Pt(float32(rightoffset), 0)).Add(gtx.Ops)
		gtx.Constraints.Min.X, gtx.Constraints.Max.X = rightsize, rightsize
		right()

		stack.Pop()
	}
}

// END WIDGET OMIT

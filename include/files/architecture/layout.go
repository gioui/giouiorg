// SPDX-License-Identifier: Unlicense OR MIT

package main

import (
	"image"
	"image/color"

	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
)

// START INSET OMIT
func inset(gtx layout.Context) layout.Dimensions {
	// Draw rectangles inside of each other, with 30dp padding.
	return layout.UniformInset(unit.Dp(30)).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		return ColorBox(gtx, gtx.Constraints.Max, red)
	})
}

// END INSET OMIT

// START STACK OMIT
func stacked(gtx layout.Context) layout.Dimensions {
	return layout.Stack{}.Layout(gtx,
		// Force widget to the same size as the second.
		layout.Expanded(func(gtx layout.Context) layout.Dimensions {
			// This will have a minimum constraint of 100x100.
			return ColorBox(gtx, gtx.Constraints.Min, red)
		}),
		layout.Stacked(func(gtx layout.Context) layout.Dimensions {
			return ColorBox(gtx, image.Pt(100, 30), green)
		}),
		layout.Stacked(func(gtx layout.Context) layout.Dimensions {
			return ColorBox(gtx, image.Pt(30, 100), blue)
		}),
	)
}

// END STACK OMIT

// START BACKGROUND OMIT
func layoutBackground(gtx layout.Context) layout.Dimensions {
	return layout.Background{}.Layout(gtx,
		func(gtx layout.Context) layout.Dimensions {
			defer clip.Rect{Max: gtx.Constraints.Min}.Push(gtx.Ops).Pop()
			paint.Fill(gtx.Ops, background)
			return layout.Dimensions{Size: gtx.Constraints.Min}
		}, func(gtx layout.Context) layout.Dimensions {
			return ColorBox(gtx, image.Pt(30, 100), blue)
		})
}

// END BACKGROUND OMIT

// START LIST OMIT
var list = layout.List{}

func listing(gtx layout.Context) layout.Dimensions {
	return list.Layout(gtx, 100, func(gtx layout.Context, i int) layout.Dimensions {
		col := color.NRGBA{R: byte(i * 20), G: 0x20, B: 0x20, A: 0xFF}
		return ColorBox(gtx, image.Pt(20, 100), col)
	})
}

// END LIST OMIT

// START FLEX OMIT
func flexed(gtx layout.Context) layout.Dimensions {
	return layout.Flex{}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return ColorBox(gtx, image.Pt(100, 100), red)
		}),
		layout.Flexed(0.5, func(gtx layout.Context) layout.Dimensions {
			return ColorBox(gtx, gtx.Constraints.Min, blue)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return ColorBox(gtx, image.Pt(100, 100), red)
		}),
		layout.Flexed(0.5, func(gtx layout.Context) layout.Dimensions {
			return ColorBox(gtx, gtx.Constraints.Min, green)
		}),
	)
}

// END FLEX OMIT

// START SPACER OMIT
func spacer(gtx layout.Context) layout.Dimensions {
	return layout.Flex{}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return ColorBox(gtx, image.Pt(100, 100), red)
		}),
		layout.Rigid(layout.Spacer{Width: 20}.Layout),
		layout.Flexed(0.5, func(gtx layout.Context) layout.Dimensions {
			return ColorBox(gtx, gtx.Constraints.Min, blue)
		}),
		layout.Rigid(layout.Spacer{Width: 20}.Layout),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return ColorBox(gtx, image.Pt(100, 100), red)
		}),
		layout.Rigid(layout.Spacer{Width: 20}.Layout),
		layout.Flexed(0.5, func(gtx layout.Context) layout.Dimensions {
			return ColorBox(gtx, gtx.Constraints.Min, green)
		}),
	)
}

// END SPACER OMIT

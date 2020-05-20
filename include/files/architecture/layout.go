// SPDX-License-Identifier: Unlicense OR MIT

package main

import (
	"image"
	"image/color"

	"gioui.org/layout"
	"gioui.org/unit"
)

// START INSET OMIT
func inset(gtx *layout.Context) {
	// Draw rectangles inside of each other, with 20dp padding.
	layout.UniformInset(unit.Dp(30)).Layout(gtx, func() {
		ColorBox(gtx, gtx.Constraints.Max, red)
	})
}

// END INSET OMIT

// START STACK OMIT
func stacked(gtx *layout.Context) {
	layout.Stack{}.Layout(gtx,
		// Force widget to the same size as the second.
		layout.Expanded(func() {
			// This will have a minimum constraint of 100x100.
			ColorBox(gtx, gtx.Constraints.Min, red)
		}),
		layout.Stacked(func() {
			ColorBox(gtx, image.Pt(100, 30), green)
		}),
		layout.Stacked(func() {
			ColorBox(gtx, image.Pt(30, 100), blue)
		}),
	)
}

// END STACK OMIT

// START LIST OMIT
var list = layout.List{}

func listing(gtx *layout.Context) {
	list.Layout(gtx, 100, func(i int) {
		col := color.RGBA{R: byte(i * 20), G: 0x20, B: 0x20, A: 0xFF}
		ColorBox(gtx, image.Pt(20, 100), col)
	})
}

// END LIST OMIT

// START FLEX OMIT
func flexed(gtx *layout.Context) {
	layout.Flex{}.Layout(gtx,
		layout.Rigid(func() {
			ColorBox(gtx, image.Pt(100, 100), red)
		}),
		layout.Flexed(0.5, func() {
			ColorBox(gtx, gtx.Constraints.Min, blue)
		}),
		layout.Rigid(func() {
			ColorBox(gtx, image.Pt(100, 100), red)
		}),
		layout.Flexed(0.5, func() {
			ColorBox(gtx, gtx.Constraints.Min, green)
		}),
	)
}

// END FLEX OMIT

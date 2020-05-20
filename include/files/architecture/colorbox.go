// SPDX-License-Identifier: Unlicense OR MIT

package main

import (
	"image"
	"image/color"

	"gioui.org/f32"
	"gioui.org/layout"
	"gioui.org/op/paint"
)

// START WIDGET OMIT

// Test colors.
var (
	background = color.RGBA{R: 0xC0, G: 0xC0, B: 0xC0, A: 0xFF}
	red        = color.RGBA{R: 0xC0, G: 0x40, B: 0x40, A: 0xFF}
	green      = color.RGBA{R: 0x40, G: 0xC0, B: 0x40, A: 0xFF}
	blue       = color.RGBA{R: 0x40, G: 0x40, B: 0xC0, A: 0xFF}
)

// ColorBox creates a widget with the specified dimensions and color.
func ColorBox(gtx *layout.Context, size image.Point, color color.RGBA) {
	gtx.Dimensions.Size = size
	bounds := f32.Rect(0, 0, float32(size.X), float32(size.Y))
	paint.ColorOp{Color: color}.Add(gtx.Ops)
	paint.PaintOp{Rect: bounds}.Add(gtx.Ops)
}

// END WIDGET OMIT

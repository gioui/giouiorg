// SPDX-License-Identifier: Unlicense OR MIT

package main

import (
	"image"
	"image/color"

	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/text"
	"gioui.org/website/include/files/architecture/internal/f32color"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

var (
	colorA    = hslColor(0.3, 1.0, 0.5)
	colorB    = hslColor(0.9, 1.0, 0.5)
	blendBias widget.Float
)

func init() { blendBias.Value = 0.5 }

func colorMixing(gtx layout.Context, th *material.Theme) layout.Dimensions {
	return layout.Flex{
		Axis: layout.Vertical,
	}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return layout.Flex{Axis: layout.Horizontal, Alignment: layout.Middle}.Layout(gtx,
				layout.Flexed(1, centeredText(gtx, th, "A")),
				layout.Flexed(8, func(gtx layout.Context) layout.Dimensions {
					return colorA.Layout(gtx, th)
				}),
				flexSpacing(),
				layout.Flexed(4, filledColor{
					Theme: th,
					Color: colorA.RGBA().SRGB(),
				}.Layout),
				flexSpacing(),
			)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return layout.Flex{Axis: layout.Horizontal, Alignment: layout.Middle}.Layout(gtx,
				layout.Flexed(1, centeredText(gtx, th, "B")),
				layout.Flexed(8, func(gtx layout.Context) layout.Dimensions {
					return colorB.Layout(gtx, th)
				}),
				flexSpacing(),
				layout.Flexed(4, filledColor{
					Theme: th,
					Color: colorB.RGBA().SRGB(),
				}.Layout),
				flexSpacing(),
			)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return layout.Flex{Axis: layout.Horizontal, Alignment: layout.Middle}.Layout(gtx,
				layout.Flexed(1, centeredText(gtx, th, "bias")),
				layout.Flexed(8, material.Slider(th, &blendBias).Layout),
				flexSpace(4),
				flexSpacing(),
			)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return layout.Flex{Axis: layout.Horizontal, Alignment: layout.Middle}.Layout(gtx,
				layout.Flexed(9, centeredText(gtx, th, "linear color blending")),
				flexSpacing(),
				layout.Flexed(4, filledColor{
					Theme: th,
					Color: f32color.BlendRGBA(blendBias.Value, colorA.RGBA(), colorB.RGBA()).SRGB(),
				}.Layout),
				flexSpacing(),
			)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return layout.Flex{Axis: layout.Horizontal, Alignment: layout.Middle}.Layout(gtx,
				layout.Flexed(9, centeredText(gtx, th, "sRGB color blending")),
				flexSpacing(),
				layout.Flexed(4, filledColor{
					Theme: th,
					Color: f32color.BlendSRGBA(blendBias.Value, colorA.SRGBA(), colorB.SRGBA()),
				}.Layout),
				flexSpacing(),
			)
		}),
	)
}

func flexSpacing() layout.FlexChild { return flexSpace(0.3) }

func flexSpace(weight float32) layout.FlexChild {
	return layout.Flexed(weight, func(gtx layout.Context) layout.Dimensions {
		return layout.Dimensions{
			Size: image.Point{X: gtx.Constraints.Max.X},
		}
	})
}

type hslColorState struct {
	hue   widget.Float
	sat   widget.Float
	light widget.Float
}

func hslColor(h, s, l float32) hslColorState {
	var c hslColorState
	c.hue.Value = h
	c.sat.Value = s
	c.light.Value = l
	return c
}

func (c hslColorState) RGBA() f32color.RGBA {
	return f32color.HSLA(c.hue.Value, c.sat.Value, c.light.Value, 1)
}

func (c hslColorState) SRGBA() color.NRGBA { return c.RGBA().SRGB() }

func (c *hslColorState) Layout(gtx layout.Context, th *material.Theme) layout.Dimensions {
	return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
		layout.Rigid(layoutNamedComponentSlider(gtx, th, "hue", &c.hue)),
		layout.Rigid(layoutNamedComponentSlider(gtx, th, "sat", &c.sat)),
		layout.Rigid(layoutNamedComponentSlider(gtx, th, "lig", &c.light)),
	)
}

func layoutNamedComponentSlider(gtx layout.Context, th *material.Theme, name string, value *widget.Float) layout.Widget {
	return func(gtx layout.Context) layout.Dimensions {
		return layout.Flex{}.Layout(gtx,
			layout.Flexed(2, centeredText(gtx, th, name)),
			layout.Flexed(8, material.Slider(th, value).Layout),
		)
	}
}

func centeredText(gtx layout.Context, th *material.Theme, s string) layout.Widget {
	return func(gtx layout.Context) layout.Dimensions {
		body := material.Body1(th, s)
		body.Alignment = text.Middle
		return body.Layout(gtx)
	}
}

type filledColor struct {
	Theme *material.Theme
	Color color.NRGBA
}

func (c filledColor) Layout(gtx layout.Context) layout.Dimensions {
	size := gtx.Constraints.Constrain(image.Pt(gtx.Constraints.Max.X, gtx.Metric.Sp(c.Theme.TextSize*4)))

	paint.FillShape(gtx.Ops, c.Color, clip.Rect{
		Max: size,
	}.Op())
	return layout.Dimensions{
		Size: size,
	}
}

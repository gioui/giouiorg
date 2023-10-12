// SPDX-License-Identifier: Unlicense OR MIT

package f32color

import "image/color"

// BlendRGBA blends to colors in linear color space.
func BlendRGBA(p float32, a, b RGBA) RGBA {
	return RGBA{
		R: lerpClamp(p, a.R, b.R),
		G: lerpClamp(p, a.G, b.G),
		B: lerpClamp(p, a.B, b.B),
		A: lerpClamp(p, a.A, b.A),
	}
}

// BlendSRGBA blends to colors in sRGB color space.
func BlendSRGBA(p float32, a, b color.NRGBA) color.NRGBA {
	return mix(a, b, sat8(p))
}

// SPDX-License-Identifier: Unlicense OR MIT

package f32color

import "math"

// HSLA converts hue-staturation-lightness to a color.NRGBA
func HSLA(h, s, l, a float32) RGBA {
	r, g, b, a := hsla(h, s, l, a)
	return RGBA{R: r, G: g, B: b, A: a}
}

func hsla(h, s, l, a float32) (r, g, b, ra float32) {
	if s == 0 {
		return l, l, l, a
	}

	h = mod32(h, 1)

	var v2 float32
	if l < 0.5 {
		v2 = l * (1 + s)
	} else {
		v2 = (l + s) - s*l
	}

	v1 := 2*l - v2
	r = hue(v1, v2, h+1.0/3.0)
	g = hue(v1, v2, h)
	b = hue(v1, v2, h-1.0/3.0)
	ra = a

	return
}

func hue(v1, v2, h float32) float32 {
	if h < 0 {
		h += 1
	}
	if h > 1 {
		h -= 1
	}
	if 6*h < 1 {
		return v1 + (v2-v1)*6*h
	} else if 2*h < 1 {
		return v2
	} else if 3*h < 2 {
		return v1 + (v2-v1)*(2.0/3.0-h)*6
	}

	return v1
}

// sat8 converts 0..1 float to 0..255 uint8
func sat8(v float32) uint8 {
	v *= 255.0
	if v >= 255 {
		return 255
	} else if v <= 0 {
		return 0
	}
	return uint8(v)
}

func lerpClamp(p, min, max float32) float32 {
	if p < 0 {
		return min
	} else if p > 1 {
		return max
	}
	return min + (max-min)*p
}

func mod32(x, y float32) float32 { return float32(math.Mod(float64(x), float64(y))) }

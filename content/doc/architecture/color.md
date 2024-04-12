---
title: Color
subtitle: Understanding color and blending
---

Color handling is something that we don’t usually don’t think about. However, a
framework can make many tradeoffs while handling color.

The short explanation is that Gio uses sRGB colors for input but uses linear
color space for blending. This results in the color blending being correct
without manually converting usual color values to linear color space.

If the short explanation wasn't sufficient, then there's a longer one below.

_Note: the following will oversimplify things to make them more understandable.
For all the gritty details, read the linked articles._

## Color primer

Most programs represent colors with red, green, and blue lightness values. The
simplest approach would be to represent the exact lightness value with the
value you use in your RGB color. However, eyes are more sensitive to darker
colors than lighter colors. With just 8 bits available per color channel, a
linear mapping of lightness wastes bits to represent lighter values that people
cannot differentiate.

One approach is [gamma correction] (https://en.wikipedia.org/wiki/Gamma_correction)
that encodes lightness values with a function that stretches the darker color
range at the cost of compressing the lighter color range.

Usually the gamma transformations look like:

```
// transforming linear color to gamma compressed color
gamma_color  := math.Pow(linear_color, gamma)
// transforming gamma compressed color to linear color
linear_color := math.Pow(gamma_color, 1/gamma)

// where
linear_color = [0..1]
gamma_color  = [0..1]
gamma        = usually 2.2 or 2.4
```

One of the problems with this function is that the [rate of color change is near infinite](https://en.wikipedia.org/wiki/SRGB#Transfer_function_\(%22gamma%22\)).
To avoid this boundary condition there is a lightness value transformation
called [sRGB color space](https://en.wikipedia.org/wiki/SRGB). sRGB conversion looks like:

```
// transforming linear color to sRGB color
if linear_color <= 0.0031308 {
	srgb_color = 12.92 * linear_color
} else { // linear_color > 0.0031308
	srgb_color = 1.055 * math.Pow(linear_color, 1/2.4) - 0.055
}

// transforming sRGB color to linear color
if srgb_color <= 0.04045 {
	linear_color = srgb_color / 12.92
} else { // srgb_color > 0.04045
	linear_color = math.Pow((srgb_color + 0.055) / 1.055, 2.4)
}
```

The details of the sRGB vs gamma corrected colors aren't that important for the
discussion, so we'll keep using the gamma transformation, because it's shorter
to write than the sRGB conversions.

## Problems with sRGB

One of the problems that sRGB and gamma-corrected colors have is that when you
directly compute the sum of them, you don't get the correct color mixing.

Let's take an example of mixing `linear_color_alpha` and `linear_color_beta`:

```
// mix colors using linear color space
linear_color = 0.5*linear_color_alpha + 0.5*linear_color_beta

// mixing colors using sRGB color space
linear_color = math.Pow(
	0.5 * math.Pow(linear_color_alpha, gamma) +
	0.5 * math.Pow(linear_color_beta, gamma),
	1/gamma)
```

<pre style="min-height: 100px" data-run="wasm" data-pkg="architecture" data-args="color-mix" data-size="730x420"></pre>

When you experiment with this example, you should notice that blending in sRGB
results often in a darker or grayer color, which ends up causing muddied
colors in blending.

The blending issues have been discussed in more detail in:

* [How software gets color wrong](https://bottosson.github.io/posts/colorwrong/)
* [Linear Gamma vs Higher Gamma](https://ninedegreesbelow.com/photography/linear-gamma-blur-normal-blend.html)
* [Gamma error in picture scaling](http://www.ericbrasseur.org/gamma.html)

## Choice for frameworks

Overall, frameworks need to choose a color space to work with.
Historically, the most common choice was sRGB because of the darker color
benefit. Similarly, as an accident or for performance reasons, people ended
up using sRGB blending. This also leads to bugs related to [resizing images](http://www.ericbrasseur.org/gamma.html).

So, due to the historical importance of sRGB, there are a few choices for a UI framework:

1. Use sRGB for input and blending: this causes incorrect blending and muddy
   colors. However, this behavior is similar to all other programs.

2. Use linear colors for input and blending: this has correct blending. However,
   people cannot use their usual "color pickers" (because they work in sRGB) and
   must manually convert images from sRGB to linear.

3. Use sRGB colors when providing input; however, blend using linear colors:
   this is compatible with programs for color selection. Mixing colors is going to
   be different from sRGB blending.

Gio has chosen approach **3**, because it's a pragmatic choice that
has correct blending and does not have the annoyances of color conversion.

_Sidenote: of course, there are more choices, such as using higher bit-depth or
wide-gamut color spaces, but for usual UI applications, there isn't a
significant benefit from them._
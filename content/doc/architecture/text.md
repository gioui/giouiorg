---
title: Text
subtitle: Low-level text management
---

## Fonts

Gio's text shaper uses the type `[]text.FontFace` to represent the collection of available fonts.

There is one font bundled in package [`gioui.org/font/gofont`](https://gioui.org/font/gofont), you can use [`gofont.Collection()`](https://gioui.org/font/gofont#Collection) to get a `[]text.FontFace` containing all of the variants of the Go fonts.

For loading other fonts there is [`gioui.org/font/opentype`](https://gioui.org/font/opentype). After parsing the font(s) using [`opentype.Parse`](https://gioui.org/font/opentype#Parse), you can append them to a `[]text.FontFace`.

<!-- TODO: code example. -->

## Shapes

For converting strings to clip shapes there is the [`gioui.org/text`](https://gioui.org/text) package.

It contains [`text.Cache`](https://gioui.org/text#Cache) that implements cached string to shape conversion, with appropriate fallbacks. Simply provide your fonts (`[]text.FontFace`) to `text.NewCache`.

In most cases you can use [`widget.Label`](https://gioui.org/widget#Label) which handles wrapping and layout constraints. Or when you are using material design [`material.LabelStyle`](https://gioui.org/widget/material#LabelStyle).

---
title: Text
subtitle: Low-level text management
---

## Fonts

[`gioui.org/font`](https://gioui.org/font) contains the central registry for fonts. This helps to reduce passing fonts around throughout the program.

There is one font bundled in package [`gioui.org/font/gofont`](https://gioui.org/font/gofont), you can use [`gofont.Register()`](https://gioui.org/font/gofont#Register) to register it to the global registry.

For loading other fonts there is [`gioui.org/font/opentype`](https://gioui.org/font/opentype). After parsing the font(s) using [`opentype.Parse`](https://gioui.org/font/opentype#Parse) or [`opentype.ParseCollection`](https://gioui.org/font/opentype#ParseCollection) they can be registered with [`font.Register`](https://gioui.org/font#Register).

<!-- TODO: code example. -->

## Shapes

For converting strings to clip shapes there is the [`gioui.org/text`](https://gioui.org/text) package.

It contains [`text.Cache`](https://gioui.org/text#Cache) that implements cached string to shape conversion, with appropriate fallbacks.

In most cases you can use [`widget.Label`](https://gioui.org/widget#Label) which handles wrapping and layout constraints. Or when you are using material design [`material.LabelStyle`](https://gioui.org/widget/material#LabelStyle).
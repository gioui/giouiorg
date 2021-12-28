---
title: Theme
subtitle: Making things look the same
---

The same abstract widget can have many visual representations, ranging from simple color changes to entirely custom graphics. To give an application a consistent appearance it is useful to have an abstraction that represents a particular "theme".

Package [`gioui.org/widget/material`](https://gioui.org/widget/material) implements a theme based on the [Material Design](https://material.io/design), and the [`Theme`](https://gioui.org/widget/material#Theme) struct encapsulates the parameters for varying colors, sizes and fonts.

To use a theme, you must first initialize it in your application loop:

<{{files/architecture/main.go}}[/START THEMELOOP OMIT/,/END THEMELOOP OMIT/]

Then in your application use the provided widgets:

<{{files/architecture/theme.go}}[/START EXAMPLE OMIT/,/END EXAMPLE OMIT/]

<pre style="min-height: 100px" data-run="wasm" data-pkg="architecture" data-args="theme" data-size="200x100"></pre>

[Kitchen example](https://git.sr.ht/~eliasnaur/gio-example/tree/main/example/kitchen/kitchen.go) shows all the different widgets available.

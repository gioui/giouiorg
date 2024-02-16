---
title: Split Widget
subtitle: Tailoring things to your own needs
---

Sometimes there's a need for writing a custom widget or layout.

To implement rendering of children, we can use:

<{{files/architecture/split-visual.go}}[/START WIDGET OMIT/,/END WIDGET OMIT/]

Then we can use the widget like:

<{{files/architecture/split-visual.go}}[/START EXAMPLE OMIT/,/END EXAMPLE OMIT/]

<pre style="min-height: 200px" data-run="wasm" data-pkg="architecture" data-args="split-visual" data-size="400x200"></pre>

## Ratio

Let's make the ratio adjustable. We should try to make zero values useful, in this case `0` could mean that it's split in the center.

<{{files/architecture/split-ratio.go}}[/START WIDGET OMIT/,/END WIDGET OMIT/]

The usage code would look like:

<{{files/architecture/split-ratio.go}}[/START EXAMPLE OMIT/,/END EXAMPLE OMIT/]

<pre style="min-height: 200px" data-run="wasm" data-pkg="architecture" data-args="split-ratio" data-size="400x200"></pre>

## Interactive

To make it more useful we could make the split draggable.

Because we also need to have an area designated for moving the split, let's add a bar into the center:

<{{files/architecture/split-interactive.go}}[/START BAR OMIT/,/END BAR OMIT/]

Now we need to store our interactive state:

<{{files/architecture/split-interactive.go}}[/START INPUTSTATE OMIT/,/END INPUTSTATE OMIT/]

And then we need to handle input events:

<{{files/architecture/split-interactive.go}}[/START INPUTCODE OMIT/,/END INPUTCODE OMIT/]

## Result

Putting the whole widget together:

<{{files/architecture/split-interactive.go}}[/START WIDGET OMIT/,/END WIDGET OMIT/]

And an example:

<{{files/architecture/split-interactive.go}}[/START EXAMPLE OMIT/,/END EXAMPLE OMIT/]

<pre style="min-height: 200px" data-run="wasm" data-pkg="architecture" data-args="split-interactive" data-size="400x200"></pre>
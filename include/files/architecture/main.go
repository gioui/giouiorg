// SPDX-License-Identifier: Unlicense OR MIT

package main

import (
	"fmt"
	"os"
	"strings"

	"gioui.org/app"
	"gioui.org/font/gofont"
	"gioui.org/io/input"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/text"
	"gioui.org/widget/material"
)

func main() {
	type command struct {
		name string
		run  func(title string) error
	}

	var commands []*command
	commands = []*command{
		// drawing section
		{name: "draw-operations", run: drawLoop(addColorOperation)},
		{name: "draw-paint", run: drawLoop(drawRedRect)},
		{name: "draw-transformation", run: drawLoop(drawRedRect10PixelsRight)},
		{name: "draw-clip", run: drawLoop(redButtonBackground)},
		{name: "draw-clip-triangle", run: drawLoop(redTriangle)},
		{name: "draw-stroke-rect", run: drawLoop(strokeRect)},
		{name: "draw-stroke-triangle", run: drawLoop(strokeTriangle)},
		{name: "draw-stack", run: drawLoop(redButtonBackgroundStack)},
		{name: "draw-draworder", run: drawLoop(drawOverlappingRectangles)},
		{name: "draw-macro", run: drawLoop(drawFiveRectangles)},
		{name: "draw-animation", run: drawQueueLoop(drawProgressBarInternal)},
		{name: "draw-cache", run: drawLoop(drawWithCache)},
		{name: "draw-image", run: drawLoop(drawImageInternal)},

		{name: "button-low", run: drawQueueLoop(doButton)},
		{name: "input-tree", run: drawQueueLoop(doPointerTree)},
		{name: "external-changes", run: externalChanges},
		{name: "button-visual", run: contextLoop(handleButtonVisual)},
		{name: "button", run: contextLoop(handleButton)},

		{name: "layout-inset", run: contextLoop(inset)},
		{name: "layout-stack", run: contextLoop(stacked)},
		{name: "layout-background", run: contextLoop(layoutBackground)},
		{name: "layout-list", run: contextLoop(listing)},
		{name: "layout-flex", run: contextLoop(flexed)},
		{name: "layout-spacer", run: contextLoop(spacer)},

		{name: "theme", run: themeLoop(themedApplication)},

		{name: "color-mix", run: themeLoop(colorMixing)},

		{name: "split-visual", run: themeLoop(exampleSplitVisual)},
		{name: "split-ratio", run: themeLoop(exampleSplitRatio)},
		{name: "split-interactive", run: themeLoop(exampleSplit)},

		{name: "all", run: func(title string) error {
			for _, cmd := range commands {
				if cmd.name == "all" {
					continue
				}

				cmd := cmd
				go func() {
					if err := cmd.run(cmd.name); err != nil {
						fmt.Fprintln(os.Stderr, err)
					}
				}()
			}
			return nil
		}},
	}

	var cmdname string
	if len(os.Args) >= 2 {
		cmdname = os.Args[1]
	}

	var cmd *command
	for _, c := range commands {
		if strings.EqualFold(c.name, cmdname) {
			cmd = c
		}
	}

	if len(os.Args) <= 1 || cmd == nil {
		if cmdname != "" {
			fmt.Fprintf(os.Stderr, "unknown command %q\n", cmdname)
		}
		fmt.Fprint(os.Stderr, "basics [command]:\n")
		for _, cmd := range commands {
			fmt.Fprintf(os.Stderr, "\t%q\n", cmd.name)
		}
		os.Exit(1)
	}

	go func() {
		err := cmd.run(cmd.name)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}()
	app.Main()
}

func drawLoop(draw func(*op.Ops)) func(title string) error {
	return func(title string) error {
		// START DRAWLOOP OMIT
		var window app.Window
		window.Option(app.Title(title))

		var ops op.Ops
		for {
			switch e := window.Event().(type) {
			case app.DestroyEvent:
				// The window was closed.
				return e.Err
			case app.FrameEvent:
				// A request to draw the window state.

				// Reset the operations back to zero.
				ops.Reset()
				// Draw the state into ops.
				draw(&ops)
				// Events the display.
				e.Frame(&ops)
			}
		}
		// END DRAWLOOP OMIT
	}
}

func drawQueueLoop(draw func(*op.Ops, input.Source)) func(title string) error {
	return func(title string) error {
		// START DRAWQUEUELOOP OMIT
		var window app.Window
		window.Option(app.Title(title))

		var ops op.Ops
		for {
			switch e := window.Event().(type) {
			case app.DestroyEvent:
				// The window was closed.
				return e.Err
			case app.FrameEvent:
				// A request to draw the window state.

				// Reset the operations back to zero.
				ops.Reset()
				// Draw the state into ops based on events in e.Queue.
				draw(&ops, e.Source)
				// Events the display.
				e.Frame(&ops)
			}
		}
		// END DRAWQUEUELOOP OMIT
	}
}

func contextLoop(draw func(layout.Context) layout.Dimensions) func(title string) error {
	return func(title string) error {
		// START CONTEXTLOOP OMIT
		var window app.Window
		window.Option(app.Title(title))

		var ops op.Ops
		for {
			switch e := window.Event().(type) {
			case app.DestroyEvent:
				// The window was closed.
				return e.Err
			case app.FrameEvent:
				// Reset the layout.Context for a new frame.
				gtx := app.NewContext(&ops, e)

				// Draw the state into ops based on events in e.Queue.
				draw(gtx)

				// Events the display.
				e.Frame(gtx.Ops)
			}
		}
		// END CONTEXTLOOP OMIT
	}
}

func themeLoop(draw func(layout.Context, *material.Theme) layout.Dimensions) func(title string) error {
	return func(title string) error {
		// START THEMELOOP OMIT
		th := material.NewTheme()
		th.Shaper = text.NewShaper(text.WithCollection(gofont.Collection()))

		var window app.Window
		window.Option(app.Title(title))

		var ops op.Ops
		for {
			switch e := window.Event().(type) {
			case app.DestroyEvent:
				// The window was closed.
				return e.Err
			case app.FrameEvent:
				// Reset the layout.Context for a new frame.
				gtx := app.NewContext(&ops, e)

				// Draw the state into ops based on events in e.Queue.
				draw(gtx, th)

				// Events the display.
				e.Frame(gtx.Ops)
			}
		}
		// END THEMELOOP OMIT
	}
}

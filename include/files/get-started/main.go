package main

import (
	"image/color"
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/op"
	"gioui.org/text"
	"gioui.org/widget/material"
)

// START MAIN OMIT
func main() {
	go func() {
		w := app.NewWindow()
		err := run(w)
		if err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	app.Main()
}

// END MAIN OMIT

// START CREATE THEME OMIT
func run(w *app.Window) error {
	th := material.NewTheme()
	// END CREATE THEME OMIT
	var ops op.Ops
	// START PROCESS EVENTS OMIT
	for {
		switch e := w.NextEvent().(type) {
		case app.DestroyEvent:
			return e.Err
		case app.FrameEvent:
			// END PROCESS EVENTS OMIT
			// START DRAW TEXT OMIT
			// This graphics context is used for managing the rendering state.
			gtx := app.NewContext(&ops, e)

			// Define an large label with an appropriate text:
			title := material.H1(th, "Hello, Gio")

			// Change the color of the label.
			maroon := color.NRGBA{R: 127, G: 0, B: 0, A: 255}
			title.Color = maroon

			// Change the position of the label.
			title.Alignment = text.Middle

			// Draw the label to the graphics context.
			title.Layout(gtx)

			// Pass the drawing operations to the GPU.
			e.Frame(gtx.Ops)
			// END DRAW TEXT OMIT
		}
	}
}

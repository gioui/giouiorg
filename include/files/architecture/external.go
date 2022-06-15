// SPDX-License-Identifier: Unlicense OR MIT

package main

import (
	"image"
	"time"

	"gioui.org/app"
	"gioui.org/io/system"
	"gioui.org/op"
)

func externalChanges() error {
	// START LOOP OMIT
	window := app.NewWindow()

	changes := time.NewTicker(time.Second)
	defer changes.Stop()

	buttonOffset := 0

	ops := new(op.Ops)
	for {
		select {
		case e := <-window.Events():
			switch e := e.(type) {
			case system.DestroyEvent:
				return e.Err
			case system.FrameEvent:
				ops.Reset()

				// Offset the button based on state.
				op.Offset(image.Pt(buttonOffset, 0)).Add(ops)

				// Handle button input and draw.
				doButton(ops, e.Queue)

				// Update display.
				e.Frame(ops)
			}

		case t := <-changes.C:
			buttonOffset = int((t.Second() % 3) * 100)
			window.Invalidate()
		}
	}
	// END LOOP OMIT
}

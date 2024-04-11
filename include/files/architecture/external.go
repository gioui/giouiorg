// SPDX-License-Identifier: Unlicense OR MIT

package main

import (
	"image"
	"sync"
	"time"

	"gioui.org/app"
	"gioui.org/op"
)

func externalChanges(title string) error {
	// START LOOP OMIT
	var window app.Window
	window.Option(app.Title(title))

	var button struct {
		lock   sync.Mutex
		offset int
	}

	updateOffset := func(v int) {
		button.lock.Lock()
		defer button.lock.Unlock()
		button.offset = v
	}
	readOffset := func() int {
		button.lock.Lock()
		defer button.lock.Unlock()
		return button.offset
	}

	go func() {
		changes := time.NewTicker(time.Second)
		defer changes.Stop()
		for t := range changes.C {
			updateOffset(int((t.Second() % 3) * 100))
			window.Invalidate()
		}
	}()

	ops := new(op.Ops)
	for {
		switch e := window.Event().(type) {
		case app.DestroyEvent:
			return e.Err
		case app.FrameEvent:
			ops.Reset()

			// Offset the button based on state.
			op.Offset(image.Pt(readOffset(), 0)).Add(ops)

			// Handle button input and draw.
			doButton(ops, e.Source)

			// Update display.
			e.Frame(ops)
		}
	}
	// END LOOP OMIT
}

// SPDX-License-Identifier: Unlicense OR MIT

package main

import (
	"image"
	"sync"
	"time"

	"gioui.org/app"
	"gioui.org/io/system"
	"gioui.org/op"
)

func externalChanges() error {
	// START LOOP OMIT
	window := app.NewWindow()

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
		switch e := window.NextEvent().(type) {
		case system.DestroyEvent:
			return e.Err
		case system.FrameEvent:
			ops.Reset()

			// Offset the button based on state.
			op.Offset(image.Pt(readOffset(), 0)).Add(ops)

			// Handle button input and draw.
			doButton(ops, e.Queue)

			// Update display.
			e.Frame(ops)
		}
	}
	// END LOOP OMIT
}

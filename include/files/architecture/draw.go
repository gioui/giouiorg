// SPDX-License-Identifier: Unlicense OR MIT

package main

import (
	"bytes"
	"image"
	"image/color"
	"image/gif"
	"time"

	"gioui.org/f32"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"golang.org/x/image/draw"
)

// START OPERATIONS OMIT
func addColorOperation(ops *op.Ops) {
	red := color.RGBA{R: 0xFF, A: 0xFF}
	paint.ColorOp{Color: red}.Add(ops)
}

// END OPERATIONS OMIT

// START DRAWING OMIT
func drawRedRect(ops *op.Ops) {
	paint.ColorOp{Color: color.RGBA{R: 0x80, A: 0xFF}}.Add(ops)
	paint.PaintOp{Rect: f32.Rect(0, 0, 100, 100)}.Add(ops)
}

// END DRAWING OMIT

// START TRANSFORMATION OMIT
func drawRedRect10PixelsRight(ops *op.Ops) {
	op.TransformOp{}.Offset(f32.Pt(100, 0)).Add(ops)
	drawRedRect(ops)
}

// END TRANSFORMATION OMIT

// START CLIPPING OMIT
func redButtonBackground(ops *op.Ops) {
	const r = 10 // roundness
	bounds := f32.Rect(0, 0, 100, 100)
	clip.Rect{Rect: bounds, SE: r, SW: r, NW: r, NE: r}.Add(ops)
	drawRedRect(ops)
}

// END CLIPPING OMIT

// START CLIP TRIANGLE OMIT
func redTriangle(ops *op.Ops) {
	var path clip.Path
	path.Begin(ops)
	path.Move(f32.Pt(50, 0))
	path.Quad(f32.Pt(0, 90), f32.Pt(50, 100))
	path.Line(f32.Pt(-100, 0))
	path.Line(f32.Pt(50, -100))
	path.End().Add(ops)
	drawRedRect(ops)
}

// END CLIP TRIANGLE OMIT

// START STACK OMIT
func redButtonBackgroundStack(ops *op.Ops) {
	defer op.Push(ops).Pop()

	const r = 1 // roundness
	bounds := f32.Rect(0, 0, 100, 100)
	clip.Rect{Rect: bounds, SE: r, SW: r, NW: r, NE: r}.Op(ops).Add(ops)
	drawRedRect(ops)
}

// END STACK OMIT

// START DRAWORDER OMIT
func drawOverlappingRectangles(ops *op.Ops) {
	// Draw a red rectangle.
	paint.ColorOp{Color: color.RGBA{R: 0x80, A: 0xFF}}.Add(ops)
	paint.PaintOp{Rect: f32.Rect(0, 0, 100, 50)}.Add(ops)

	// Draw a green rectangle.
	paint.ColorOp{Color: color.RGBA{G: 0x80, A: 0xFF}}.Add(ops)
	paint.PaintOp{Rect: f32.Rect(0, 0, 50, 100)}.Add(ops)
}

// END DRAWORDER OMIT

// START MACRO OMIT
func drawFiveRectangles(ops *op.Ops) {
	// Record drawRedRect operations into the macro.
	macro := op.Record(ops)
	drawRedRect(ops)
	c := macro.Stop()

	// “Play back” the macro 5 times, each time
	// translated vertically 20px and horizontally 110 pixels.
	for i := 0; i < 5; i++ {
		c.Add(ops)
		op.TransformOp{}.Offset(f32.Pt(110, 20)).Add(ops)
	}
}

// END MACRO OMIT

func drawProgressBarInternal(ops *op.Ops) {
	drawProgressBar(ops, time.Now())
}

// START ANIMATION OMIT
var startTime = time.Now()
var duration = 10 * time.Second

func drawProgressBar(ops *op.Ops, now time.Time) {
	// Calculate how much of the progress bar to draw,
	// based on the current time.
	elapsed := now.Sub(startTime)
	progress := elapsed.Seconds() / duration.Seconds()
	if progress < 1 {
		// The progress bar hasn’t yet finished animating.
		op.InvalidateOp{}.Add(ops)
	} else {
		progress = 1
	}

	paint.ColorOp{Color: color.RGBA{G: 0x80, A: 0xFF}}.Add(ops)
	width := 200 * float32(progress)
	paint.PaintOp{Rect: f32.Rect(0, 0, width, 20)}.Add(ops)
}

// END ANIMATION OMIT

// START CACHE OMIT
func drawWithCache(ops *op.Ops) {
	// Save the operations in an independent ops value (the cache).
	cache := new(op.Ops)
	macro := op.Record(cache)
	paint.ColorOp{Color: color.RGBA{G: 0x80, A: 0xFF}}.Add(cache)
	paint.PaintOp{Rect: f32.Rect(0, 0, 100, 100)}.Add(cache)
	call := macro.Stop()

	// Draw the operations from the cache.
	call.Add(ops)
}

// END CACHE OMIT

var exampleImage image.Image

func createExampleImage() (image.Image, error) {
	gif, err := gif.Decode(bytes.NewReader(gifData[:]))
	if err != nil {
		return nil, err
	}
	scaled := image.NewRGBA(image.Rect(0, 0, 100, 100))
	draw.NearestNeighbor.Scale(scaled, scaled.Bounds(), gif, gif.Bounds(), draw.Over, nil)

	return scaled, nil
}

func drawImageInternal(ops *op.Ops) {
	if exampleImage == nil {
		var err error
		exampleImage, err = createExampleImage()
		if err != nil {
			exampleImage = image.NewUniform(color.RGBA{R: 0xFF, A: 0xFF})
		}
	}
	drawImage(ops, exampleImage)
}

// START IMAGE OMIT
func drawImage(ops *op.Ops, img image.Image) {
	imageOp := paint.NewImageOp(img)
	imageOp.Add(ops)
	paint.PaintOp{Rect: f32.Rect(0, 0, 100, 100)}.Add(ops)
}

// END IMAGE OMIT

var gifData = [...]byte{
	0x47, 0x49, 0x46, 0x38, 0x37, 0x61, 0x19, 0x00, 0x19, 0x00, 0xa2, 0x00,
	0x00, 0x00, 0x00, 0x00, 0xff, 0xff, 0xff, 0x57, 0x72, 0x82, 0x39, 0x46,
	0x54, 0x96, 0xd6, 0xff, 0x00, 0x00, 0x00, 0x6c, 0x55, 0x19, 0x00, 0x00,
	0x00, 0x21, 0xf9, 0x04, 0x09, 0x0a, 0x00, 0x00, 0x00, 0x2c, 0x00, 0x00,
	0x00, 0x00, 0x19, 0x00, 0x19, 0x00, 0x00, 0x03, 0x81, 0x18, 0xba, 0xdc,
	0xfe, 0x50, 0x08, 0xf8, 0xc6, 0x90, 0x4b, 0x10, 0x32, 0x29, 0x1b, 0x9b,
	0x24, 0x86, 0xde, 0x07, 0x6e, 0x28, 0xd7, 0x95, 0x67, 0x9a, 0xae, 0x54,
	0xeb, 0xa2, 0x98, 0xd7, 0x2e, 0x68, 0xc6, 0x6e, 0x0d, 0xd1, 0xc0, 0x8e,
	0x93, 0xa2, 0x50, 0x08, 0xf8, 0x02, 0xc4, 0x00, 0xb0, 0x21, 0x44, 0x16,
	0x8f, 0xc9, 0xe5, 0x27, 0xa7, 0x20, 0x24, 0x15, 0x83, 0xdd, 0xac, 0xba,
	0xc9, 0xda, 0x66, 0x33, 0x6f, 0x6c, 0x63, 0x00, 0x77, 0x0d, 0x5f, 0x82,
	0xa1, 0x0c, 0x1e, 0xac, 0xc7, 0x64, 0xb6, 0x4b, 0x0c, 0x07, 0x13, 0x8b,
	0x4e, 0x88, 0x2c, 0x75, 0xc7, 0x5f, 0x2b, 0x76, 0x45, 0x78, 0x0b, 0x7d,
	0x4c, 0x16, 0x87, 0x03, 0x7f, 0x79, 0x85, 0x0d, 0x6f, 0x0a, 0x06, 0x49,
	0x91, 0x82, 0x43, 0x0e, 0x8e, 0x01, 0x6b, 0x85, 0x92, 0x8a, 0x14, 0x99,
	0x82, 0x77, 0x25, 0xa0, 0x0b, 0x09, 0x00, 0x3b,
}

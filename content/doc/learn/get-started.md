---
title: Get Started
subtitle: Hello, Gio!
---

This example does a really quick introduction on getting something up and
running. It does not explain all the details, those will be covered in
another tutorial.

Ensure that you have followed [installation instructions](/doc/install).
If everything is setup correctly, then running:

```
$ go run gioui.org/example/hello@latest
```

Should display a pretty "Hello, Gio!" message.

## Creating a new package

_If you are unfamiliar with Go, then more help can be found at [go.dev/learn](https://go.dev/learn/)._

First step in creating a Go program requires setting up the module.

We'll use `gio.test` as our module name, however, it's recommended to use a
repository name when you want to upload it. The module name can be later changed.

```
$ go mod init gio.test
```

## Creating the program

Let's create `main.go` with the following code:

```
package main

import (
	"image/color"
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/font/gofont"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/text"
	"gioui.org/widget/material"
)

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

func run(w *app.Window) error {
	th := material.NewTheme(gofont.Collection())
	var ops op.Ops
	for {
		e := <-w.Events()
		switch e := e.(type) {
		case system.DestroyEvent:
			return e.Err
		case system.FrameEvent:
			gtx := layout.NewContext(&ops, e)

			title := material.H1(th, "Hello, Gio")
			maroon := color.NRGBA{R: 127, G: 0, B: 0, A: 255}
			title.Color = maroon
			title.Alignment = text.Middle
			title.Layout(gtx)

			e.Frame(gtx.Ops)
		}
	}
}
```

Let's then update all the dependencies with:

```
$ go mod tidy
```

Once that succeeds, the program should start up with:

```
$ go run .
```

Now to explain what's happening.

## Creating the window

Every program requires a window, the `main` starts up the application loop that
talks to the operating system and starts the window logic in a separate
goroutine.

```
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
```

## Creating a theme

Applications need to define their fonts and different color settings.
Themes contain all the necessary information.

```
func run(w *app.Window) error {
	th := material.NewTheme(gofont.Collection())
	...
```

## Listening for events

The communication with the operating system (i.e. keyboard, mouse, GPU) happens
through events. Gio uses the following approach to listen for events:

```
	for {
		e := <-w.Events()
		switch e := e.(type) {
		case system.DestroyEvent:
			return e.Err
		case system.FrameEvent:
			...
		}
	}
```

`system.DestroyEvent` means the user pressed the close button.

`system.FrameEvent` means the program should handle input and render a new
frame.

## Drawing the text

To draw the text it needs to go through several stages:

```
	var ops op.Ops
	for {
		...
		case system.FrameEvent:
			// This graphics context is used for managing the rendering state.
			gtx := layout.NewContext(&ops, e)

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
		}
	}
```

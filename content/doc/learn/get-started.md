---
title: Get Started
subtitle: Hello, Gio!
---

This example does a really quick introduction on getting something up and
running. It does not explain all the details, those will be covered in
another tutorial.

Ensure that you have followed [installation instructions](/doc/install).
If everything is setup correctly, then running:

``` sh
go run gioui.org/example/hello@latest
```

Should display a pretty "Hello, Gio!" message.

## Creating a new package

_If you are unfamiliar with Go, then more help can be found at [go.dev/learn](https://go.dev/learn/)._

First step in creating a Go program requires setting up the module.

We'll use `gio.test` as our module name, however, it's recommended to use a
repository name when you want to upload it. The module name can be later changed.

``` sh
go mod init gio.test
```

## Creating the program

Let's create `main.go` with the following code:

<{{files/get-started/main.go}}

Let's then update all the dependencies with:

``` sh
go mod tidy
```

Once that succeeds, the program should start up with:

``` sh
go run .
```

Now to explain what's happening.

## Creating the window

Every program requires a window, the `main` starts up the application loop that
talks to the operating system and starts the window logic in a separate
goroutine.

<{{files/get-started/main.go}}[/START MAIN OMIT/,/END MAIN OMIT/]

## Creating a theme

Applications need to define their fonts and different color settings.
Themes contain all the necessary information.

<{{files/get-started/main.go}}[/START CREATE THEME OMIT/,/END CREATE THEME OMIT/]

## Listening for events

The communication with the operating system (i.e. keyboard, mouse, GPU) happens
through events. Gio uses the following approach to process events:

<{{files/get-started/main.go}}[/START PROCESS EVENTS OMIT/,/END PROCESS EVENTS OMIT/]

* `app.DestroyEvent` means the user pressed the close button.
* `app.FrameEvent` means the program should handle input and render a new
frame.

## Drawing the text

To draw the text it needs to go through several stages:

<{{files/get-started/main.go}}[/START DRAW TEXT OMIT/,/END DRAW TEXT OMIT/]

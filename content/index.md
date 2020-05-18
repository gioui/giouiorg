# Gio

Gio implements portable immediate mode GUI programs in Go. Gio programs run on
all the major platforms: iOS/tvOS, Android, Linux (Wayland/X11), macOS,
Windows, FreeBSD, OpenBSD, and experimental support for browsers (Webassembly/WebGL).
There is a [unikernel port](https://eliasnaur.com/unik) for running Gio programs in virtual machines.

Gio includes an efficient vector renderer based on the Pathfinder project (https://github.com/servo/pathfinder),
implemented on OpenGL ES and Direct3D 11.
Text and other shapes are rendered using only their outlines without baking them into texture images,
to support efficient animations, transformed drawing and pixel resolution independence.

This is a screenshot of the [Kitchen
example](https://git.sr.ht/~eliasnaur/gio/tree/master/example/kitchen/kitchen.go). If your browser
supports WebAssembly and WebGL, run the example by pressing the run
button.

{data-run="wasm" data-pkg="kitchen" data-size="800x600"}
<img src="/files/wasm/kitchen.png" alt="Kitchen screenshot" width="800"/>


## Documentation

Documentation is sparse. The
[examples](https://godoc.org/gioui.org/example) gives a feel of the
structure of typical Gio programs.

[Operations](https://godoc.org/gioui.org/op) and stateful operation
lists are the low-level primitives of Gio. The important operations
are for [drawing](https://godoc.org/gioui.org/op/paint) and
[clipping](https://godoc.org/gioui.org/op/clip), as well as
[pointer](https://godoc.org/gioui.org/io/pointer) and
[keyboard](https://godoc.org/gioui.org/io/key) input.

The [layout](https://godoc.org/gioui.org/layout) package implements
useful layouts, while the [widget](https://godoc.org/gioui.org/widget)
and [widget/material](https://godoc.org/gioui.org/widget/material)
packages implement common user interface widgets. The
[gesture](https://godoc.org/gioui.org/gesture) package detects common
gestures from lower-level input events.

Layouts, widgets and gestures are all implemented in terms of operations.

Package [app](https://godoc.org/gioui.org/app) is for creating
windows and apply operations to them. Only the app package and its
sub-packages have native dependencies, making Gio [highly
portable](https://godoc.org/gioui.org/example/glfw).

[![GoDoc](https://godoc.org/gioui.org?status.svg)](https://godoc.org/gioui.org)

## Installation

Gio is designed to work with very few dependencies. It depends only on the platform libraries for
window management, input and GPU drawing.

- [Linux](/doc/install#linux)
- [macOS, iOS, tvOS](/doc/install#apple)
- [Windows](/doc/install#windows)
- [Android](/doc/install#android)
- [WebAssembly](/doc/install#wasm)

Gio supports the latest released version of
[Go](https://golang.org/dl) in module mode. Earlier versions of Go and
GOPATH mode might work, but no effort is made to keep them working.

## Running Gio programs

Use the `go` tool to initialize a new module and run the "hello"
example:

	$ go mod init example.com
	$ go run gioui.org/example/hello

should display a simple message in a window.

The command

	$ go run gioui.org/example/kitchen

is another example that demonstrates the material design widgets.

## Running on mobiles

For Android, iOS, tvOS the `gogio` tool can build and package a Gio program for you.

To build an Android .apk file from the `gophers` example:

	$ go run gioui.org/cmd/gogio -target android gioui.org/example/kitchen

To build for the iOS simulator:

	$ go run gioui.org/cmd/gogio -target ios -appid <bundle-id> gioui.org/example/gophers

See the [running on mobile](/doc/mobile) page for more information.


## Webassembly/WebGL

To run a Gio program in a compatible browser, the `gogio` tool can output a directory ready to
serve. With the `goxec` tool you don't even need a web server:

	$ go run gioui.org/cmd/gogio -target js gioui.org/example/gophers
	$ go get github.com/shurcooL/goexec
	$ goexec 'http.ListenAndServe(":8080", http.FileServer(http.Dir("gophers")))'

Open http://localhost:8080 in a browser to run the program.


## Integration with existing projects

See the [integration guide](/doc/integrate) for details on using
Gio with existing projects.


## Programs using Gio

- [Scatter](https://scatter.im), an implementation of the Signal protocol over email.


## Resources

- [Immediate Mode GUI Programming](https://eliasnaur.com/blog/immediate-mode-gui-programming)
- [FAQ](/doc/faq).
- [Gophercon 2019 talk](https://www.youtube.com/watch?v=9D6eWP4peYM) about Gio and [Scatter](https://scatter.im).
[Slides](https://go-talks.appspot.com/github.com/eliasnaur/gophercon-2019-talk/gophercon-2019.slide),
[Demos](https://github.com/eliasnaur/gophercon-2019-talk).
- [Gophercon UK 2019 talk](https://www.youtube.com/watch?v=PxnL3-Sex3o) demonstrating a Gio program built from scratch.
[Slides](https://go-talks.appspot.com/github.com/eliasnaur/gophercon-uk-2019-talk/gophercon-uk-2019-live.slide),
[Demos](https://github.com/eliasnaur/gophercon-uk-2019-talk).

## Community Calls

- [2020-04-21](https://www.youtube.com/watch?v=4qiHYE81nIE)

## Source code

The source code, mailing lists and issue tracker for Gio are [hosted on sourcehut](https://sr.ht/~eliasnaur/gio).

## Issues

File bugs and TODOs through the [issue tracker](https://todo.sr.ht/~eliasnaur/gio) or send an email
to [~eliasnaur/gio@todo.sr.ht](mailto:~eliasnaur/gio@todo.sr.ht). For general discussion, use the
[~eliasnaur/gio@lists.sr.ht mailing list](https://lists.sr.ht/~eliasnaur/gio) or the
[#gioui](https://gophers.slack.com/archives/CM87SNCGM) Gophers Slack channel.


## Contributing

Post patches to the [gio-patches list](https://lists.sr.ht/~eliasnaur/gio-patches). No Sourcehut
account is required and you can post without being subscribed.

You can also use the Sourcehut web-based flow for submitting patches,
similar to other source forges. See the [contribution
guide](/doc/contribute) for more details.


## Sponsors

<div class="sponsor">
	<a href="https://orijtech.com/?referrer=gioui.org">
		<img srcset="/files/orijtech/orijtech.png,
					 /files/orijtech/orijtech@2x.png 2x,
					 /files/orijtech/orijtech@3x.png 3x"
					 src="/files/orijtech/orijtech@3x.png" alt="Orijtech, Inc." width="350">
		<em>"Observability and infrastructure for high performance systems and the cloud."</em>
	</a>
</div>


Gio's main developer is working full-time on Gio, 100% supported by
sponsorships. Please consider [sponsoring Gio](https://github.com/sponsors/eliasnaur) if you find it
useful. Sponsorships are handled by GitHub Sponsors and are easy to
set up.

## Donations

Bitcoin donations are gladly accepted to [bc1q8xw95urett00f4xs3v66p2l6xp2mfk5erpe5ug](bitcoin:bc1q8xw95urett00f4xs3v66p2l6xp2mfk5erpe5ug).
Donations will go toward hosting expenses and for supporting the author's full time work on Gio.

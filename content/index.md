---
title: Gio UI
subtitle: Cross-Platform GUI for Go
childrennolink: true
children:
    - doc/install
    - doc/learn
    - doc/architecture
    - doc/community
    - doc/contribute
    - doc/faq
---

Gio implements portable immediate mode GUI programs in Go. Gio programs run on
all the major platforms: iOS/tvOS, Android, Linux (Wayland/X11), macOS,
Windows, FreeBSD, OpenBSD, and experimental support for browsers (Webassembly/WebGL).
There is a [unikernel port](https://eliasnaur.com/unik) for running Gio programs in virtual machines.

Gio includes an efficient vector renderer based on the [Pathfinder
project](https://github.com/servo/pathfinder), and an experimental renderer
based on the [piet-gpu project](https://github.com/linebender/piet-gpu). Both
renderers support Vulkan, Metal, Direct3D 11, and OpenGL ES. For low-end
devices there is a CPU fallback that runs on unextended OpenGL ES 2.0.

Text and other shapes are rendered using only their outlines without baking them into texture images,
to support efficient animations, transformed drawing and display resolution independence.

This is a screenshot of the [Kitchen
example](https://git.sr.ht/~eliasnaur/gio-example/tree/main/kitchen/kitchen.go). If your browser
supports WebAssembly and WebGL, run the example by pressing the run
button.

{data-run="wasm" data-pkg="kitchen" data-size="800x600"}
<img src="/files/wasm/kitchen.png" alt="Kitchen screenshot" width="800"/>


## Documentation

The [architecture
document](/doc/architecture) is a good introduction to Gio concepts
and API.

The [examples](https://pkg.go.dev/gioui.org/example) give a feel of the
structure of typical Gio programs.

Jon Strand's [tutorial](https://jonegil.github.io/gui-with-gio/) is an great
step-by-step guide to writing Gio programs.

The ["Immediate Mode GUI Programming"](https://eliasnaur.com/blog/immediate-mode-gui-programming)
article compares Gio's immediate mode design with the traditional
retained mode APIs such as the browser DOM.

## Reference documentation

[Operations](https://pkg.go.dev/gioui.org/op) and stateful operation
lists are the low-level primitives of Gio. The important operations
are for [drawing](https://pkg.go.dev/gioui.org/op/paint) and
[clipping](https://pkg.go.dev/gioui.org/op/clip), as well as
[pointer](https://pkg.go.dev/gioui.org/io/pointer) and
[keyboard](https://pkg.go.dev/gioui.org/io/key) input.

The [layout](https://pkg.go.dev/gioui.org/layout) package implements
useful layouts, while the [widget](https://pkg.go.dev/gioui.org/widget)
and [widget/material](https://pkg.go.dev/gioui.org/widget/material)
packages implement common user interface widgets. The
[gesture](https://pkg.go.dev/gioui.org/gesture) package detects common
gestures from lower-level input events.

Layouts, widgets and gestures are all implemented in terms of operations.

Package [app](https://pkg.go.dev/gioui.org/app) is for creating
windows and apply operations to them. Only the app package and its
sub-packages have native dependencies, making Gio [highly
portable](https://pkg.go.dev/gioui.org/example/glfw).

[![GoDoc](https://pkg.go.dev/badge/gioui.org.svg)](https://pkg.go.dev/gioui.org)

## Installation

Gio is designed to work with very few dependencies. It depends only on the
platform libraries for window management, input and GPU drawing.

<div class="big-links">
    <a href="/doc/install/linux">Linux</a>
    <a href="/doc/install/windows">Windows</a>
    <a href="/doc/install/macos">macOS</a>
    <a href="/doc/install/android">Android</a>
    <a href="/doc/install/ios">iOS / tvOS</a>
    <a href="/doc/install/wasm">WebAssembly</a>
</div>

Currently Gio targets the latest released version of [Go](https://golang.org/dl)
in module mode. Earlier versions of Go and `GOPATH` mode might work, but no
effort is made to keep them working.

See [Installation](/doc/install) for further information.

## Programs using Gio

- [godcr](https://github.com/planetdecred/godcr), a cross-platform desktop wallet for the Decred cryptocurrency.
- [sprig](https://git.sr.ht/~whereswaldon/sprig), a client for the [Arbor chat system](https://arbor.chat).
- [Tailscale](https://github.com/tailscale/tailscale-android), a [Tailscale](https://tailscale.com) Android client.
- [Protonet](https://play.google.com/store/apps/details?id=live.protonet), a peer-to-peer chat application. [GitHub Repo](https://github.com/mearaj/protonet)
- [Wormhole William](https://play.google.com/store/apps/details?id=io.sanford.wormhole_william), an end-to-end encrypted file transfer application using the Magic Wormhole protocol. [GitHub repository](https://github.com/psanford/wormhole-william-mobile).
- [Sointu](https://github.com/vsariola/sointu/), a modular software synthesizer to easily produce music for 4k intros.
- [Photon](https://gitlab.com/microo8/photon), a fast RSS reader as light as a photon.

## Sponsors

<div class="sponsors">
	<a class="sponsor" href="https://decred.org/">
		<img src="/files/sponsors/decred.png" alt="decred.org">
		<em>"Decred - Secure. Adaptable. Sustainable."</em>
	</a>
</div>

Development of Gio is funded by sponsorships. If you find Gio useful, please consider sponsoring the
[project on OpenCollective](https://opencollective.com/gioui) or one or more of its developers directly.

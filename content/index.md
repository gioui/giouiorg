---
title: Gio UI
subtitle: Cross-Platform GUI for Go
childrennolink: true
children:
    - news
    - doc/install
    - doc/learn
    - doc/showcase
    - doc/architecture
    - doc/community
    - doc/contribute
    - doc/faq
---

Gio is a library for writing cross-platform immediate mode GUI-s in Go. Gio
supports all the major platforms: Linux, macOS, Windows, Android, iOS, FreeBSD,
OpenBSD and WebAssembly.

For a quick demonstration take a look at the WebAssembly demo below.
_This requires a browser that supports WebAssembly._

{data-run="wasm" data-pkg="kitchen" data-size="800x600"}
<img src="/files/wasm/kitchen.png" alt="Kitchen screenshot" width="800"/>

The source for the [Kitchen project](https://git.sr.ht/~eliasnaur/gio-example/tree/main/kitchen/kitchen.go).

## Getting Started

Gio is designed to work with very few dependencies. It depends only on the
platform libraries for window management, input and GPU drawing.

To install the necessary dependencies, take a look at:

<div class="big-links">
    <a href="/doc/install/linux">Linux</a>
    <a href="/doc/install/windows">Windows</a>
    <a href="/doc/install/macos">macOS</a>
    <a href="/doc/install/android">Android</a>
    <a href="/doc/install/ios">iOS / tvOS</a>
    <a href="/doc/install/wasm">WebAssembly</a>
</div>

Once you have everything installed head over to [Learn](/doc/learn), which
contains links to get you started with Gio.

<div class="big-links">
    <a href="/doc/learn/get-started">First Project<p>Hello World.</p></a>
    <a href="/doc/learn">Learn<p>More helpful resources.</p></a>
</div>

## Showcase

<div class="tiles">
    <a href="/doc/showcase/godcr" style="background-image: url('/doc/showcase/godcr/1.png')">
        <div class="title">godcr</div>
    </a>
    <a href="/doc/showcase/tailscale" style="background-image: url('/doc/showcase/tailscale/1.png')">
        <div class="title">Tailscale</div>
    </a>
    <a href="/doc/showcase/gotraceui" style="background-image: url('/doc/showcase/gotraceui/1.webp')">
        <div class="title">gotraceui</div>
    </a>
    <a href="/doc/showcase/sointu" style="background-image: url('/doc/showcase/sointu/1.png')">
        <div class="title">Sointu</div>
    </a>
    <a href="/doc/showcase/protonet" style="background-image: url('/doc/showcase/protonet/1.png')">
        <div class="title">Protonet</div>
    </a>
    <a class="centered" href="/doc/showcase"><div class="title">More here ...</div></a>
</div>

## Why?

Gio helps Go developers to build efficient, fluid, and portable GUIs across
all major platforms. It combines bleeding-edge 2D graphics technology with the
flexibility of the immediate mode graphics paradigm to create a compelling and
consistent foundation for application development.

Gio includes an efficient vector renderer based on the [Pathfinder project]
(https://github.com/servo/pathfinder) implemented on OpenGL ES and Direct3D 11,
and is migrating towards an even more efficient compute-shader-based renderer
built atop [piet-gpu](https://github.com/linebender/piet-gpu). Text and other
shapes are rendered using only their outlines without baking them into texture
images, to support efficient animations, transformed drawing and pixel
resolution independence.

## Sponsors

Development of Gio is funded by sponsorships. If you find Gio useful, please consider sponsoring the
[project on OpenCollective](https://opencollective.com/gioui) or one or more of its developers directly.

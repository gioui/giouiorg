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

- [godcr](https://github.com/planetdecred/godcr), a cross-platform desktop wallet for the Decred cryptocurrency.
- [sprig](https://git.sr.ht/~whereswaldon/sprig), a client for the [Arbor chat system](https://arbor.chat).
- [Tailscale](https://github.com/tailscale/tailscale-android), a [Tailscale](https://tailscale.com) Android client.
- [Protonet](https://play.google.com/store/apps/details?id=live.protonet), a peer-to-peer chat application. [GitHub Repo](https://github.com/mearaj/protonet)
- [Wormhole William](https://play.google.com/store/apps/details?id=io.sanford.wormhole_william), an end-to-end encrypted file transfer application using the Magic Wormhole protocol. [GitHub repository](https://github.com/psanford/wormhole-william-mobile).
- [Sointu](https://github.com/vsariola/sointu/), a modular software synthesizer to easily produce music for 4k intros.
- [Photon](https://gitlab.com/microo8/photon), a fast RSS reader as light as a photon.

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

<div class="sponsors">
	<a class="sponsor" href="https://decred.org/">
		<img src="/files/sponsors/decred.png" alt="decred.org">
		<em>"Decred - Secure. Adaptable. Sustainable."</em>
	</a>
</div>

Development of Gio is funded by sponsorships. If you find Gio useful, please consider sponsoring the
[project on OpenCollective](https://opencollective.com/gioui) or one or more of its developers directly.
